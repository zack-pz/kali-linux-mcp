package executor

import (
	"context"
	"io"
	"strings"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type DockerExecutor struct {
	Client      *client.Client
	Ctx         context.Context
	ContainerID string
}

func NewDockerExecutor(client *client.Client, ctx context.Context, containerID string) *DockerExecutor {
	return &DockerExecutor{
		Client:      client,
		Ctx:         ctx,
		ContainerID: containerID,
	}
}

func (d *DockerExecutor) Run(command string) (string, error) {
	execConfig := container.ExecOptions{
		Cmd:          []string{"/bin/sh", "-c", command},
		AttachStdout: true,
		AttachStderr: true,
		Tty:          true,
	}
	respID, err := d.Client.ContainerExecCreate(d.Ctx, d.ContainerID, execConfig)
	if err != nil {
		return "", err
	}

	resp, err := d.Client.ContainerExecAttach(d.Ctx, respID.ID, container.ExecStartOptions{})
	if err != nil {
		return "", err
	}
	defer resp.Close()

	output, err := io.ReadAll(resp.Reader)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}
