package commands

import (
	"fmt"
	"github.com/osrgroup/product-model-toolkit/cnst"
)

func RunExport(exportId, exportType, exportPath string) error {
	// creating a new http client
	client := newClient(cnst.ServerBaseURL)

	// log server version with respect to client
	logServerVersion(client)

	url := "http://localhost:8081/api/v1/products/export"

	payload := fmt.Sprintf("{\n\t\"exportId\":\"%v\",\n\t\"exportPath\":\"%v\",\n\t\"exportType\":\"%v\"\n}", exportId, exportPath, exportType)

	_, err := client.postData(url, []byte(payload))
	return err
}

