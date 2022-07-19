// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/osrgroup/product-model-toolkit/cnst"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"reflect"
	"sort"
	"strconv"
	"time"

	"github.com/osrgroup/product-model-toolkit/model"
	"github.com/osrgroup/product-model-toolkit/pkg/server/commands"
	convert "github.com/osrgroup/product-model-toolkit/pkg/server/services/convert"
	composer "github.com/osrgroup/product-model-toolkit/pkg/server/services/convert/composer"
	hasher "github.com/osrgroup/product-model-toolkit/pkg/server/services/convert/hasher"
	"github.com/pkg/errors"
	"github.com/spdx/tools-golang/builder"
	"github.com/spdx/tools-golang/reporter"
	"github.com/spdx/tools-golang/spdx"
	"github.com/spdx/tools-golang/spdxlib"
	"github.com/spdx/tools-golang/tvloader"
	"github.com/spdx/tools-golang/tvsaver"
)

var (
	ErrNotFound = errors.New("entity not found")
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// Repository provides access to the product db.
type Repository interface {
	// FindAllProducts returns a list of all products saved in db.
	FindAllProducts() ([]model.Product, error)
	// FindProductByID returns the product with the given ID.
	FindProductByID(id int) (model.Product, error)

	// DeleteProductByID deletes the product with the given ID.
	DeleteProductByID(id int) error

	Download([]string) (*model.DownloadData, error)

	SaveProduct(prod *model.Product) (model.Product, error)

	StoreDownloadedRepo(data *model.DownloadData) error

	FindAllDownloadedRepos() ([]model.DownloadData, error)
}

// Service  provides product querying operations.
type Service interface {
	FindAllProducts() ([]model.Product, error)
	FindProductByID(id int) (model.Product, error)
	DeleteProductByID(id int) error
	Download([]string) (*model.DownloadData, error)
	CheckLicenseCompatibility(id int) ([]localResult, error)
	// import
	ComposerImport(io.Reader, string) (*model.Product, error)
	SPDXImport(io.Reader, string) (*model.Product, error)
	FileHasherImport(io.Reader, string) (*model.Product, error)
	ScannerImport(io.Reader, string) (*model.Product, error)

	// export
	SPDXExport(exportId, exportPath string) (*spdx.Document2_2, string, error)
	ReportExport(exportId, exportPath string) (string, error)
	CompatibilityExport(exportId, exportPath string) (string, string, error)
	Scan(scanDetails []string) (string, error)
	StoreDownloadedRepo(data *model.DownloadData) error
	FindAllDownloadedRepos() ([]model.DownloadData, error)
}

type service struct {
	r Repository
}

// NewService creates a querying service with all necessary dependencies.
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) FindAllDownloadedRepos() ([]model.DownloadData, error) {
	return s.r.FindAllDownloadedRepos()
}

func (s *service) StoreDownloadedRepo(data *model.DownloadData) error {
	return s.r.StoreDownloadedRepo(data)
}

// FindAllProducts returns all existing products.
func (s *service) FindAllProducts() ([]model.Product, error) {
	return s.r.FindAllProducts()
}

// FindProductByID returns the product with the given ID.
func (s *service) FindProductByID(id int) (model.Product, error) {
	return s.r.FindProductByID(id)
}

func (s *service) DeleteProductByID(id int) error {
	return s.r.DeleteProductByID(id)
}

func (s *service) Download(downloadDetails []string) (*model.DownloadData, error) {
	return s.r.Download(downloadDetails)
}
func (s *service) Scan(scanDetails []string) (string, error) {
	scannerName, source, output := scanDetails[0], scanDetails[1], scanDetails[2]

	sort.Slice(commands.Available[:], func(i, j int) bool {
		return commands.Available[i].Name <= commands.Available[j].Name
	})

	idx := sort.Search(len(commands.Available), func(i int) bool {
		return string(commands.Available[i].Name) >= scannerName
	})

	if item := commands.Available[idx]; item.Name == scannerName {
		if item.DockerCmd == "" {
			return "", errors.New("The scanner has not been defined yet!")
		}
		dockerCmd := fmt.Sprintf(item.DockerCmd, source, output, item.DockerImg)
		// execute the command
		_, err := exec.Command("/bin/sh", "-c", dockerCmd).CombinedOutput()
		if err != nil {
			log.Printf("docker command failed: %s", dockerCmd)
			log.Printf("error: %v\n", err.Error())
			return "", err
		}
		return fmt.Sprintf("The output path: %v", output), nil

	} else {
		err := errors.New("The scanner has not been defined yet!")
		return "", err
	}
}

