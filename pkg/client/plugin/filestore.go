// SPDX-FileCopyrightText: 2021 Cristian Mogildea
//
// SPDX-License-Identifier: Apache-2.0

package plugin

import "errors"

type filestore struct {
	results [][]byte
}

var resultsFilestore []filestore

func initializeFilestore(length int) error {
	if length == 0 {
		return errors.New("length cannot be zero")
	}
	if len(resultsFilestore) == 0 {
		resultsFilestore = []filestore{{results: *new([][]byte)}}

		for i := 0; i < length-1; i++ {
			resultsFilestore = append(resultsFilestore, filestore{results: *new([][]byte)})
		}
	}

	return nil
}

func saveResultFile(id int, bytes []byte) error {
	if id > len(resultsFilestore)-1 {
		return errors.New("id out of bounds")
	}
	resultsFilestore[id].results = append(resultsFilestore[id].results, bytes)

	return nil
}

func getResultFiles(id int) [][]byte {
	return resultsFilestore[id].results
}
