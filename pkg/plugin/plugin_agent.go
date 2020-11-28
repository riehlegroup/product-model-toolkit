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
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// ExecResponse struct is used to get the output of an executed command line
type ExecResponse struct {
	StdOut   string
	StdErr   string
	ExitCode int
}

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
	if err != nil {
		return resp, ctx, cli, err
	}
	io.Copy(os.Stdout, reader)

	resp, err = cli.ContainerCreate(ctx,
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
	if err != nil {
		return resp, ctx, cli, err
	}

	return resp, ctx, cli, err
}

// GetRegistryAuth returns authentication string required to pull container from container registry
func GetRegistryAuth() (string, error) {
	var authStr string
	if os.Getenv("USER") != "" && os.Getenv("TOKEN") != "" {
		authConfig := types.AuthConfig{
			Username: os.Getenv("USER"),
			Password: os.Getenv("TOKEN"),
		}
		encodedJSON, err := json.Marshal(authConfig)
		if err != nil {
			return "", err
		}
		authStr = base64.URLEncoding.EncodeToString(encodedJSON)
		return authStr, nil
	}

	log.Println("[Plugin agent] No authentication credentials provided, please check if environment variables are set")
	return "", errors.New("no authentication credentials provided")
}

// GetShell returns the Unix shell required to run the command lines
func GetShell(cfg *Config) string {
	return strings.Split(cfg.Cmd, " ")[0]
}

// ExecAllPluginCmd executes all required command lines in the container
func ExecAllPluginCmd(ctx context.Context, containerId string, cfg *Config) error {
	currentCmd := "echo test"
	currentOutput := "test"
	err := ExecPluginCmd(ctx, containerId, cfg, currentCmd, currentOutput, true)
	if err != nil {
		return err
	}

	currentCmd = "mkdir /result"
	currentOutput = ""
	err = ExecPluginCmd(ctx, containerId, cfg, currentCmd, currentOutput, false)
	if err != nil {
		return err
	}

	currentCmd = cfg.Cmd[strings.Index(cfg.Cmd, "-c")+3 : len(cfg.Cmd)]
	currentOutput = ""
	err = ExecPluginCmd(ctx, containerId, cfg, currentCmd, currentOutput, false)
	if err != nil {
		return err
	}

	currentCmd = fmt.Sprintf("for i in /result/*; do curl -F name=%s -F result=@$i http://[::]:%d/save; done", cfg.Name, GetPortNumber())
	currentOutput = ""
	err = ExecPluginCmd(ctx, containerId, cfg, currentCmd, currentOutput, false)
	if err != nil {
		return err
	}

	return nil
}

// ExecPluginCmd executes the command line and checks if successful
func ExecPluginCmd(ctx context.Context, containerId string, cfg *Config, cmd string, output string, outputCheck bool) error {
	log.Printf("[Plugin agent] Executing command line \"%v\" in container\n", cmd)

	idResponse, err := ExecContainerCmd(ctx, containerId, PrepareCmd(cfg, cmd))
	if err != nil {
		log.Printf("[Plugin agent] Error when executing command line \"%v\" in container\n", cmd)
		return err
	}

	execResponse, err := GetExecResponse(ctx, idResponse)
	if err != nil {
		log.Printf("[Plugin agent] Unable to get output of executed command line \"%v\"\n", cmd)
		return err
	}

	if outputCheck == true && execResponse.StdOut != output {
		log.Printf("[Plugin agent] Incorrect output of executed command line \"%v\", got \"%v\", but expected \"%v\"\n", cmd, execResponse.StdOut, "test")
		return errors.New("incorrect output of executed command line")
	}

	log.Printf("[Plugin agent] Command line \"%v\" successfuly executed\n", cmd)

	return nil
}

// PrepareCmd generates the complete command line that specifies the Unix shell
func PrepareCmd(cfg *Config, cmd string) []string {
	bashCmd := strings.Split(cfg.Cmd[0:strings.Index(cfg.Cmd, "-c")+2], " ")
	return append(bashCmd, cmd)
}

// ExecContainerCmd executes the command line in the specified container
func ExecContainerCmd(ctx context.Context, containerId string, command []string) (types.IDResponse, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return types.IDResponse{}, err
	}

	config := types.ExecConfig{
		AttachStderr: true,
		AttachStdout: true,
		Cmd:          command,
	}

	return cli.ContainerExecCreate(ctx, containerId, config)
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

	execResponse.StdOut = strings.TrimSuffix(string(stdout), "\n")
	execResponse.StdErr = strings.TrimSuffix(string(stderr), "\n")
	execResponse.ExitCode = res.ExitCode
	return execResponse, nil
}

// StopContainer stops the specified container
func StopContainer(containerId string) error {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}

	fmt.Printf("[Plugin agent] Stopping container %v... \n", containerId[:10])
	err = cli.ContainerStop(ctx, containerId, nil)
	if err != nil {
		return err
	}
	fmt.Printf("[Plugin agent] Container %v stopped successfuly\n", containerId[:10])

	return nil
}
