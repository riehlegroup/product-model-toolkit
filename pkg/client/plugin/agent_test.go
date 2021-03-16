// SPDX-FileCopyrightText: 2021 Cristian Mogildea
//
// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"reflect"
	"testing"
)

func Test_prepareCmd(t *testing.T) {
	config1 := &Config{
		Plugin: Plugin{
			Shell: "/bin/bash",
			Cmd:   "test",
		},
	}
	config2 := &Config{
		Plugin: Plugin{
			Shell: "/bin/sh",
			Cmd:   "test -test",
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
		{name: "test#1", args: args{cfg: config1, cmd: config1.Cmd}, want: []string{"/bin/bash", "-c", "test"}},
		{name: "test#2", args: args{cfg: config2, cmd: config2.Cmd}, want: []string{"/bin/sh", "-c", "test -test"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prepareCmd(tt.args.cfg, tt.args.cmd); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prepareCmd() = %v, want %v", got, tt.want)
			}
		})
	}
}
