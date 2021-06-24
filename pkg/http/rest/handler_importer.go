// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package rest

import (
	"fmt"
	"github.com/osrgroup/product-model-toolkit/pkg/server/services"
	"net/http"
	"strings"

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
			prod, err = iSrv.SPDX(r)
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
			fmt.Sprintf("successfully parsed content from scanner %s.\n Found %v packages\n", scanner, len(prod.Components)),
		)
	}
}

// func productLocation(path string, id int) string {
// 	i := strings.LastIndex(path, "/")
// 	prodPath := path[:i+1]

// 	return fmt.Sprintf("%s%s", prodPath, strconv.Itoa(id))
// }
