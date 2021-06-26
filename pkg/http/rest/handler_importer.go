// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package rest

import (
	"fmt"
	"github.com/osrgroup/product-model-toolkit/pkg/server/services"
	"net/http"
	"strings"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/osrgroup/product-model-toolkit/model"
	"github.com/pkg/errors"
)

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
			fmt.Sprintf("successfully parsed content from scanner %s.\nProduct id: %v\nFound %v packages\n", scanner,prod.ID, len(prod.Components)),
		)
	}
}

func getJSONRawBody(c echo.Context) (map[string]string, error)  {

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
		if err !=nil {
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
			if err !=nil {
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
			if err !=nil {
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