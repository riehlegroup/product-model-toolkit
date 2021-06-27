// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package services

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"bytes"
	"strconv"
	"encoding/json"
	convert "github.com/osrgroup/product-model-toolkit/pkg/server/services/convert"
	composer "github.com/osrgroup/product-model-toolkit/pkg/server/services/convert/composer"
	hasher "github.com/osrgroup/product-model-toolkit/pkg/server/services/convert/hasher"
	"math/rand"
	"time"
	"github.com/osrgroup/product-model-toolkit/model"
	"github.com/pkg/errors"
	"github.com/spdx/tools-golang/builder"
	"github.com/spdx/tools-golang/reporter"
	"github.com/spdx/tools-golang/spdx"
	"github.com/spdx/tools-golang/spdxlib"
	"github.com/spdx/tools-golang/tvloader"
	"github.com/spdx/tools-golang/tvsaver"
)

var (
	// ErrNotFound if a entity couldn't be found in the db.
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

	SaveProduct(prod *model.Product) (model.Product, error)
}

// Service  provides product querying operations.
type Service interface {
	FindAllProducts() ([]model.Product, error)
	FindProductByID(id int) (model.Product, error)

	// import
	ComposerImport(io.Reader) (*model.Product, error)
	SPDXImport(io.Reader) (*model.Product, error)
	FileHasherImport(io.Reader) (*model.Product, error)
	ScannerImport(io.Reader) (*model.Product, error)

	// export
	SPDXExport(exportId, exportPath string) (*spdx.Document2_2, string, error)
	ReportExport(exportId, exportPath string) (string, error)
}

type service struct {
	r Repository
}

// NewService creates a querying service with all necessary dependencies.
func NewService(r Repository) Service {
	return &service{r}
}

// FindAllProducts returns all existing products.
func (s *service) FindAllProducts() ([]model.Product, error) {
	return s.r.FindAllProducts()
}

// FindProductByID returns the product with the given ID.
func (s *service) FindProductByID(id int) (model.Product, error) {
	return s.r.FindProductByID(id)
}

// ComposerImport import a Composer representation of the BOM and store it in the DB.
func (s *service) ComposerImport(input io.Reader) (*model.Product, error) {
	var c convert.Converter = new(composer.Composer)
	prod, err := c.Convert(input)
	if err != nil {
		msg := fmt.Sprintf("Error while parsing Composer JSON body: %v", err)
		return nil, errors.New(msg)
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
			Pkg:     p.PackageSummary,
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
		ClearingState:     nil,
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
		fmt.Printf("error while opening %v for writing: %v\n", exportPath, err)
		return nil, "", err

	}
	defer w.Close()

	packageRootDir := exportPath

	doc, err := builder.Build2_2(prod.Name, packageRootDir, config)
	if err != nil {
		fmt.Printf("error while building document: %v\n", err)
		return nil, "", err
	}

	packages := make(map[spdx.ElementID]*spdx.Package2_2)
	for _, component := range prod.Components {
		eId := fmt.Sprintf("Package-%v\n", component.Name)

		pk := new(spdx.Package2_2)
		pk.PackageSPDXIdentifier = spdx.ElementID(component.UID)
		pk.PackageName = component.Name
		pk.PackageSummary = component.Pkg
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
		fmt.Printf("error while saving %v: %v", exportPath, err)
		return nil, "", err
	}

	return doc, exportPath, nil
}

// SPDX import a SPDX representation of the BOM.
func (s *service) SPDXImport(input io.Reader) (*model.Product, error) {

	doc, err := tvloader.Load2_1(input)
	if err != nil {
		msg := fmt.Sprintf("error while parsing SPDX body: %v", err)
		return nil, errors.New(msg)
	}
	prod, err := spdxToProduct(doc)
	if err != nil {
		return nil, err
	}

	pSaved, err := s.r.SaveProduct(prod)
	if err != nil {
		msg := fmt.Sprintf("error while saving the product to the DB: %v", err)
		return nil, errors.New(msg)
	}

	return &pSaved, nil

}

// FileHasherImport import a File-Hasher representation of the BOM and store it in the DB.
func (s *service) FileHasherImport(input io.Reader) (*model.Product, error) {
	var c convert.Converter = new(hasher.Hasher)

	prod, err := c.Convert(input)
	if err != nil {
		return nil, errors.Wrap(err, "error while parsing File-Hasher body")
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
		return "", err
	}
	if (len(doc.Packages)) > 0 {
		// check whether the SPDX file has at least one package that it describes
		pkgIDs, err := spdxlib.GetDescribedPackageIDs2_2(doc)
		if err != nil {
			fmt.Printf("Unable to get describe packages from SPDX document: %v\n", err)
			return "", err
		}

		// it does, so we'll go through each one
		for _, pkgID := range pkgIDs {
			pkg, ok := doc.Packages[pkgID]
			if !ok {
				fmt.Printf("package %s has described relationship but ID not found\n", string(pkgID))
				continue
			}

			// check whether the package had its files analyzed
			if !pkg.FilesAnalyzed {
				fmt.Printf("package %s (%s) had FilesAnalyzed: false\n", string(pkgID), pkg.PackageName)
				return exportPath, err
			}

			// also check whether the package has any files present
			if pkg.Files == nil || len(pkg.Files) < 1 {
				fmt.Printf("package %s (%s) has no Files\n", string(pkgID), pkg.PackageName)
				return exportPath, err
			}

			// if we got here, there's at least one file
			// generate and print a report of the Package's Files' LicenseConcluded
			// values, sorted by # of occurrences
			fmt.Printf("============================\n")
			fmt.Printf("Package %s (%s)\n", string(pkgID), pkg.PackageName)
			err = reporter.Generate2_2(pkg, os.Stdout)
			if err != nil {
				fmt.Printf("error while generating report: %v\n", err)
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
		return nil, "", err
	}
	prod, err := s.FindProductByID(id)
	if err != nil {
		return nil, "", err
	}

	// convert the product to spdx
	doc, exportPath, err := productToSPDX(&prod, exportPath)
	if err != nil {
		return nil, "", err
	}
	return doc, exportPath, nil
}

// ComposerImport import a Composer representation of the BOM and store it in the DB.
func (s *service) ScannerImport(input io.Reader) (*model.Product, error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(input)

	prod := new(model.Product)
	err := json.Unmarshal(buf.Bytes(), &prod)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	prod.Name = fmt.Sprintf("product-%s", randStringRunes(10))

	pSaved, err := s.r.SaveProduct(prod)
	if err != nil {
		msg := fmt.Sprintf("error while saving the product to the DB: %v", err)
		return nil, errors.New(msg)
	}

	return &pSaved, nil
	
}