type localPackage struct {
	Name    string `json:"name"`
	License string `json:"license"`
}

type localProd struct {
	ProductId string `json:"product_id"`
	License   string `json:"license"`
}

type localResult struct {
	Package      localPackage
	Product      localProd
	IsCompatible bool `json:"is_compatible"`
}

func (s *service) CheckLicenseCompatibility(id int) ([]localResult, error) {
	// get the product from id
	prod, err := s.FindProductByID(id)
	if err != nil {
		return nil, err
	}

	configFileData, err := ioutil.ReadFile("./licenseCompatibility.json")
	if err != nil {
		return nil, err
	}

	g, err := convertConfigFileToGraph(configFileData)
	if err != nil {
		return nil, err
	}

	var finalResult []localResult
	// iterate over the list of licenses
	for _, v := range prod.Components {
		lr := new(localResult)
		lr.Package = localPackage{
			v.Package,
			v.License.SPDXID,
		}
		lr.Product = localProd{
			string(prod.ID),
			prod.License,
		}
		if !IsAncestor(g, prod.License, v.License.SPDXID) {
			lr.IsCompatible = false
		} else {
			lr.IsCompatible = true
		}
		finalResult = append(finalResult, *lr)

	}
	return finalResult, nil

}

// ComposerImport import a Composer representation of the BOM and store it in the DB.
func (s *service) ComposerImport(input io.Reader, importName string) (*model.Product, error) {
	var c convert.Converter = new(composer.Composer)
	prod, err := c.Convert(input)
	if err != nil {
		msg := fmt.Sprintf("Error while parsing Composer JSON body: %v", err)
		return nil, errors.New(msg)
	}

	if importName != cnst.Empty {
		prod.Name = importName
	}

	pSaved, err := s.r.SaveProduct(prod)
	if err != nil {
		msg := fmt.Sprintf("error while saving the product to the DB: %v", err)
		return nil, errors.New(msg)
	}

	return &pSaved, nil
}

func spdxToProduct(doc *spdx.Document2_1) (*model.Product, error) {

	components := []model.Component{}

	for _, p := range doc.Packages {

		cmp := model.Component{
			UID:     string(p.PackageSPDXIdentifier),
			Name:    p.PackageName,
			Package: p.PackageSummary,
			Version: p.PackageVersion,
			License: model.License{
				SPDXID:           string(p.PackageSPDXIdentifier),
				DeclaredLicense:  p.PackageLicenseDeclared,
				ConcludedLicense: p.PackageLicenseConcluded,
			},
		}
		components = append(components, cmp)
	}

	var ref string
	if len(doc.CreationInfo.ExternalDocumentReferences) > 0 {
		ref = fmt.Sprintf("%v", reflect.TypeOf(doc.CreationInfo.ExternalDocumentReferences))
	}

	prod := &model.Product{
		Name:              doc.CreationInfo.DocumentName,
		Version:           doc.CreationInfo.SPDXVersion,
		Description:       doc.CreationInfo.DocumentComment,
		Comment:           doc.CreationInfo.CreatorComment,
		HomepageURL:       doc.CreationInfo.DocumentNamespace,
		ExternalReference: ref,
		Components:        components,
	}

	return prod, nil
}

