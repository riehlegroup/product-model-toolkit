// SPDX-FileCopyrightText: Cristian Mogildea
//
// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

// ExecResponse struct is used to get the output of an executed command line
type ExecResponse struct {
	StdOut   string
	StdErr   string
	ExitCode int
}

const envDockerUser = "DOCKER_USER"
const envDockerToken = "DOCKER_TOKEN"

// ExecPlugin executes the plugin and returns nil if successful
func ExecPlugin(cfg *Config) error {
	resp, ctx, cli, err := PrepareContainer(cfg)
	if err != nil {
		return err
	}

	err = cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
	if err != nil {
		return err
	}

	err = StartPluginServer(cfg)
	if err != nil {
		return err
	}

	err = ExecAllPluginCmd(ctx, resp.ID, cfg)
	if err != nil {
		return err
	}

	err = StopContainer(resp.ID)
	if err != nil {
		return err
	}

	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			return err
		}
	case <-statusCh:
	}

	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		return err
	}

	stdcopy.StdCopy(os.Stdout, os.Stderr, out)

	return nil
}

// PrepareContainer pulls image from container registry and prepares container for execution
func PrepareContainer(cfg *Config) (container.ContainerCreateCreatedBody, context.Context, *client.Client, error) {
	var resp container.ContainerCreateCreatedBody
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return resp, ctx, cli, err
	}

	authStr, err := GetRegistryAuth()
	if err != nil {
		return resp, ctx, cli, err
	}

	reader, err := cli.ImagePull(ctx, cfg.DockerImg, types.ImagePullOptions{RegistryAuth: authStr})
	if err == nil {
		io.Copy(os.Stdout, reader)
	}
	if err != nil {
		log.Printf("[Plugin agent] Unable to pull image from container registry, got following error: %v\n", err)
	}

	resp, err = containerCreate(ctx, cli, cfg)
	if err != nil {
		return resp, ctx, cli, err
	}

	return resp, ctx, cli, err
}

func containerCreate(ctx context.Context, cli *client.Client, cfg *Config) (container.ContainerCreateCreatedBody, error) {
	return cli.ContainerCreate(ctx,
		&container.Config{
			Image: cfg.DockerImg,
			Cmd:   []string{GetShell(cfg)},
			Tty:   true,
		},
		&container.HostConfig{
			Mounts: []mount.Mount{
				{
					Type:   mount.TypeBind,
					Source: cfg.InDir,
					Target: "/input",
				},
			},
			NetworkMode: "host",
		}, nil, "")
}

// GetRegistryAuth returns authentication string required to pull container from container registry
func GetRegistryAuth() (string, error) {
	user := os.Getenv(envDockerUser)
	token := os.Getenv(envDockerToken)

	if user == "" || token == "" {
		log.Println("[Plugin agent] No authentication credentials provided, please check if environment variables are set")
		return "", errors.New("no authentication credentials provided")
	}

	authConfig := types.AuthConfig{
		Username: user,
		Password: token,
	}

	encodedJSON, err := json.Marshal(authConfig)
	if err != nil {
		return "", err
	}

	authStr := base64.URLEncoding.EncodeToString(encodedJSON)
	return authStr, nil
}

// GetShell returns the Unix shell required to run the command lines
func GetShell(cfg *Config) string {
	return strings.Split(cfg.Cmd, " ")[0]
}

// ExecAllPluginCmd executes all required command lines in the container
func ExecAllPluginCmd(ctx context.Context, containerID string, cfg *Config) error {
	logFile, err := CreateLogFile()
	if err != nil {
		return err
	}

	currentCmd := "echo test"
	expectedOutput := "test\n"
	err = ExecPluginCmd(ctx, containerID, cfg, currentCmd, expectedOutput, true, logFile)
	if err != nil {
		return err
	}

	currentCmd = "mkdir /result"
	expectedOutput = ""
	err = ExecPluginCmd(ctx, containerID, cfg, currentCmd, expectedOutput, false, logFile)
	if err != nil {
		return err
	}

	currentCmd = cfg.Cmd[strings.Index(cfg.Cmd, "-c")+3 : len(cfg.Cmd)]
	expectedOutput = ""
	err = ExecPluginCmd(ctx, containerID, cfg, currentCmd, expectedOutput, false, logFile)
	if err != nil {
		return err
	}

	currentCmd = fmt.Sprintf("for i in /result/*; do curl -F name=%s -F result=@$i http://127.0.0.1:%d/save; done", cfg.Name, GetPortNumber())
	expectedOutput = ""
	err = ExecPluginCmd(ctx, containerID, cfg, currentCmd, expectedOutput, false, logFile)
	if err != nil {
		return err
	}

	log.Printf("[Plugin agent] All commands were executed, check logfile %v for outputs of executed command lines\n", logFile)

	return nil
}

