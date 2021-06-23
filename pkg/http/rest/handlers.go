// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package rest

import (
	"fmt"
	"net/http"
	"strconv"

	// "github.com/osrgroup/product-model-toolkit/pkg/services/querying"
	// "github.com/osrgroup/product-model-toolkit/pkg/services/version"
	"github.com/osrgroup/product-model-toolkit/pkg/server/services"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
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

// findAllLicenses
// func findAllLicenses(srv services.Service) echo.HandlerFunc {
	// return func(c echo.Context) error {
		// licenses, err := srv.FindAllLicenses()
		// if err != nil {
			// c.Error(errors.Wrap(err, "unable to find all licenses"))
		// }
		// return c.JSON(http.StatusOK, licenses)
	// }
// }

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
