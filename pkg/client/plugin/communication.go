// SPDX-FileCopyrightText: 2021 Cristian Mogildea
//
// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"archive/tar"
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/docker/docker/client"
	"github.com/labstack/echo"
)

var portNr = -1

func getResultsFromContainer(cfg *Config, cli *client.Client, ctx context.Context, id string) error {
	for _, s := range cfg.Results {
		err := doGetResultsFromContainer(cfg, cli, ctx, id, s)
		if err != nil {
			return err
		}
	}

	return nil
}

func doGetResultsFromContainer(cfg *Config, cli *client.Client, ctx context.Context, id string, file string) error {
	tarStream, _, err := cli.CopyFromContainer(ctx, id, fmt.Sprintf("/result/%v", file))
	if err != nil {
		return err
	}

	tr := tar.NewReader(tarStream)
	if _, err := tr.Next(); err != nil {
		return err
	}

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, tr); err != nil {
		return err
	}

	resultsFilestore[cfg.Id].results = append(resultsFilestore[cfg.Id].results, buf.Bytes())

	if coreConfigValues.SaveResultsLocally == true {
		err := writeFile(cfg.Name, buf.Bytes(), file)
		if err != nil {
			return err
		}
	}

	return nil
}

func startServer() error {
	if portNr != -1 {
		return nil
	}

	l, err := net.Listen("tcp", ":0")
	if err != nil {
		return err
	}
	portNr = l.Addr().(*net.TCPAddr).Port

	server := echo.New()
	server.POST("/save", saveResultFile)
	server.Listener = l

	go func() {
		server.Logger.Fatal(server.Start(""))
	}()

	log.Println("[Server] Waiting 2 seconds for REST API to start")
	time.Sleep(2 * time.Second)

	return nil
}

func getPortNumber() int {
	return portNr
}

// saveResultFile saves received result file into resultsFilestore
func saveResultFile(c echo.Context) error {
	name := c.FormValue("name")
	id, err := strconv.Atoi(c.FormValue("id"))
	if err != nil {
		return err
	}
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

	resultsFilestore[id].results = append(resultsFilestore[id].results, buf.Bytes())

	if coreConfigValues.SaveResultsLocally == true {
		err = writeFile(name, buf.Bytes(), result.Filename)
		if err != nil {
			return err
		}
	}

	return c.HTML(http.StatusOK, fmt.Sprintf("Result file %s received from %s\n", result.Filename, name))
}

// writeFile saves file locally
func writeFile(pluginName string, fileContent []byte, filename string) error {
	if coreConfigValues.PathDirResults == "" {
		return errors.New("cannot save file locally, path to directory unspecified")
	}

	dst, err := os.Create(filepath.Join(coreConfigValues.PathDirResults, fmt.Sprintf("%v_%v", pluginName, filepath.Base(filename))))
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, bytes.NewReader(fileContent)); err != nil {
		return err
	}

	return nil
}
