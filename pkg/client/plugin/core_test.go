// SPDX-FileCopyrightText: 2021 Cristian Mogildea
//
// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"os"
	"testing"
)

func TestGetRemoteRepoAuth(t *testing.T) {
	setAuthEnv()
	authStr, err := getRemoteRepoAuth()
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	expected := "eyJ1c2VybmFtZSI6Im15VXNlciIsInBhc3N3b3JkIjoic2VjcmV0In0="
	if authStr != expected {
		t.Errorf("Expected auth str to be '%s', but got '%s'", expected, authStr)
	}

	unsetAuthEnv()
}

func TestGetRemoteRepoAuth_MissingEnv(t *testing.T) {
	unsetAuthEnv()

	_, err := getRemoteRepoAuth()
	if err == nil {
		t.Errorf("Expected to return error when auth env is not present")
	}
}

func setAuthEnv() {
	os.Setenv(envRemoteRepoUser, "myUser")
	os.Setenv(envRemoteRepoPass, "secret")
}

func unsetAuthEnv() {
	os.Unsetenv(envRemoteRepoUser)
	os.Unsetenv(envRemoteRepoPass)
}
