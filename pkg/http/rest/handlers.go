// SPDX-FileCopyrightText: 2022 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/osrgroup/product-model-toolkit/model"
	"github.com/osrgroup/product-model-toolkit/pkg/server/services"
	"github.com/pkg/errors"
	"github.com/spdx/tools-golang/idsearcher"
	"github.com/spdx/tools-golang/tvsaver"

	"github.com/labstack/echo/v4"
)

type result struct {
	Result interface{} `json:"result"`
}

func handleEntryPoint(c echo.Context) error {
	return c.JSON(http.StatusOK, result{Result: c.Echo().Routes()})
}

func handleVersion(c echo.Context) error {
	return c.JSON(http.StatusOK, "1.0.0")
}

func handleHealth(c echo.Context) error {

	return c.JSON(http.StatusOK, result{Result: "UP"})
}

func findAllProducts(srv services.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		prods, err := srv.FindAllProducts()
		if err != nil {
			log.Printf("Error: %s\n", err.Error())
			return c.JSON(http.StatusNotFound, result{Result: err.Error()})
		}

		return c.JSON(http.StatusOK, prods)
	}
}

func findProductByID(srv services.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Printf("Error: %s\n", err.Error())
			return c.JSON(http.StatusBadRequest, result{Result: err.Error()})
		}

		prod, err := srv.FindProductByID(id)
		if err != nil {
			log.Printf("Error: %s\n", err.Error())
			return c.JSON(http.StatusNotFound, result{Result: fmt.Sprintf("unable fo find product with ID %v", id)})
		}

		return c.JSON(http.StatusOK, prod)
	}
}

func updateProductByID(srv services.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Printf("Error: %s\n", err.Error())
			return c.JSON(http.StatusBadRequest, result{Result: err.Error()})
		}

		updateDetails, err := getUpdateDetails(c)
		if err != nil {
			log.Printf("Error: %s\n", err.Error())
			return c.JSON(http.StatusBadRequest, result{Result: err.Error()})
		}

		name, version := updateDetails[0], updateDetails[1]
		if err := srv.UpdateProductByID(id, name, version); err != nil {
			log.Printf("Error: %s\n", err.Error())
			return c.JSON(http.StatusNotFound, result{Result: fmt.Sprintf("unable fo find product with ID %v", id)})
		}

		return c.JSON(http.StatusOK, result{Result: "updated!"})
	}
}

func deleteProductByID(srv services.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Printf("Error: %s\n", err.Error())
			return c.JSON(http.StatusBadRequest, result{Result: err.Error()})
		}
		err = srv.DeleteProductByID(id)
		if err != nil {
			log.Printf("Error: %s\n", err.Error())
			return c.JSON(http.StatusNotFound, result{Result: fmt.Sprintf("unable fo find product with ID %v", id)})
		}
		return c.JSON(http.StatusOK, result{Result: fmt.Sprintf("product %v deleted", id)})
	}
}

// This handler is responsible for getting the required url
// from user and download the git file and store it on a predefined path
func download(srv services.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		downloadDetails, err := getDownloadDetails(c)
		if err != nil {
			log.Printf("error: %s\n", err.Error())
			return c.JSON(http.StatusBadRequest, result{Result: err.Error()})
		}

		downloadData, err := srv.Download(downloadDetails)
		if err != nil {
			log.Printf("error: %s\n", err.Error())
			return c.JSON(http.StatusBadRequest, result{Result: err.Error()})
		}

		err = srv.StoreDownloadedRepo(downloadData)
		if err != nil {
			log.Printf("error: %s\n", err.Error())
			// remove the path directory
			path := downloadDetails[1]
			err := os.Remove(path)
			if err != nil {
				log.Printf("error: %s\n", err.Error())
				return c.JSON(http.StatusInternalServerError, result{Result: err.Error()})
			}
			return c.JSON(http.StatusInternalServerError, result{Result: err.Error()})
		}

		return c.JSON(http.StatusOK, result{Result: "Download completed"})
	}
}

func getAllDownloadedRepos(srv services.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		data, err := srv.FindAllDownloadedRepos()
		if err != nil {
			log.Printf("error: %s\n", err.Error())
			return c.JSON(http.StatusBadRequest, result{Result: err.Error()})
		}

		return c.JSON(http.StatusOK, result{Result: data})
	}
}