func productToSPDX(prod *model.Product, exportPath string) (*spdx.Document2_2, string, error) {
	config := &builder.Config2_2{
		NamespacePrefix: "https://example.com/whatever/testdata-",
		CreatorType:     "Tool",
		Creator:         "Product Model Toolkit",
		PathsIgnored: []string{
			"/.git/",
			"**/__pycache__/",
			"/.ignorefile",
			"**/.DS_Store",
			"/vendor/",
		},
	}

	w, err := os.Create(exportPath)
	if err != nil {
		log.Printf("error while creating the export file %q: %v", exportPath, err)
		return nil, "", err

	}
	defer func() {
		if err := w.Close(); err != nil {
			log.Printf("error while closing the export file %q: %v", exportPath, err)
		}
	}()

	packageRootDir := exportPath

	doc, err := builder.Build2_2(prod.Name, packageRootDir, config)
	if err != nil {
		log.Printf("error while building the SPDX document: %v", err)
		return nil, "", err
	}

	packages := make(map[spdx.ElementID]*spdx.Package2_2)
	for _, component := range prod.Components {
		eId := fmt.Sprintf("Package-%v\n", component.Package)

		pk := new(spdx.Package2_2)
		pk.PackageSPDXIdentifier = spdx.ElementID(component.UID)
		pk.PackageName = component.Name
		pk.PackageSummary = component.Package
		pk.PackageVersion = component.Version
		pk.PackageSPDXIdentifier = spdx.ElementID(component.License.SPDXID)
		pk.PackageLicenseDeclared = component.License.DeclaredLicense
		pk.PackageLicenseConcluded = component.License.ConcludedLicense
		pk.FilesAnalyzed = true

		packages[spdx.ElementID(eId)] = pk
	}

	doc.CreationInfo.DocumentName = prod.Name
	doc.CreationInfo.SPDXVersion = prod.Version
	doc.CreationInfo.DocumentComment = prod.Description
	doc.CreationInfo.CreatorComment = prod.Comment
	doc.CreationInfo.DocumentNamespace = prod.HomepageURL
	doc.CreationInfo.ExternalDocumentReferences = []string{prod.ExternalReference}
	doc.Packages = packages

	err = tvsaver.Save2_2(doc, w)
	if err != nil {
		log.Printf("error while saving the SPDX document %q: %v", exportPath, err)
		return nil, "", err
	}

	return doc, exportPath, nil
}

// SPDX import a SPDX representation of the BOM.
func (s *service) SPDXImport(input io.Reader, importName string) (*model.Product, error) {

	doc, err := tvloader.Load2_1(input)
	if err != nil {
		log.Printf("error while loading the SPDX document: %v", err)
		return nil, err
	}
	prod, err := spdxToProduct(doc)
	if err != nil {
		log.Printf("error while converting the SPDX document: %v", err)
		return nil, err
	}

	if importName != cnst.Empty {
		prod.Name = importName
	}

	pSaved, err := s.r.SaveProduct(prod)
	if err != nil {
		log.Printf("error while saving the product: %v", err)
		return nil, err
	}

	return &pSaved, nil

}

// FileHasherImport import a File-Hasher representation of the BOM and store it in the DB.
func (s *service) FileHasherImport(input io.Reader, importName string) (*model.Product, error) {
	var c convert.Converter = new(hasher.Hasher)

	prod, err := c.Convert(input)
	if err != nil {
		return nil, errors.Wrap(err, "error while parsing File-Hasher body")
	}

	if importName != cnst.Empty {
		prod.Name = importName
	}

	pSaved, err := s.r.SaveProduct(prod)
	if err != nil {
		return nil, errors.Wrap(err, "error while saving the product to the DB")
	}

	return &pSaved, nil
}

