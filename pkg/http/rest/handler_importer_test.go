// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package rest

import (
	"testing"
)

func TestProductLocation(t *testing.T) {
	loc := productLocation("/api/v1/products/something", 42)

	should := "/api/v1/products/42"
	if loc != should {
		t.Errorf("Expected product location to be '%v', but got '%v'", should, loc)
	}
}
