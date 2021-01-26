// SPDX-FileCopyrightText: 2021 Cristian Mogildea
//
// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/labstack/echo"
)

var listener net.Listener

var pluginCfg *Config

var results [][]byte

func startPluginServer(cfg *Config) error {
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
	server.POST("/save", saveResult)
	server.Listener = l

	go func() {
		server.Logger.Fatal(server.Start(""))
	}()

	log.Println("[Plugin server] Waiting 2 seconds for REST API to start")
	time.Sleep(2 * time.Second)

	return nil
}

func getPortNumber() int {
	return listener.Addr().(*net.TCPAddr).Port
}

// saveResult saves the received result file in a list
func saveResult(c echo.Context) error {
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

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, src); err != nil {
		return err
	}
	results = append(results, buf.Bytes())

	err = writeFile(buf.Bytes(), result.Filename)
	if err != nil {
		return err
	}

	return c.HTML(http.StatusOK, fmt.Sprintf("Result file %s received from %s\n", result.Filename, name))
}

// GetResults returns all result files as a list
func GetResults() [][]byte {
	return results
}

// writeFile saves the file locally to the specified path
func writeFile(fileContent []byte, filename string) error {
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

	if _, err = io.Copy(dst, bytes.NewReader(fileContent)); err != nil {
		return err
	}

	return nil
}