func getDiffProducts(srv services.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		
		// get json body
		jsonBody, err := getJSONRawBody(c)
		if err != nil {
			log.Printf("Error: %s\n", err.Error())
			return c.JSON(http.StatusBadRequest, result{Result: err.Error()})
		}

		// read data
		firstFile := jsonBody["first"]
		secondFile := jsonBody["second"]

		if firstFile == "" {
			err := errors.New("empty parameter")
			return c.JSON(http.StatusBadRequest, result{Result: err.Error()})
		}

		if secondFile == "" {
			err := errors.New("empty parameter")
			return c.JSON(http.StatusBadRequest, result{Result: err.Error()})
		}

		data, err := srv.FindAllDiff(firstFile, secondFile)
		if err != nil {
			log.Printf("error: %s\n", err.Error())
			return c.JSON(http.StatusBadRequest, result{Result: err.Error()})
		}
		return c.JSON(http.StatusOK, result{Result: data})
	}
}

func importFromScanner(iSrv services.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		importDetails, err := getImportDetails(c)
		if err != nil {
			log.Printf("Error: %s\n", err.Error())
			return c.JSON(http.StatusBadRequest, result{Result: err.Error()})
		}
		scanner, importPath, importName := importDetails[0], importDetails[1], importDetails[2]

		rb, err := os.ReadFile(importPath)
		if err != nil {
			log.Printf("Error: %s\n", err.Error())
			return c.JSON(http.StatusBadRequest, result{Result: err.Error()})
		}
		// get the scanner from the url param
		// scanner := strings.ToLower(c.Param("scanner"))
		r := bytes.NewReader(rb)
		// define product and error variable
		var prod *model.Product

		// switch over the scanner name
		switch scanner {
		case "spdx":
			prod, err = iSrv.SPDXImport(r, importName)
		case "composer":
			prod, err = iSrv.ComposerImport(r, importName)
		case "file-hasher":
			prod, err = iSrv.FileHasherImport(r, importName)
		case "scanner":
			prod, err = iSrv.ScannerImport(r, importName)
		default:
			return c.JSON(http.StatusBadRequest, result{Result: fmt.Sprintf(
				"received result file with content length %d, "+
					"but will not import content, because there is no importer for the scanner '%s'",
				c.Request().ContentLength, scanner)})
		}

		// check error
		if err != nil {
			c.Error(errors.Wrap(err, fmt.Sprintf("unable to perform import for scanner %s", scanner)))
		}

		return c.JSON(http.StatusCreated, result{Result: fmt.Sprintf("successfully parsed content from scanner %s."+
			"\nProduct id: %v\nFound %v packages\n", scanner, prod.ID, len(prod.Components))})
	}
}

func scan(iSrv services.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		// get the scan details
		scanDetails, err := getScanDetails(c)

		// check the errors
		if err != nil {
			log.Printf("Error: %s\n", err.Error())
			return c.JSON(http.StatusBadRequest, result{Result: err.Error()})
		}

		// send scan data to the service
		res, err := iSrv.Scan(scanDetails)

		// check the errors
		if err != nil {
			log.Printf("Error: %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, result{Result: err.Error()})
		}

		// return the required json data
		return c.JSON(
			http.StatusOK,
			result{Result: struct {
				Report string `json:"report"`
				//Data   string `json:"data"`
			}{
				fmt.Sprintf(res)},
			},
		)
	}
}

func checkLicenseCompatibility(srv services.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Printf("Error: %s\n", err.Error())
			return c.JSON(
				http.StatusInternalServerError,
				result{Result: err.Error()},
			)
		}

		res, err := srv.CheckLicenseCompatibility(id)
		if err != nil {
			log.Printf("Error: %s\n", err.Error())
			return c.JSON(
				http.StatusNotFound,
				result{Result: err.Error()},
			)
		}
		return c.JSON(http.StatusOK, result{Result: res})
	}
}

func getScanDetails(c echo.Context) ([]string, error) {
	// get json body
	jsonBody, err := getJSONRawBody(c)
	if err != nil {
		log.Printf("Error: %s\n", err.Error())
		return nil, err
	}

	// read data
	scannerName := jsonBody["scannerName"]
	source := jsonBody["source"]
	output := jsonBody["output"]

	return []string{scannerName, source, output}, nil
}

func getDownloadDetails(c echo.Context) ([]string, error) {
	// get json body
	jsonBody, err := getJSONRawBody(c)
	if err != nil {
		return nil, err
	}

	// read data
	url := jsonBody["url"]
	output := jsonBody["path"]

	return []string{url, output}, nil
}

