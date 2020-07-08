// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package model

import (
	"testing"
)

func TestComponent_ID(t *testing.T) {
	type fields struct {
		Name    string
		Pkg     string
		Version string
	}
	tests := []struct {
		name   string
		fields fields
		want   CmpID
	}{
		{
			name:   "myCmp",
			fields: fields{Name: "myCmp", Pkg: "org.pmt.model", Version: "1.2.3-beta"},
			want:   "org.pmt.model:myCmp:1.2.3-beta",
		},
		{
			name:   "empty component",
			fields: fields{},
			want:   "::",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Component{
				Name:    tt.fields.Name,
				Pkg:     tt.fields.Pkg,
				Version: tt.fields.Version,
			}
			if got := c.ID(); got != tt.want {
				t.Errorf("Component.ID() = %v, want %v", got, tt.want)
			}
		})
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
			name: "empty array",
			args: args{cmps: []Component{}, cid: "abc"},
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
