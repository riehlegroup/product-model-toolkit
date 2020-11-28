// SPDX-FileCopyrightText: Cristian Mogildea
//
// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"os"
	"testing"
)

func TestGetRegistryAuth(t *testing.T) {
	setAuthEnv()
	authStr, err := GetRegistryAuth()
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	expected := "eyJ1c2VybmFtZSI6Im15VXNlciIsInBhc3N3b3JkIjoic2VjcmV0In0="
	if authStr != expected {
		t.Errorf("Expected auth str to be '%s', but got '%s'", expected, authStr)
	}

	unsetAuthEnv()
}

func TestGetRegistryAuth_MissingEnv(t *testing.T) {
	unsetAuthEnv()

	_, err := GetRegistryAuth()
	if err == nil {
		t.Errorf("Expected to return error when auth env is not present")
	}
}

func TestGetShell(t *testing.T) {
	cfg := &Config{
		Plugin: Plugin{
			Cmd: "/bin/sh ls",
		},
	}

	sh := GetShell(cfg)
	expected := "/bin/sh"
	if sh != expected {
		t.Errorf("Expected returned shell to be '%s', but got '%s'", expected, sh)
	}
}

func setAuthEnv() {
	os.Setenv(envDockerUser, "myUser")
	os.Setenv(envDockerToken, "secret")
}

func unsetAuthEnv() {
	os.Unsetenv(envDockerUser)
	os.Unsetenv(envDockerToken)
}
