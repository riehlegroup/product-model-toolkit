package container

import "testing"

func TestExecutor_CmdStr(t *testing.T) {
	type fields struct {
		Img        string
		Cmd        string
		InputDir   string
		ResultDir  string
		ResultFile string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "tool1",
			fields: fields{Img: "myImage", Cmd: "/bin/sh", InputDir: "/input1", ResultDir: "/result1"},
			want:   "docker run -v /input1:/input -v /result1:/result myImage /bin/sh",
		},
		{
			name:   "licensee",
			fields: fields{Img: "myRepo/licensee:9.13.0", Cmd: "licensee --version", InputDir: "/input", ResultDir: "/result"},
			want:   "docker run -v /input:/input -v /result:/result myRepo/licensee:9.13.0 licensee --version",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Executor{
				Img:        tt.fields.Img,
				Cmd:        tt.fields.Cmd,
				InputDir:   tt.fields.InputDir,
				ResultDir:  tt.fields.ResultDir,
				ResultFile: tt.fields.ResultFile,
			}
			if got := e.CmdStr(); got != tt.want {
				t.Errorf("Executor.CmdStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
