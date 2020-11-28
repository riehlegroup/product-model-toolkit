// SPDX-FileCopyrightText: Cristian Mogildea
//
// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/labstack/echo"
)

var listener net.Listener

var pluginCfg *Config

var results []multipart.File

func StartPluginServer(cfg *Config) error {
	if listener != nil {
		return errors.New("plugin server already started")
	}

	pluginCfg = cfg

	l, err := net.Listen("tcp", ":0")
	if err != nil {
		return err
	}
	listener = l

	server := echo.New()
	server.POST("/save", SaveResult)
	server.Listener = l

	go func() {
		server.Logger.Fatal(server.Start(""))
	}()

	log.Println("[Plugin server] Waiting 2 seconds for REST API to start")
	time.Sleep(2 * time.Second)

	return nil
}

func GetPortNumber() int {
	return listener.Addr().(*net.TCPAddr).Port
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

// GetResults returns all result files as a list
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
