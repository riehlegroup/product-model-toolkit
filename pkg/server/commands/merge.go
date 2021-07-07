// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package commands

import (
	"github.com/osrgroup/product-model-toolkit/cnst"
)

func RunMerge(first, second, output string) error {
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


