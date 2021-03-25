// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package scanning

import (
	"bytes"
	"fmt"
	"log"

	"github.com/osrgroup/product-model-toolkit/pkg/client/http/rest"
)

var restClient *rest.Client

func LogServerVersion(c *rest.Client) {
	restClient = c

	v, err := c.GetServerVersion()
	if err != nil {
		log.Printf("[Client (REST)] Unable to read PMT server version: %s", err)
		return
	}

	log.Printf("[Client (REST)] PMT server version: %s", v)
}

func SendResults(files [][]byte, pluginName string) {
	url := fmt.Sprintf("/products/import/%s", pluginName)

	for _, f := range files {
		loc, err := restClient.PostResult(url, bytes.NewReader(f))
		if err != nil {
			log.Printf("[Client] Error while sending %s results to server [%s]: %s", pluginName, url, err)
			return
		}

		log.Printf("[Client] Successfully sent %s results to server. Path to created resource: %s", pluginName, loc)
	}
}
