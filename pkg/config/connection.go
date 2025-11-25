package config

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"golang.org/x/crypto/ssh"
)

type IExecutor interface {
	Run(command string) (string, error)
}

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

type SSHExecutor struct {
	client *ssh.Client
}

func NewSSHExecutor(client *ssh.Client) *SSHExecutor {
	return &SSHExecutor{
		client: client,
	}
}

func (s *SSHExecutor) Run(command string) (string, error) {
	session, err := s.client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	var b bytes.Buffer
	session.Stdout = &b
	session.Stderr = &b

	if err := session.Run(command); err != nil {
		return b.String(), fmt.Errorf("error cmd ssh: %v", err)
	}
	return strings.TrimSpace(b.String()), nil
}

// type LocalExecutor struct{}

// func (l *LocalExecutor) Run(command string) (string, error) {
// 	cmd := exec.Command("/bin/sh", "-c", command)
// 	output, err := cmd.CombinedOutput()
// 	if err != nil {
// 		return "", fmt.Errorf("error cmd local: %v", err)
// 	}
// 	return strings.TrimSpace(string(output)), nil
// }
