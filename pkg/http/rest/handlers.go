// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package rest

import (
	"os/exec"
	"encoding/json"
	"fmt"
	"github.com/osrgroup/product-model-toolkit/model"
	"github.com/osrgroup/product-model-toolkit/pkg/server/services"
	"github.com/pkg/errors"
	"github.com/spdx/tools-golang/idsearcher"
	"github.com/spdx/tools-golang/tvsaver"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

func handleEntryPoint(c echo.Context) error {
	return c.JSON(http.StatusOK, c.Echo().Routers())
}

func handleVersion(c echo.Context) error {
	return c.String(http.StatusOK, "1.0.0")
}

func handleHealth(c echo.Context) error {
	type status struct {
		Status string `json:"status"`
	}

	return c.JSON(http.StatusOK, status{Status: "UP"})
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
			c.String(
				http.StatusNotFound,
				fmt.Sprintf("unable fo find product with ID %v", id))
		}

		return c.JSON(http.StatusOK, prod)
	}
}

func importFromScanner(iSrv services.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		// get the scanner from the url param
		scanner := strings.ToLower(c.Param("scanner"))

		// read request body
		r := c.Request().Body

		// define product and error variable
		var prod *model.Product
		var err error

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
		
		// get json body
		jsonBody, err := getJSONRawBody(c)
		if err != nil {
			return c.String(
				http.StatusInternalServerError,
				err.Error(),
			)
		}

		// read data
		scannerName := jsonBody["scannerName"]
		source := jsonBody["source"]
		output := jsonBody["output"]

		// switch over the scanner name
		switch scannerName {
		case "phpscanner":

		dockerImageName := "docker.pkg.github.com/osrgroup/product-model-toolkit/php-scanner:1.0.0" // TODO

		fmt.Println(source)
		if source == "." {
			source = "$PWD"
		}
		if output == "." {
			output = "$PWD"
		}

		// create the dockerCmd from input values
		dockerCmd := fmt.Sprintf("sudo docker run "+
			"-v %v:/source "+
			"-v %v:/output %v",
			source, output, dockerImageName)

		// log information
		fmt.Println("Running crawler")

		// execute docker command
		fmt.Println("Executing the docker command ...")

		// print the docker command
		fmt.Println(dockerCmd)

		// executing the command
		_, err := exec.Command("/bin/sh", "-c", dockerCmd).CombinedOutput()
		// check error
		if err != nil {
			return err
		}

		fmt.Println("Crawling licenses successfully completed")
		fmt.Printf("The output path: %v\n", output)
		return c.String(
			http.StatusOK,
			fmt.Sprintf("The output path: %v\n", output),
		)

		// case "human-read":
			// fmt.Println("inja1")
			// exportPath, err = iSrv.ReportExport(exportId, exportPath)
			// if err != nil {
			// 	return c.String(
			// 		http.StatusInternalServerError,
			// 		err.Error(),
			// 	)
			// }
			// return c.String(
			// 	http.StatusCreated,
			// 	fmt.Sprintf("export path: %v", exportPath),
			// )
			case "licensee":
					
		// get json body
		jsonBody, err := getJSONRawBody(c)
		if err != nil {
			return c.String(
				http.StatusInternalServerError,
				err.Error(),
			)
		}

		// read data
		scannerName := jsonBody["scannerName"]
		source := jsonBody["source"]
		output := jsonBody["output"]

		// switch over the scanner name
		switch scannerName {
		case "phpscanner":

		dockerImageName := "docker.pkg.github.com/osrgroup/product-model-toolkit/php-scanner:1.0.0" // TODO

		fmt.Println(source)
		if source == "." {
			source = "$PWD"
		}
		if output == "." {
			output = "$PWD"
		}

		// create the dockerCmd from input values
		dockerCmd := fmt.Sprintf("sudo docker run "+
			"-v %v:/source "+
			"-v %v:/output %v",
			source, output, dockerImageName)

		// log information
		fmt.Println("Running crawler")

		// execute docker command
		fmt.Println("Executing the docker command ...")

		// print the docker command
		fmt.Println(dockerCmd)

		// executing the command
		_, err := exec.Command("/bin/sh", "-c", dockerCmd).CombinedOutput()
		// check error
		if err != nil {
			return err
		}

		fmt.Println("Crawling licenses successfully completed")
		fmt.Printf("The output path: %v\n", output)
		return c.String(
			http.StatusOK,
			fmt.Sprintf("The output path: %v\n", output),
		)


		default:
			return c.String(
				http.StatusNotAcceptable,
				"file received but couldn't accept it",
			)
		}

	}
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
			fmt.Println("inja1")
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
		default:
			return c.String(
				http.StatusNotAcceptable,
				"file received but couldn't accept it",
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