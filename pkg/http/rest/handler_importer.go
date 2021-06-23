// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package rest

import (
	"fmt"
	"github.com/osrgroup/product-model-toolkit/pkg/server/services"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/osrgroup/product-model-toolkit/model"
	"github.com/pkg/errors"
)

func importFromScanner(iSrv services.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		scanner := strings.ToLower(c.Param("scanner"))
		r := c.Request().Body

		// if scanner == "scancode" || scanner == "spdx" {
		// 	doc, err := iSrv.SPDX(r)
		// 	if err != nil {
		// 		c.Error(errors.Wrap(err, "unable to perform SPDX import"))
		// 	}
			
		// 	return c.String(
		// 		http.StatusOK,
		// 		fmt.Sprintf("successfully parsed SPDX document.\nFound %v packages", len(doc.Packages)),
				
		// 	)
		// }

		var prod *model.Product
		var err error
		switch scanner {
		case "spdx":
			prod, err = iSrv.SPDX(r)
		case "composer":
			prod, err = iSrv.ComposerImport(r)
		case "file-hasher":
			prod, err = iSrv.FileHasherImport(r)
		default:
			return c.String(
				http.StatusOK,
				fmt.Sprintf("Received result file with content length %d, but will not import content, because there is no importer for the scanner '%s'", c.Request().ContentLength, scanner))
		}

		if err != nil {
			c.Error(errors.Wrap(err, fmt.Sprintf("Unable to perform import for scanner %s", scanner)))
		}

		loc := productLocation(c.Path(), prod.ID)
		c.Response().Header().Set(echo.HeaderLocation, loc)

		return c.String(
			http.StatusCreated,
			fmt.Sprintf("Successfully parsed content from scanner %s.\n Found %v packages\n", scanner, len(prod.Components)),
		)
	}
}

func productLocation(path string, id int) string {
	i := strings.LastIndex(path, "/")
	prodPath := path[:i+1]

	return fmt.Sprintf("%s%s", prodPath, strconv.Itoa(id))
}
