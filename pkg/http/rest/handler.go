// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/osrgroup/product-model-toolkit/pkg/querying"
	"github.com/osrgroup/product-model-toolkit/pkg/version"
)

// Handler handle all request for the given route group.
func Handler(g *echo.Group, q querying.Service) {
	g.GET("/", handleEntryPoint)
	g.GET("/version", handleVersion)
	g.GET("/health", handleHealth)

	g.GET("/products", findAllProducts(q))
	g.GET("/products/:id", findProductByID(q))
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

	up := status{Status: "UP"}
	return c.JSON(http.StatusOK, up)
}

func findAllProducts(q querying.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		prods, err := q.FindAllProducts()
		if err != nil {
			c.Error(err)
		}

		return c.JSON(http.StatusOK, prods)
	}
}

func findProductByID(q querying.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		prod, err := q.FindProductByID(id)
		if err != nil {
			c.Error(err)
		}

		return c.JSON(http.StatusOK, prod)
	}
}