func getImportDetails(c echo.Context) ([]string, error) {
	// get json body
	jsonBody, err := getJSONRawBody(c)
	if err != nil {
		return nil, err
	}

	// read data
	importType := jsonBody["importType"]
	importPath := jsonBody["importPath"]
	importName := jsonBody["importName"]

	return []string{importType, importPath, importName}, nil
}


func getUpdateDetails(c echo.Context) ([]string, error) {
	// get json body
	jsonBody, err := getJSONRawBody(c)
	if err != nil {
		return nil, err
	}

	// read data
	name := jsonBody["name"]
	version := jsonBody["version"]

	return []string{name, version}, nil
}

func getJSONRawBody(c echo.Context) (map[string]string, error) {

	if c.Request().ContentLength == 0 {
		return nil, errors.New("no content received")
	}

	jsonBody := make(map[string]string)
	err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	if err != nil {
		log.Printf("Error: %s\n", err.Error())
		return nil, err
	}

	return jsonBody, nil
}

func exportWithType(iSrv services.Service) echo.HandlerFunc {
	return func(c echo.Context) error {

		// get json body
		jsonBody, err := getJSONRawBody(c)
		if err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				result{Result: err.Error()},
			)
		}

		// read data
		exportId := jsonBody["exportId"]
		exportType := jsonBody["exportType"]
		exportPath := jsonBody["exportPath"]

		// switch over the scanner name
		switch exportType {
		case "spdx":
			_, exportPath, err = iSrv.SPDXExport(exportId, exportPath)
			if err != nil {
				return c.JSON(
					http.StatusInternalServerError,
					result{Result: err.Error()},
				)
			}
			return c.JSON(
				http.StatusCreated,
				result{Result: fmt.Sprintf("export path: %v", exportPath)},
			)
		case "human-read":
			exportPath, err = iSrv.ReportExport(exportId, exportPath)
			if err != nil {
				return c.JSON(
					http.StatusInternalServerError,
					result{Result: err.Error()},
				)
			}
			return c.JSON(
				http.StatusCreated,
				result{Result: fmt.Sprintf("export path: %v", exportPath)},
			)
		case "compatibility":
			report, exportPath, err := iSrv.CompatibilityExport(exportId, exportPath)
			if err != nil {
				return c.JSON(
					http.StatusInternalServerError,
					result{Result: err.Error()},
				)
			}
			fmt.Println(report)
			return c.JSON(
				http.StatusCreated,
				result{Result: struct {
					Report string `json:"report"`
					Path   string `json:"path"`
				}{
					Report: report,
					Path:   exportPath,
				}},
			)
		default:
			return c.JSON(
				http.StatusNotAcceptable,
				result{
					Result: fmt.Sprintf("export type %q not supported", exportType),
				},
			)
		}
	}
}

func searchSPDX(srv services.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		// get json body
		jsonBody, err := getJSONRawBody(c)
		if err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				result{Result: err.Error()},
			)
		}
		// read data
		packageName := jsonBody["name"]
		packageRootDir := jsonBody["dir"]
		fileOut := jsonBody["out"]

		config := &idsearcher.Config2_2{

			NamespacePrefix: "https://example.com/whatever/testdata-",

			BuilderPathsIgnored: []string{
				"/.git/",
				"**/__pycache__/",
				"/.ignorefile",
				"**/.DS_Store",
				"/vendor/",
			},

			SearcherPathsIgnored: []string{

				"/Documentation/process/license-rules.rst",
				"/LICENSES/",
			},
		}

		doc, err := idsearcher.BuildIDsDocument2_2(packageName, packageRootDir, config)
		if err != nil {
			log.Printf("error: %v\n", err.Error())
			return c.JSON(http.StatusInternalServerError, result{Result: err.Error()})
		}

		log.Printf("successfully created document and searched for IDs for package %s\n", packageName)

		w, err := os.Create(fileOut)
		if err != nil {
			log.Printf("error: %v\n", err.Error())
			return c.JSON(http.StatusInternalServerError, result{Result: err.Error()})
		}
		defer func() {
			if err := w.Close(); err != nil {
				log.Printf("error: %v\n", err.Error())
			}
		}()

		err = tvsaver.Save2_2(doc, w)
		if err != nil {
			log.Printf("error: %v\n", err.Error())
			return c.JSON(http.StatusInternalServerError, result{Result: err.Error()})
		}

		return c.JSON(
			http.StatusOK,
			result{Result: fmt.Sprintf(
				"successfully created document and searched for IDs for package %s", packageName)},
		)
	}
}
