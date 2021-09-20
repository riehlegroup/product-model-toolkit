// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"testing"
)

func Test_introScreen(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "introScreen",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			introScreen()
		})
	}
}

func Test_callscanner(t *testing.T) {
	type args struct {
		name   string
		source string
		output string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "callscanner",
			args: args{
				name:   "test",
				source: "test",
				output: "test",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := callscanner(tt.args.name, tt.args.source, tt.args.output); (err != nil) != tt.wantErr {
				t.Errorf("callscanner() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
