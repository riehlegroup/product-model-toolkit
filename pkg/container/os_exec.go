package container

import (
	"fmt"
	"os/exec"
)

type Executor struct {
	Img       string
	Cmd       string
	InputDir  string
	ResultDir string
	ResultFile string
}

func (e *Executor) Exec() {
	fmt.Println("Execute: ", e.CmdStr())
	out, err := exec.Command("/bin/sh", "-c", e.CmdStr()).CombinedOutput()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("Output: \n%s", out)
}

func (e *Executor) CmdStr() string {
	return fmt.Sprintf("docker run -v %s:/input -v %s:/result %s %s", e.InputDir, e.ResultDir, e.Img, e.Cmd)
}
