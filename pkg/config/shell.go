package config

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"strings"
	"sync"
	"time"
)

type TerminalClient struct {
	cmd     *exec.Cmd
	stdin   io.WriteCloser
	stdout  io.ReadCloser
	mu      sync.Mutex
	marker  string
	scanner *bufio.Scanner
}

func NewTerminalClient() (IExecutor, error) {
	cmd := exec.Command("bash")

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	client := &TerminalClient{
		cmd:     cmd,
		stdin:   stdin,
		stdout:  stdout,
		marker:  "__END_OF_COMMAND__" + fmt.Sprint(time.Now().UnixNano()),
		scanner: bufio.NewScanner(stdout),
	}

	return client, nil
}

func (t *TerminalClient) Run(cmd string) (string, error) {
	t.mu.Lock()
	defer t.mu.Unlock()

	fullCmd := fmt.Sprintf("%s; echo %s", cmd, t.marker)

	_, err := fmt.Fprintln(t.stdin, fullCmd)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer

	for t.scanner.Scan() {
		line := t.scanner.Text()

		if strings.Contains(line, t.marker) {
			break
		}

		buf.WriteString(line + "\n")
	}

	if err := t.scanner.Err(); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (t *TerminalClient) Close() {
	t.stdin.Close()
	t.cmd.Process.Kill()
}
