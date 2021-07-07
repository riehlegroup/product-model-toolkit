// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package commands

import (
	"github.com/osrgroup/product-model-toolkit/cnst"
	// "github.com/osrgroup/product-model-toolkit/pkg/server/services"
	"fmt"
)

func RunImport(importType, importPath string) error {

	client := newClient(cnst.ServerBaseURL)

	// log server version with respect to client
	logServerVersion(client)

	url := "http://localhost:8081/api/v1/products/import"

	payload := fmt.Sprintf("{\n\t\"importType\":\"%v\",\n\t\"importPath\":\"%v\"\n}", importType, importPath)

	_, err := client.postData(url, []byte(payload))
	return err
}

