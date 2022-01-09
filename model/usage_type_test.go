// SPDX-FileCopyrightText: 2022 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package model

import (
	"testing"
)

func TestUsageType_isValid(t *testing.T) {
	tests := []struct {
		name    string
		ut      UsageType
		wantErr bool
	}{
		{
			name:    "OnPremise",
			ut:      UsageType("on-premise"),
			wantErr: false,
		},
		{
			name:    "CloudService",
			ut:      UsageType("cloud-service"),
			wantErr: false,
		},
		{
			name:    "Library",
			ut:      UsageType("library"),
			wantErr: false,
		},
		{
			name:    "Internal",
			ut:      UsageType("internal"),
			wantErr: false,
		},
		{
			name:    "Unknown",
			ut:      UsageType("unknown usage type"),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.ut.isValid(); (err != nil) != tt.wantErr {
				t.Errorf("UsageType.isValid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
