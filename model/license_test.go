// SPDX-FileCopyrightText: 2022 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package model

import "testing"

func TestLicense_toString(t *testing.T) {
	type fields struct {
		SPDXID           string
		DeclaredLicense  string
		ConcludedLicense string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "should return SPDX id",
			fields: fields{SPDXID: "MIT", DeclaredLicense: "", ConcludedLicense: ""},
			want:   "MIT",
		},
		{
			name:   "should return declared license",
			fields: fields{SPDXID: "", DeclaredLicense: "Apache", ConcludedLicense: ""},
			want:   "Apache",
		},
		{
			name:   "should return concluded license",
			fields: fields{SPDXID: "", DeclaredLicense: "Apache", ConcludedLicense: "Apache-2.0"},
			want:   "Apache-2.0",
		},
		{
			name:   "should return empty string",
			fields: fields{SPDXID: "", DeclaredLicense: "", ConcludedLicense: ""},
			want:   "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &License{
				SPDXID:           tt.fields.SPDXID,
				DeclaredLicense:  tt.fields.DeclaredLicense,
				ConcludedLicense: tt.fields.ConcludedLicense,
			}
			if got := l.toString(); got != tt.want {
				t.Errorf("License.toString() = %v, want %v", got, tt.want)
			}
		})
	}
}
