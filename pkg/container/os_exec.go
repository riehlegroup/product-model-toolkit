package container

import (
	"fmt"
	"os/exec"
)

// Executor defines what is needed to execute a tool within a container.
type Executor struct {
	Img        string
	Cmd        string
	InputDir   string
	ResultDir  string
	ResultFile string
}

// Exec executes a tool within a Docker container.
func (e *Executor) Exec() {
	fmt.Println("Execute: ", e.CmdStr())
	out, err := exec.Command("/bin/sh", "-c", e.CmdStr()).CombinedOutput()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("Output: \n%s", out)
}

// CmdStr creates a cli string which can be executed by a shell.
func (e *Executor) CmdStr() string {
	return fmt.Sprintf("docker run -v %s:/input -v %s:/result %s %s", e.InputDir, e.ResultDir, e.Img, e.Cmd)
}
