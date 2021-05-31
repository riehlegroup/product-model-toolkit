// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package rest

import (
	"fmt"
	"github.com/osrgroup/product-model-toolkit/pkg/services/diff"
	importing2 "github.com/osrgroup/product-model-toolkit/pkg/services/importing"
	querying2 "github.com/osrgroup/product-model-toolkit/pkg/services/querying"
	version2 "github.com/osrgroup/product-model-toolkit/pkg/services/version"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

// Handler handle all request for the given route group.
func Handler(g *echo.Group, diff diff.Service, qSrv querying2.Service, iSrv importing2.Service) {
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
	return c.String(http.StatusOK, version2.Name())
}

func handleHealth(c echo.Context) error {
	type status struct {
		Status string `json:"status"`
	}

	return c.JSON(http.StatusOK, status{Status: "UP"})
}

func findAllProducts(q querying2.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		prods, err := q.FindAllProducts()
		if err != nil {
			c.Error(errors.Wrap(err, "Unable to find all products"))
		}

		return c.JSON(http.StatusOK, prods)
	}
}

func findProductByID(q querying2.Service) echo.HandlerFunc {
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
