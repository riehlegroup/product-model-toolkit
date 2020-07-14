// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package rest

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/osrgroup/product-model-toolkit/pkg/importing"
)

func importSPDX(iSrv importing.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		r := c.Request().Body

		doc, err := iSrv.SPDX(r)
		if err != nil {
			c.Error(err)
		}

		msg := fmt.Sprintf("Successfully parsed SDPX document.\nFound %v packages", len(doc.Packages))
		return c.String(http.StatusOK, msg)
	}
}

func importComposer(iSrv importing.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		r := c.Request().Body

		prod, err := iSrv.Composer(r)
		if err != nil {
			c.Error(err)
		}

		msg := fmt.Sprintf("Successfully parsed Composer JSON.\nFound %v packages", len(prod.Components))
		return c.String(http.StatusOK, msg)
	}
}