func (s *service) ReportExport(exportId, exportPath string) (string, error) {
	doc, exportPath, err := s.SPDXExport(exportId, exportPath)
	if err != nil {
		log.Printf("error: %v\n", err.Error())
		return "", err
	}
	if (len(doc.Packages)) > 0 {
		// check whether the SPDX file has at least one package that it describes
		pkgIDs, err := spdxlib.GetDescribedPackageIDs2_2(doc)
		if err != nil {
			log.Printf("error: %v\n", err.Error())
			return "", err
		}

		// it does, so we'll go through each one
		for _, pkgID := range pkgIDs {
			pkg, ok := doc.Packages[pkgID]
			if !ok {
				err := errors.New(fmt.Sprintf(
					"package %s has described relationship but ID not found\n", string(pkgID)))
				log.Printf("error: %v\n", err.Error())
				continue
			}

			// check whether the package had its files analyzed
			if !pkg.FilesAnalyzed {
				err := errors.New(fmt.Sprintf(
					"package %s (%s) had FilesAnalyzed: false\n", string(pkgID), pkg.PackageName))

				log.Printf("error: %v\n", err)
				return exportPath, err
			}

			// also check whether the package has any files present
			if pkg.Files == nil || len(pkg.Files) < 1 {
				err := errors.New(fmt.Sprintf(
					"package %s (%s) has no Files\n", string(pkgID), pkg.PackageName))
				return exportPath, err
			}

			// if we got here, there's at least one file
			// generate and print a report of the Package's Files' LicenseConcluded
			// values, sorted by # of occurrences
			log.Printf("Package %s (%s)\n", string(pkgID), pkg.PackageName)
			err = reporter.Generate2_2(pkg, os.Stdout)
			if err != nil {
				log.Printf("error: %v\n", err.Error())
				return "", err
			}
		}
		return fmt.Sprintf("successfully exported to: %v", exportPath), nil
	} else {
		err := errors.New("the length of packages is zero")
		return "", err
	}
}

func (s *service) SPDXExport(exportId, exportPath string) (*spdx.Document2_2, string, error) {
	// get the product from id
	id, err := strconv.Atoi(exportId)
	if err != nil {
		log.Printf("error: %v\n", err.Error())
		return nil, "", err
	}
	prod, err := s.FindProductByID(id)
	if err != nil {
		log.Printf("error: %v\n", err.Error())
		return nil, "", err
	}

	// convert the product to spdx
	doc, exportPath, err := productToSPDX(&prod, exportPath)
	if err != nil {
		log.Printf("error: %v\n", err.Error())
		return nil, "", err
	}
	return doc, exportPath, nil
}

func (s *service) CompatibilityExport(exportId, exportPath string) (string, string, error) {
	// get the product from id
	id, err := strconv.Atoi(exportId)
	if err != nil {
		return "", "", err
	}
	prod, err := s.FindProductByID(id)
	if err != nil {
		return "", "", err
	}

	configFileData, err := ioutil.ReadFile("./licenseCompatibility.json")
	if err != nil {
		return "", "", err
	}

	g, err := convertConfigFileToGraph(configFileData)
	if err != nil {
		return "", "", err
	}

	var rp string
	// iterate over the list of licenses
	for _, v := range prod.Components {
		if !IsAncestor(g, v.License.SPDXID, prod.License) {
			localResult := fmt.Sprintf("The [PACKAGE] %s with [DATABASE ID] %d and [LICENSE] %s, is not compatible with [PRODUCT ID] %d with [LICENSE] %s\n", v.Package, v.ID, v.License.SPDXID, prod.ID, prod.License)
			rp += localResult
		}
	}

	// create a report file and write down all strings into it
	reportFile, err := os.Create(exportPath)
	if err != nil {
		return "", "", err
	}
	defer reportFile.Close()

	reportFile.WriteString(rp)

	return rp, exportPath, nil
}

func convertConfigFileToGraph(data []byte) (*Graph, error) {
	var cnf map[string]Value
	if err := json.Unmarshal(data, &cnf); err != nil {
		return nil, err
	}

	fmt.Println(cnf)
	// create a graph
	g := NewDirectedGraph()
	for k, _ := range cnf {
		g.AddVertex(k)
	}

	for k, v := range cnf {
		for _, vv := range v.IncludableIn {
			g.AddEdge(k, vv)
		}
	}

	return g, nil
}

type Value struct {
	IncludableIn []string `json:"includable_in"`
}

// ComposerImport import a Composer representation of the BOM and store it in the DB.
func (s *service) ScannerImport(input io.Reader, importName string) (*model.Product, error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(input)
	strInput := buf.String()

	prod := new(model.Product)
	err := json.Unmarshal([]byte(strInput), &prod)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if importName != cnst.Empty {
		prod.Name = importName
	} else {
		prod.Name = fmt.Sprintf("product-%s", randStringRunes(10))
	}

	pSaved, err := s.r.SaveProduct(prod)
	if err != nil {
		msg := fmt.Sprintf("error while saving the product to the DB: %v", err)
		return nil, errors.New(msg)
	}

	return &pSaved, nil

}
