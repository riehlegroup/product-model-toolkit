// SPDX-FileCopyrightText: Cristian Mogildea
//
// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"testing"
)

func TestGetRegistryAuth(t *testing.T) {
	_, err := GetRegistryAuth()
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}
