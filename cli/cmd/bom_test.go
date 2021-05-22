package cmd

import (
	product "pmt/bom/proto"
	"testing"
)

func Test_normaliseTypeValue(t *testing.T) {
	type args struct {
		typeValue string
	}
	tests := []struct {
		name string
		args args
		want product.InputType
	}{
		{
			name: "test_1",
			args: args{
				typeValue: "Error",
			},
			want: product.InputType_SPDX,
		},
		{
			name: "test_2",
			args: args{
				typeValue: "1",
			},
			want: product.InputType_HUMAN,
		},
		{
			name: "test_3",
			args: args{
				typeValue: "2",
			},
			want: product.InputType_CUSTOM,
		},
		{
			name: "test_4",
			args: args{
				typeValue: "10",
			},
			want: product.InputType_SPDX,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := normaliseTypeValue(tt.args.typeValue); got != tt.want {
				t.Errorf("normaliseTypeValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createBomWithType(t *testing.T) {
	type args struct {
		path      string
		typeValue string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test_1",
			args: args{
				path: ".",
				typeValue: "0",
			},
			wantErr: false,
		},
		{
			name: "test_2",
			args: args{
				path: ".",
				typeValue: "1",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := createBomWithType(tt.args.path, tt.args.typeValue); (err != nil) != tt.wantErr {
				t.Errorf("createBomWithType() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}