// SPDX-FileCopyrightText: Cristian Mogildea
//
// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

func StartContainer(cfg *Config) error {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}

	authStr, err := GetRegistryAuth()
	if err != nil {
		return err
	}

	reader, err := cli.ImagePull(ctx, cfg.DockerImg, types.ImagePullOptions{RegistryAuth: authStr})
	if err != nil {
		return err
	}
	io.Copy(os.Stdout, reader)

	containerCmd := GenerateCmd(cfg)

	resp, err := cli.ContainerCreate(ctx,
		&container.Config{
			Image: cfg.DockerImg,
			Cmd:   containerCmd,
			Tty:   false,
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
		return err
	}

	go StartServer(cfg)

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
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
	} else if os.Getenv("GITHUB_TOKEN") != "" { // TODO: check required
		authStr = os.Getenv("GITHUB_TOKEN")
		return authStr, nil
	}

	log.Println("No authentication credentials provided, please check if environment variables are set")
	return "", errors.New("no authentication credentials provided")
}

// GenerateCmd returns a string with complete command to be executed in the container
func GenerateCmd(cfg *Config) []string {
	bashCmd := strings.Split(cfg.Cmd[0:strings.Index(cfg.Cmd, "-c")+2], " ")
	curlCmd := fmt.Sprintf("&& curl -F name=%s -F result=@/result/%s http://localhost:8082/save",
		cfg.Name,
		cfg.Results[0])
	completeCmd := fmt.Sprintf("%s %s %s",
		"mkdir /result &&",
		cfg.Cmd[strings.Index(cfg.Cmd, "-c")+3:len(cfg.Cmd)],
		curlCmd)

	fmt.Println(append(bashCmd, completeCmd))
	return append(bashCmd, completeCmd)
}
