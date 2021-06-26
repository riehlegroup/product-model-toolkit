package commands

import (
	"fmt"
	"github.com/osrgroup/product-model-toolkit/cnst"
)

func RunSearch(name, dir, output string) error {
	// creating a new http client
	client := newClient(cnst.ServerBaseURL)

	// log server version with respect to client
	logServerVersion(client)

	// get the command-line arguments

	url := "http://localhost:8081/api/v1/spdx/search"

	payload := fmt.Sprintf("{\n\t\"name\":\"%v\",\n\t\"dir\":\"%v\",\n\t\"out\":\"%v\"\n}", name, dir, output)

	_, err := client.postData(url, []byte(payload))
	return err
}



