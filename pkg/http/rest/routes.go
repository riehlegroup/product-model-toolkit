// SPDX-FileCopyrightText: 2022 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package rest

import (
	"github.com/osrgroup/product-model-toolkit/pkg/server/services"

	"github.com/labstack/echo/v4"
)

// Handler handles all request for the given route group.
func Handler(g *echo.Group, srv services.Service) {
	// default routes
	g.GET("/", handleEntryPoint)
	g.GET("/health", handleHealth)
	g.GET("/version", handleVersion)

	// product routes
	g.GET("/products", findAllProducts(srv))
	g.GET("/products/:id", findProductByID(srv))
	g.PUT("/products/:id", updateProductByID(srv))
	g.DELETE("/products/:id", deleteProductByID(srv))
	g.POST("/products/import", importFromScanner(srv))
	g.POST("/products/export", exportWithType(srv))

	// functionalities
	g.POST("/scanner", scan(srv))
	g.POST("/spdx/search", searchSPDX(srv))
	g.GET("/lc/:id", checkLicenseCompatibility(srv))
	g.POST("/download", download(srv))
	g.GET("/downloads", getAllDownloadedRepos(srv))

	g.POST("/diff", getDiffProducts(srv))
}
