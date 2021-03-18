// SPDX-FileCopyrightText: 2021 Cristian Mogildea
//
// SPDX-License-Identifier: Apache-2.0

package plugin

type filestore struct {
	results [][]byte
}

var resultsFilestore []filestore

func initializeFilestore(length int) {
	if len(resultsFilestore) == 0 {
		resultsFilestore = []filestore{{results: *new([][]byte)}}

		for i := 0; i < length-1; i++ {
			resultsFilestore = append(resultsFilestore, filestore{results: *new([][]byte)})
		}
	}
}

func saveResultFile(id int, bytes []byte) {
	resultsFilestore[id].results = append(resultsFilestore[id].results, bytes)
}

func getResultFiles(id int) [][]byte {
	return resultsFilestore[id].results
}
