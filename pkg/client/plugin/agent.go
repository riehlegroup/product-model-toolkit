// SPDX-FileCopyrightText: 2021 Cristian Mogildea
//
// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"sync"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/osrgroup/product-model-toolkit/pkg/client/scanning"
)

type agentCfg struct {
	cfg *Config
}

type agent interface {
	execPlugin(*sync.WaitGroup) error
}

// execResponse struct represents output of an executed command
type execResponse struct {
	StdOut   string
	StdErr   string
	ExitCode int
}

func createAgent(cfg *Config) *agentCfg {
	return &agentCfg{cfg: cfg}
}

// execPlugin executes the plugin and returns nil if successful
func (p agentCfg) execPlugin(wg *sync.WaitGroup) error {
	defer wg.Done()

	resp, ctx, cli, err := prepareContainer(p.cfg)
	if err != nil {
		return err
	}

	err = cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
	if err != nil {
		return err
	}

	if coreConfigValues.RestApi == true {
		err = startServer()
		if err != nil {
			return err
		}
	}

	err = execAllPluginCmd(ctx, resp.ID, p.cfg)
	if err != nil {
		return err
	}

	if coreConfigValues.RestApi == false {
		err = getResultsFromContainer(p.cfg, cli, ctx, resp.ID)
		if err != nil {
			return err
		}
	}

	err = stopContainer(p.cfg.Name, resp.ID)
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

	scanning.SendResults(getResultFiles(p.cfg.Id), p.cfg.Name)

	return nil
}

// prepareContainer pulls image from container registry and prepares container for execution
func prepareContainer(cfg *Config) (container.ContainerCreateCreatedBody, context.Context, *client.Client, error) {
	var resp container.ContainerCreateCreatedBody
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return resp, ctx, cli, err
	}

	authStr, err := getRemoteRepoAuth()
	if err != nil {
		return resp, ctx, cli, err
	}

	reader, err := cli.ImagePull(ctx, cfg.DockerImg, types.ImagePullOptions{RegistryAuth: authStr})
	if err == nil {
		io.Copy(os.Stdout, reader)
	}
	if err != nil {
		log.Printf("[Plugin agent] [%v] Unable to pull image from container registry, got following error: %v\n", cfg.Name, err)
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
			Cmd:   []string{cfg.Shell},
			Tty:   true,
		},
		&container.HostConfig{
			Mounts: []mount.Mount{
				{
					Type:     mount.TypeBind,
					Source:   cfg.InDir,
					Target:   "/input",
					ReadOnly: true,
				},
			},
			NetworkMode: "host",
		}, nil, nil, "")
}

// execAllPluginCmd executes all necessary commands in the container
func execAllPluginCmd(ctx context.Context, containerID string, cfg *Config) error {
	logFile, err := createLogFile(cfg.Name)
	if err != nil {
		return err
	}

	coreCfg, err := loadCoreEngineConfig()
	if err != nil {
		return err
	}

	err = compatibilityCheck(ctx, containerID, cfg, logFile, coreCfg)
	if err != nil {
		return err
	}

	currentCmd := "mkdir /result"
	expectedOutput := ""
	err = execPluginCmd(ctx, containerID, cfg, currentCmd, expectedOutput, false, logFile)
	if err != nil {
		return err
	}

	currentCmd = cfg.Cmd
	expectedOutput = ""
	err = execPluginCmd(ctx, containerID, cfg, currentCmd, expectedOutput, false, logFile)
	if err != nil {
		return err
	}

	if coreConfigValues.RestApi == true {
		currentCmd = fmt.Sprintf("for i in /result/*; do curl -F name=%s -F id=%d -F result=@$i http://127.0.0.1:%d/save; done", cfg.Name, cfg.Id, getPortNumber())
		expectedOutput = ""
		err = execPluginCmd(ctx, containerID, cfg, currentCmd, expectedOutput, false, logFile)
		if err != nil {
			return err
		}
	}

	log.Printf("[Plugin agent] [%v] All commands were executed, check log file %v for outputs of executed commands\n", cfg.Name, logFile)

	return nil
}

