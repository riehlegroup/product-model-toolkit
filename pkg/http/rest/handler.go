// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package rest

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/osrgroup/product-model-toolkit/pkg/importing"
	"github.com/osrgroup/product-model-toolkit/pkg/querying"
	"github.com/osrgroup/product-model-toolkit/pkg/version"
	"github.com/pkg/errors"
)

// Handler handle all request for the given route group.
func Handler(g *echo.Group, qSrv querying.Service, iSrv importing.Service) {
	g.GET("/", handleEntryPoint)
	g.GET("/version", handleVersion)
	g.GET("/health", handleHealth)

	g.GET("/products", findAllProducts(qSrv))
	g.GET("/products/:id", findProductByID(qSrv))
	g.POST("/products/import/:scanner", importFromScanner(iSrv))
}

func handleEntryPoint(c echo.Context) error {
	return c.JSON(http.StatusOK, c.Echo().Routers())
}

func handleVersion(c echo.Context) error {
	return c.String(http.StatusOK, version.Name())
}

func handleHealth(c echo.Context) error {
	type status struct {
		Status string `json:"status"`
	}

	return c.JSON(http.StatusOK, status{Status: "UP"})
}

func findAllProducts(q querying.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		prods, err := q.FindAllProducts()
		if err != nil {
			c.Error(errors.Wrap(err, "Unable to find all products"))
		}

		return c.JSON(http.StatusOK, prods)
	}
}

func findProductByID(q querying.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.Error(errors.Wrap(err, fmt.Sprintf("Unable to convert query param id with value '%v' to int", idStr)))
		}

		prod, err := q.FindProductByID(id)
		if err != nil {
			c.String(
				http.StatusNotFound,
				fmt.Sprintf("Unable fo find product with ID %v", id))
		}

		return c.JSON(http.StatusOK, prod)
	}
}