func CreateLogFile() (string, error) {
	file, err := ioutil.TempFile("", time.Now().Format("pmt_container_output_20060102150405_*.log"))
	if err != nil {
		return "", err
	}

	return file.Name(), nil
}

// ExecPluginCmd executes the command line and checks if successful
func ExecPluginCmd(ctx context.Context, containerID string, cfg *Config, cmd string, output string, outputCheck bool, logFile string) error {
	log.Printf("[Plugin agent] Executing command line \"%v\" in container\n", cmd)

	idResponse, err := ExecContainerCmd(ctx, containerID, PrepareCmd(cfg, cmd))
	if err != nil {
		log.Printf("[Plugin agent] Error when executing command line \"%v\" in container\n", cmd)
		return err
	}

	execResponse, err := GetExecResponse(ctx, idResponse)
	if err != nil {
		log.Printf("[Plugin agent] Unable to get output of executed command line \"%v\"\n", cmd)
		return err
	}
	if execResponse.StdOut != "" {
		WriteToLogFile(logFile, cmd, "stdout", execResponse.StdOut)
	}
	if execResponse.StdErr != "" {
		WriteToLogFile(logFile, cmd, "stderr", execResponse.StdErr)
	}

	if outputCheck == true && execResponse.StdOut != output {
		log.Printf("[Plugin agent] Incorrect output of executed command line \"%v\", got \"%v\", but expected \"%v\"\n", cmd, execResponse.StdOut, "test")
		return errors.New("incorrect output of executed command line")
	}

	log.Printf("[Plugin agent] Command line \"%v\" successfully executed\n", cmd)

	return nil
}

// PrepareCmd generates the complete command line that specifies the Unix shell
func PrepareCmd(cfg *Config, cmd string) []string {
	bashCmd := strings.Split(cfg.Cmd[0:strings.Index(cfg.Cmd, "-c")+2], " ")
	return append(bashCmd, cmd)
}

// ExecContainerCmd executes the command line in the specified container
func ExecContainerCmd(ctx context.Context, containerID string, command []string) (types.IDResponse, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return types.IDResponse{}, err
	}

	config := types.ExecConfig{
		AttachStderr: true,
		AttachStdout: true,
		Cmd:          command,
	}

	return cli.ContainerExecCreate(ctx, containerID, config)
}

// GetExecResponse returns the output of the executed command line
func GetExecResponse(ctx context.Context, idResponse types.IDResponse) (ExecResponse, error) {
	var execResponse ExecResponse

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return execResponse, err
	}

	resp, err := cli.ContainerExecAttach(ctx, idResponse.ID, types.ExecStartCheck{})
	if err != nil {
		return execResponse, err
	}

	var outBuf, errBuf bytes.Buffer
	outputDone := make(chan error)

	go func() {
		_, err = stdcopy.StdCopy(&outBuf, &errBuf, resp.Reader)
		outputDone <- err
	}()

	select {
	case err := <-outputDone:
		if err != nil {
			return execResponse, err
		}
		break
	case <-ctx.Done():
		return execResponse, ctx.Err()
	}

	stdout, err := ioutil.ReadAll(&outBuf)
	if err != nil {
		return execResponse, err
	}
	stderr, err := ioutil.ReadAll(&errBuf)
	if err != nil {
		return execResponse, err
	}

	res, err := cli.ContainerExecInspect(ctx, idResponse.ID)
	if err != nil {
		return execResponse, err
	}

	execResponse.StdOut = string(stdout)
	execResponse.StdErr = string(stderr)
	execResponse.ExitCode = res.ExitCode
	return execResponse, nil
}

func WriteToLogFile(logFile string, cmd string, stream string, text string) error {
	src, err := os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer src.Close()

	_, err = src.WriteString(fmt.Sprintf("%s for %s\n%s\n", stream, cmd, text))
	if err != nil {
		return err
	}

	return nil
}

// StopContainer stops the specified container
func StopContainer(containerID string) error {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}

	log.Printf("[Plugin agent] Stopping container %v... \n", containerID[:10])
	err = cli.ContainerStop(ctx, containerID, nil)
	if err != nil {
		return err
	}

	log.Printf("[Plugin agent] Container %v stopped successfully\n", containerID[:10])
	return nil
}