func compatibilityCheck(ctx context.Context, containerID string, cfg *Config, logFile string, coreCfg *coreConfig) error {
	currentCmd := "echo test"
	expectedOutput := "^test"
	err := execPluginCmd(ctx, containerID, cfg, currentCmd, expectedOutput, true, logFile)
	if err != nil {
		log.Printf("[Plugin agent] [%v] Plugin is not compatible with core engine, failed command: %v\n", cfg.Name, currentCmd)
		return err
	}

	if coreCfg.RestApi == true {
		currentCmd := "curl -V"
		expectedOutput := "^curl"
		err := execPluginCmd(ctx, containerID, cfg, currentCmd, expectedOutput, true, logFile)
		if err != nil {
			log.Printf("[Plugin agent] [%v] Plugin is not compatible with core engine, failed command: %v\n", cfg.Name, currentCmd)
			return err
		}
	}

	return nil
}

// execPluginCmd executes command and checks if successful
func execPluginCmd(ctx context.Context, containerID string, cfg *Config, cmd string, expectedOutput string, outputCheck bool, logFile string) error {
	log.Printf("[Plugin agent] [%v] Executing following command in container: %v\n", cfg.Name, cmd)

	idResponse, err := execContainerCmd(ctx, containerID, prepareCmd(cfg, cmd))
	if err != nil {
		log.Printf("[Plugin agent] [%v] Error when executing following command in container: %v\n", cfg.Name, cmd)
		return err
	}

	execResponse, err := getExecResponse(ctx, idResponse)
	if err != nil {
		log.Printf("[Plugin agent] [%v] Unable to get output of following executed command: %v\n", cfg.Name, cmd)
		return err
	}
	if execResponse.StdOut != "" {
		err = writeToLogFile(logFile, cmd, "stdout", execResponse.StdOut)
		if err != nil {
			log.Printf("[Plugin agent] [%v] Unable to write to log file stdout of following executed command: %v\n", cfg.Name, cmd)
			return err
		}
	}
	if execResponse.StdErr != "" {
		err = writeToLogFile(logFile, cmd, "stderr", execResponse.StdErr)
		if err != nil {
			log.Printf("[Plugin agent] [%v] Unable to write to log file stderr of following executed command: %v\n", cfg.Name, cmd)
			return err
		}
	}

	if outputCheck == true {
		match, err := regexp.MatchString(expectedOutput, execResponse.StdOut)
		if err != nil {
			return err
		}
		if match != true {
			log.Printf("[Plugin agent] [%v] Incorrect output of executed command: %v; got: %v; expected %v\n", cfg.Name, cmd, execResponse.StdOut, expectedOutput)
			return errors.New("incorrect output of executed command")
		}
	}

	log.Printf("[Plugin agent] [%v] Following command successfully executed: %v\n", cfg.Name, cmd)

	return nil
}

// prepareCmd generates complete command
func prepareCmd(cfg *Config, cmd string) []string {
	return append([]string{cfg.Shell, "-c"}, cmd)
}

// execContainerCmd executes command in specified container
func execContainerCmd(ctx context.Context, containerID string, command []string) (types.IDResponse, error) {
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

// getExecResponse returns output of executed command
func getExecResponse(ctx context.Context, idResponse types.IDResponse) (execResponse, error) {
	var execResponse execResponse

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

// stopContainer stops specified container
func stopContainer(pluginName string, containerID string) error {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}

	log.Printf("[Plugin agent] [%v] Stopping container: %v...\n", pluginName, containerID[:10])
	err = cli.ContainerStop(ctx, containerID, nil)
	if err != nil {
		return err
	}

	log.Printf("[Plugin agent] [%v] Container stopped successfully: %v...\n", pluginName, containerID[:10])
	return nil
}
