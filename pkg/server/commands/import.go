package commands

import (
	"github.com/osrgroup/product-model-toolkit/cnst"
	// "github.com/osrgroup/product-model-toolkit/pkg/server/services"
	"os"
	"fmt"
)

func RunImport(importType, importPath string) error {
	// creating a new http client
	client := newClient(cnst.ServerBaseURL)

	// log server version with respect to client
	logServerVersion(client)

	// read the file
	// open the first spdx file
	r, err := os.ReadFile(importPath)
	if err != nil {
		fmt.Printf("error while opening %v for reading: %v", importPath, err)
		return err
	}

	url := fmt.Sprintf("%s/products/import/%s", client.baseURL, importType)
	// get the id from the database
	_, err = client.postData(url, r)
	if err != nil {
		return err
	}

	
	

	// return
	return nil
}

