package executor

import (
	"bytes"
	"fmt"
	"strings"

	"golang.org/x/crypto/ssh"
)

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
