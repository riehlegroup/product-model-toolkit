// SPDX-FileCopyrightText: 2021 Cristian Mogildea
//
// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"os"
	"reflect"
	"testing"
)

func TestGetRegistryAuth(t *testing.T) {
	setAuthEnv()
	authStr, err := getRegistryAuth()
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

	_, err := getRegistryAuth()
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

	sh := getShell(cfg)
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

func Test_prepareCmd(t *testing.T) {
	config1 := &Config{
		Plugin: Plugin{
			Cmd: "/bin/sh -c test",
		},
	}
	config2 := &Config{
		Plugin: Plugin{
			Cmd: "/bin/bash -c test",
		},
	}

	type args struct {
		cfg *Config
		cmd string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "test#1", args: args{cfg: config1, cmd: "test"}, want: []string{"/bin/sh", "-c", "test"}},
		{name: "test#2", args: args{cfg: config2, cmd: "test"}, want: []string{"/bin/bash", "-c", "test"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prepareCmd(tt.args.cfg, tt.args.cmd); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prepareCmd() = %v, want %v", got, tt.want)
			}
		})
	}
}
