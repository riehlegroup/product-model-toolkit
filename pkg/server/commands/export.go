package commands

import (
	// "bytes"
	"encoding/json"
	// "errors"
	"fmt"
	"io/ioutil"

	"github.com/osrgroup/product-model-toolkit/cnst"
	"github.com/osrgroup/product-model-toolkit/model"
	// "github.com/spdx/tools-golang/tvloader"
)

func RunExport(exportId, exportPath string) error {
	// creating a new http client
	client := newClient(cnst.ServerBaseURL)

	// log server version with respect to client
	logServerVersion(client)

	// get the id from the database
	result, err := client.GetProductId(exportId)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return err
	}

	var prod model.Product
	if err := json.Unmarshal(body, &prod); err != nil {
		return err
	}

	// build the spdx file

	// save the spdx file
	
	// return the path of the spdx file

	fmt.Println(prod)
	// return
	return nil
}

