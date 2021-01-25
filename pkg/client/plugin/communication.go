// SPDX-FileCopyrightText: 2021 Cristian Mogildea
//
// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"archive/tar"
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/docker/docker/client"
)

var resultFiles [][]byte

func getResultsFromContainer(cfg *Config, cli *client.Client, ctx context.Context, id string) error {
	for _, s := range cfg.Results {
		err := doGetResultsFromContainer(cli, ctx, id, fmt.Sprintf("/result/%v", s))
		if err != nil {
			return err
		}
	}

	return nil
}

func doGetResultsFromContainer(cli *client.Client, ctx context.Context, id string, path string) error {
	tarStream, _, err := cli.CopyFromContainer(ctx, id, path)
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

	resultFiles = append(resultFiles, buf.Bytes())

	return nil
}

func GetResultFiles() [][]byte {
	return resultFiles
}
