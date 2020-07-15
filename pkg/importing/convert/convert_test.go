// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package convert

import "testing"

func TestTrimUTF8prefix(t *testing.T) {
	data := []byte{0xef, 0xbb, 0xbf, 6, 7, 8, 9, 10}

	result := TrimUTF8prefix(data)

	if len(result) != 5 {
		t.Errorf("Expected data without prefix to have length of 5, but got %v", len(result))
	}

	if result[0] != 6 {
		t.Errorf("Expected data byte array start with value %v, but got %v", 6, result[0])
	}
}
