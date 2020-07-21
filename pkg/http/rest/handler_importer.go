// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package rest

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/osrgroup/product-model-toolkit/pkg/importing"
	"github.com/pkg/errors"
)

func importSPDX(iSrv importing.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		r := c.Request().Body

		doc, err := iSrv.SPDX(r)
		if err != nil {
			c.Error(errors.Wrap(err, "Unable to perform SPDX import"))
		}

		return c.String(
			http.StatusOK,
			fmt.Sprintf("Successfully parsed SDPX document.\nFound %v packages", len(doc.Packages)),
		)
	}
}

func importComposer(iSrv importing.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		r := c.Request().Body

		prod, err := iSrv.ComposerRead(r)
		if err != nil {
			c.Error(errors.Wrap(err, "Unable to perform Composer import"))
		}

		loc := productLocation(c.Path(), prod.ID)
		c.Response().Header().Set(echo.HeaderLocation, loc)

		return c.String(
			http.StatusCreated,
			fmt.Sprintf("Successfully parsed Composer JSON.\nFound %v packages\n", len(prod.Components)),
		)
	}
}

func productLocation(path string, id int) string {
	i := strings.LastIndex(path, "/")
	prodPath := path[:i+1]

	return fmt.Sprintf("%s%s", prodPath, strconv.Itoa(id))
}
