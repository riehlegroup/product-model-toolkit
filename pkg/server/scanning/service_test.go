// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package scanning

import (
	scanner2 "github.com/osrgroup/product-model-toolkit/pkg/client/commands/scanner"
	"reflect"
	"testing"
)

var dummyT1 = scanner2.Tool{
	Name:      "Licensee",
	Version:   "9.13.0",
	DockerImg: "docker.pkg.github.com/osrgroup/product-model-toolkit/scanner-licensee:9.13.0",
	Cmd:       `/bin/bash -c "licensee detect /input/ --json > /result/result.json"`,
	Results:   []string{"result.json", "result.spdx"},
}

var dummyT2 = scanner2.Tool{
	Name:      "Abc",
	Version:   "1.13.9-beta",
	DockerImg: "docker.pkg.github.com/some-user/some-repo/img:v1.13",
	Cmd:       `/bin/bash -c "Abc scan /in/"`,
	Results:   []string{"my-result.spdx"},
}

var cfg1 = &scanner2.Config{
	Tool:      dummyT1,
	InDir:     "/input",
	ResultDir: "/result",
}

var cfg2 = &scanner2.Config{
	Tool:      dummyT2,
	InDir:     "/in",
	ResultDir: "/out",
}

func Test_execStr(t *testing.T) {
	type args struct {
		cfg *scanner2.Config
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Dummy Tool 1",
			args: args{cfg: cfg1},
			want: `docker run --rm -v /input:/input -v /result:/result docker.pkg.github.com/osrgroup/product-model-toolkit/scanner-licensee:9.13.0 /bin/bash -c "licensee detect /input/ --json > /result/result.json"`,
		},
		{
			name: "Dummy Tool 2",
			args: args{cfg: cfg2},
			want: `docker run --rm -v /in:/input -v /out:/result docker.pkg.github.com/some-user/some-repo/img:v1.13 /bin/bash -c "Abc scan /in/"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := execStr(tt.args.cfg); got != tt.want {
				t.Errorf("execStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_contains(t *testing.T) {
	type args struct {
		slice []string
		val   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "present",
			args: args{[]string{"me"}, "me"},
			want: true,
		},
		{
			name: "present with other",
			args: args{[]string{"me", "other"}, "me"},
			want: true,
		},
		{
			name: "not resent",
			args: args{[]string{"other", "another"}, "me"},
			want: false,
		},
		{
			name: "empty",
			args: args{[]string{}, "me"},
			want: false,
		},
		{
			name: "nil",
			args: args{[]string{}, "me"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := contains(tt.args.slice, tt.args.val); got != tt.want {
				t.Errorf("contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

var fileInfoMock struct {
}

func Test_findFiles(t *testing.T) {
	type args struct {
		names    []fileName
		expected []string
	}
	tests := []struct {
		name string
		args args
		want []fileName
	}{
		{
			name: "should find",
			args: args{names: []fileName{"a.out", "b.out"}, expected: []string{"a.out"}},
			want: []fileName{"a.out"},
		},
		{
			name: "empty names slice",
			args: args{names: []fileName{}, expected: []string{"a.out"}},
			want: []fileName{},
		},
		{
			name: "empty expected slice",
			args: args{names: []fileName{"a.out", "b.out"}, expected: []string{}},
			want: []fileName{},
		},
		{
			name: "more expected than found",
			args: args{names: []fileName{"a.out", "b.out"}, expected: []string{"a.out", "b.out", "c.out"}},
			want: []fileName{"a.out", "b.out"},
		},
		{
			name: "nil",
			args: args{names: nil, expected: nil},
			want: []fileName{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findFiles(tt.args.names, tt.args.expected); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}
