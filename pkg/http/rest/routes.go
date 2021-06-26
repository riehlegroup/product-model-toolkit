// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package rest

import (

	"github.com/osrgroup/product-model-toolkit/pkg/server/services"

	"github.com/labstack/echo/v4"
)

// Handler handle all request for the given route group.
func Handler(g *echo.Group, srv services.Service) {
	// default routes
	g.GET("/", handleEntryPoint)
	g.GET("/health", handleHealth)

	// product routes
	g.GET("/products", findAllProducts(srv))
	g.GET("/products/:id", findProductByID(srv))
	g.POST("/products/import/:scanner", importFromScanner(srv))

	g.POST("/products/export", exportWithType(srv))

	// // crawler routes
	// g.GET("/crawlers/list", listAllCrawlers(srv))
	// g.POST("/crawlers/license/add", addLicense(srv))
	
	// // diff routes
	// g.POST("/diff/id", diffById(srv))
	// g.POST("/diff/path", diffByPath(srv))
	
	// // export routes
	// g.POST("/export", exportFile(srv))

	// // import routes
	// g.POST("/import", importFile(srv))

	// // merge routes
	// g.POST("/merge", mergeFiles(srv))

	// // search routes
	// g.GET("/search", searchData(srv))

	// // version routes
	// g.GET("/version", handleVersion)
}