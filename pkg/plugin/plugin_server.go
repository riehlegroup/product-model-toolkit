// SPDX-FileCopyrightText: Cristian Mogildea
//
// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo"
)

var pluginCfg *Config

var results []multipart.File

func StartServer(cfg *Config) {
	pluginCfg = cfg
	server := echo.New()
	server.POST("/save", SaveResult)
	server.Logger.Fatal(server.Start(":8082"))
}

// SaveResult saves the received result file in a list
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

	results = append(results, src)

	err = WriteFile(src, result.Filename)
	if err != nil {
		return err
	}

	return c.HTML(http.StatusOK, fmt.Sprintf("Result file %s received from %s\n", result.Filename, name))
}

// GetResults returns all result files in a list
func GetResults() []multipart.File {
	return results
}

// WriteFile saves the file locally to the specified path
func WriteFile(file multipart.File, filename string) error {
	if _, err := os.Stat(pluginCfg.ResultDir); os.IsNotExist(err) {
		err := os.Mkdir(pluginCfg.ResultDir, 0755)
		if err != nil {
			return err
		}
	}

	dst, err := os.Create(filepath.Join(pluginCfg.ResultDir, filepath.Base(filename)))
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, file); err != nil {
		return err
	}

	return nil
}
