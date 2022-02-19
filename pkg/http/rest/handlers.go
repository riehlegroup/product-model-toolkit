// SPDX-FileCopyrightText: 2022 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	return c.String(http.StatusOK, "1.0.0")
}

func handleHealth(c echo.Context) error {

	return c.JSON(http.StatusOK, result{Result: "UP"})
}

func findAllProducts(srv services.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		prods, err := srv.FindAllProducts()
		if err != nil {
			c.Error(errors.Wrap(err, "unable to find all products"))
		}

		return c.JSON(http.StatusOK, prods)
	}
}

func findProductByID(srv services.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.Error(errors.Wrap(err, fmt.Sprintf("unable to convert query param id with value '%v' to int", idStr)))
		}

		prod, err := srv.FindProductByID(id)
		if err != nil {
			return c.JSON(http.StatusNotFound, result{Result: fmt.Sprintf("unable fo find product with ID %v", id)})

		}

		return c.JSON(http.StatusOK, prod)
	}
}

func deleteProductByID(srv services.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.Error(errors.Wrap(err, fmt.Sprintf("unable to convert query param id with value '%v' to int", idStr)))
		}
		err = srv.DeleteProductByID(id)
		if err != nil {
			return c.JSON(http.StatusNotFound, result{Result: fmt.Sprintf("unable fo find product with ID %v", id)})
		}
		return c.JSON(http.StatusOK, map[string]string{"result": fmt.Sprintf("product %v deleted", id)})
	}
}

// This handler is responsible for getting the required url
// from user and download the git file and store it on a predefined path
func download(iSrv services.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		downloadDetails, err := getDownloadDetails(c)
		if err != nil {
			return c.JSON(http.StatusBadRequest, result{Result: err.Error()})
		}

		err = iSrv.Download(downloadDetails)
		if err != nil {
			return c.JSON(http.StatusBadRequest, result{Result: err.Error()})
		}

		return c.JSON(http.StatusOK, result{Result: "Downloaded!"})
	}
}

func importFromScanner(iSrv services.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		importDetails, err := getImportDetails(c)
		if err != nil {
			return err
		}
		scanner, importPath := importDetails[0], importDetails[1]

		rb, err := os.ReadFile(importPath)
		if err != nil {
			fmt.Printf("error while opening %v for reading: %v", importPath, err)
			return err
		}
		// get the scanner from the url param
		// scanner := strings.ToLower(c.Param("scanner"))
		r := bytes.NewReader(rb)
		// define product and error variable
		var prod *model.Product

		// switch over the scanner name
		switch scanner {
		case "spdx":
			prod, err = iSrv.SPDXImport(r)
		case "composer":
			prod, err = iSrv.ComposerImport(r)
		case "file-hasher":
			prod, err = iSrv.FileHasherImport(r)
		case "scanner":
			prod, err = iSrv.ScannerImport(r)
		default:
			return c.String(
				http.StatusOK,
				fmt.Sprintf("received result file with content length %d, but will not import content, because there is no importer for the scanner '%s'", c.Request().ContentLength, scanner))
		}

		// check error
		if err != nil {
			c.Error(errors.Wrap(err, fmt.Sprintf("unable to perform import for scanner %s", scanner)))
		}

		return c.String(
			http.StatusCreated,
			fmt.Sprintf("successfully parsed content from scanner %s.\nProduct id: %v\nFound %v packages\n", scanner, prod.ID, len(prod.Components)),
		)
	}
}

func scan(iSrv services.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		scanDetails, err := getScanDetails(c)
		if err != nil {
			return err
		}
		result, err := iSrv.Scan(scanDetails)

		if err != nil {
			return c.String(
				http.StatusInternalServerError,
				err.Error(),
			)
		}

		return c.String(
			http.StatusOK,
			result,
		)
	}
}

func checkLicenseCompatibility(srv services.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				result{Result: err.Error()},
			)
		}

		res, err := srv.CheckLicenseCompatibility(id)
		if err != nil {
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
	output := jsonBody["output"]

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

	return []string{importType, importPath}, nil
}

func getJSONRawBody(c echo.Context) (map[string]string, error) {

	jsonBody := make(map[string]string)
	err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	if err != nil {
		return nil, err
	}

	return jsonBody, nil
}

func exportWithType(iSrv services.Service) echo.HandlerFunc {
	return func(c echo.Context) error {

		// get json body
		jsonBody, err := getJSONRawBody(c)
		if err != nil {
			return c.String(
				http.StatusInternalServerError,
				err.Error(),
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
				return c.String(
					http.StatusInternalServerError,
					err.Error(),
				)
			}
			return c.String(
				http.StatusCreated,
				fmt.Sprintf("export path: %v", exportPath),
			)
		case "human-read":
			exportPath, err = iSrv.ReportExport(exportId, exportPath)
			if err != nil {
				return c.String(
					http.StatusInternalServerError,
					err.Error(),
				)
			}
			return c.String(
				http.StatusCreated,
				fmt.Sprintf("export path: %v", exportPath),
			)
		case "compatibility":
			report, exportPath, err := iSrv.CompatibilityExport(exportId, exportPath)
			if err != nil {
				return c.String(
					http.StatusInternalServerError,
					err.Error(),
				)
			}
			fmt.Println(report)
			return c.String(
				http.StatusCreated,
				fmt.Sprintf("export path: %v", exportPath),
			)
		default:
			return c.String(
				http.StatusNotAcceptable,
				"invalid type",
			)
		}
	}
}

func searchSPDX(srv services.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		// get json body
		jsonBody, err := getJSONRawBody(c)
		if err != nil {
			return c.String(
				http.StatusInternalServerError,
				err.Error(),
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
			fmt.Printf("Error while building document: %v\n", err)
			return err
		}

		fmt.Printf("successfully created document and searched for IDs for package %s\n", packageName)

		w, err := os.Create(fileOut)
		if err != nil {
			fmt.Printf("error while opening %v for writing: %v\n", fileOut, err)
			return err
		}
		defer w.Close()

		err = tvsaver.Save2_2(doc, w)
		if err != nil {
			fmt.Printf("error while saving %v: %v", fileOut, err)
			return err
		}

		return c.String(
			http.StatusOK,
			fmt.Sprintf("successfully saved: %v", fileOut),
		)
	}
}
