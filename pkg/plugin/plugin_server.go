// SPDX-FileCopyrightText: Cristian Mogildea
//
// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"fmt"
	"github.com/labstack/echo"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

var ResultDir string

func StartServer(cfg *Config) {
	ResultDir = cfg.ResultDir
	server := echo.New()
	server.POST("/save", SaveResult)
	server.Logger.Fatal(server.Start(":8082"))
}

func SaveResult(c echo.Context) error {
	name := c.FormValue("name")
	result, err := c.FormFile("result")
	if err != nil {
		return err
	}

	src, err := result.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(filepath.Join(ResultDir, filepath.Base(result.Filename)))
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, fmt.Sprintf("<b>Message received from %s</b>", name))
}
