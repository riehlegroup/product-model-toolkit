// SPDX-FileCopyrightText: 2021 Cristian Mogildea
//
// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"bytes"
	"io"
	"reflect"
	"strings"
	"testing"
)

func Test_initializeFilestore(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "1", args: args{length: 1}},
		{name: "5", args: args{length: 5}},
		{name: "10", args: args{length: 10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initializeFilestore(tt.args.length)
			if got := len(resultsFilestore); !reflect.DeepEqual(got, tt.args.length) {
				t.Errorf("len(resultsFilestore) = %v, want %v", got, tt.args.length)
			}
			resultsFilestore = nil
		})
	}
}

func Test_initializeFilestore_ZeroLength(t *testing.T) {
	err := initializeFilestore(0)
	if err == nil {
		t.Errorf("Expected to return error when length is zero")
	}
}

func Test_saveResultFile(t *testing.T) {
	err := initializeFilestore(5)
	if err != nil {
		t.Errorf("Could not initialize filestore: %v", err)
	}

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, strings.NewReader("test")); err != nil {
		t.Errorf("Could not copy to buffer: %v", err)
	}

	err = saveResultFile(0, buf.Bytes())
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	err = saveResultFile(2, buf.Bytes())
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	err = saveResultFile(4, buf.Bytes())
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

func Test_saveResultFile_IdOutOfBounds(t *testing.T) {
	err := initializeFilestore(5)
	if err != nil {
		t.Errorf("Could not initialize filestore: %v", err)
	}

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, strings.NewReader("test")); err != nil {
		t.Errorf("Could not copy to buffer: %v", err)
	}

	err = saveResultFile(5, buf.Bytes())
	if err == nil {
		t.Errorf("Expected to return error when id out of bounds")
	}
}
