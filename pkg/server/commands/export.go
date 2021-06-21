package commands

import (
	"github.com/osrgroup/product-model-toolkit/cnst"
)

func RunExport(exportId, exportPath, postPath string) error {
	// creating a new http client
	client := newClient(cnst.ServerBaseURL)

	// log server version with respect to client
	logServerVersion(client)

	// log information

	// execute docker command

	// check error

	// export the data as files

	// return
	return nil
}

