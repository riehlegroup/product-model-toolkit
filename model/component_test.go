// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package model

import (
	"testing"
)

func TestID(t *testing.T) {
	c := &Component{Name: "Product", Pkg: "org.pmt.model", Version: "1.2.3-beta"}

	result := string(c.ID())

	expect := "org.pmt.model:Product:1.2.3-beta"

	if result != expect {
		t.Errorf("Expected component ID to be '%v', but got '%v'.", expect, result)
	}
}

func TestContainsComp(t *testing.T) {
	type args struct {
		cmps []Component
		cid  CmpID
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should find CompA",
			args: args{cmps: dummyCmps(), cid: "org.test:CompA:1.0.0"},
			want: true,
		},
		{
			name: "should find Comp_B",
			args: args{cmps: dummyCmps(), cid: "com.pmt:Comp_B:0.0.1"},
			want: true,
		},
		{
			name: "should find Comp C",
			args: args{cmps: dummyCmps(), cid: "org.pmt.model:Comp C:0.1.0-beta"},
			want: true,
		},
		{
			name: "should find Comp-D",
			args: args{cmps: dummyCmps(), cid: "pmt:Comp-D:10.134.4"},
			want: true,
		},
		{
			name: "empty string",
			args: args{cmps: dummyCmps(), cid: ""},
			want: false,
		},
		{
			name: "wrong package name",
			args: args{cmps: dummyCmps(), cid: "org.test2:CompA:1.0.0"},
			want: false,
		},
		{
			name: "wrong version number",
			args: args{cmps: dummyCmps(), cid: "org.test:CompA:1.0.1"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsComp(tt.args.cmps, tt.args.cid); got != tt.want {
				t.Errorf("ContainsComp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func dummyCmps() []Component {
	return []Component{
		{
			Name:    "CompA",
			Pkg:     "org.test",
			Version: "1.0.0",
		},
		{
			Name:    "Comp_B",
			Pkg:     "com.pmt",
			Version: "0.0.1",
		},
		{
			Name:    "Comp C",
			Pkg:     "org.pmt.model",
			Version: "0.1.0-beta",
		},
		{
			Name:    "Comp-D",
			Pkg:     "pmt",
			Version: "10.134.4",
		},
	}
}
