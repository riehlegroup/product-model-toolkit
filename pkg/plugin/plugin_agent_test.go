// SPDX-FileCopyrightText: Cristian Mogildea
//
// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"fmt"
	"os"
	"testing"
)

func TestGetRegistryAuth(t *testing.T) {
	_, err := GetRegistryAuth()
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

func TestGenerateCmd(t *testing.T) {
	wd, _ := os.Getwd()
	plugin := Plugin{
		Name:      "Test",
		Version:   "1.0",
		DockerImg: "",
		Cmd:       "/bin/bash -c ./test",
		Results:   []string{"result1.file", "result2.file", "result3.file"},
	}
	cfg := Config{
		Plugin:    plugin,
		InDir:     wd,
		ResultDir: "/result/dir/",
	}

	result := GenerateCmd(&cfg, 55555)

	// TODO
	fmt.Println(result)
}
