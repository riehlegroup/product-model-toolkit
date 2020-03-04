package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/osrgroup/product-model-toolkit/pkg/version"
)

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

func handleEntryPoint(c echo.Context) error {
	return c.JSON(http.StatusOK, c.Echo().Routers())
}
