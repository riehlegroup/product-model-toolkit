// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package rest

import (
	"testing"

	"github.com/labstack/echo/v4"
)

func TestInstance_Addr(t *testing.T) {
	e := echo.New()
	type fields struct {
		httpSrv *echo.Echo
		addr    string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "with port 1234",
			fields: fields{httpSrv: e, addr: "1234"},
			want:   "1234",
		},
		{
			name:   "with port 8080",
			fields: fields{httpSrv: e, addr: "8080"},
			want:   "8080",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &Instance{
				httpSrv: tt.fields.httpSrv,
				addr:    tt.fields.addr,
			}
			if got := srv.Addr(); got != tt.want {
				t.Errorf("Instance.Addr() = %v, want %v", got, tt.want)
			}
		})
	}
}
