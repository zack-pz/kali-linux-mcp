package main

import (
	"context"
	"net/url"
	"os"

	"github.com/mark3labs/mcp-go/server"
	"github.com/zack-pz/kali-linux-mcp/internal/di"
	"github.com/zack-pz/kali-linux-mcp/pkg/config"
	"github.com/zack-pz/kali-linux-mcp/pkg/executor"
	"github.com/zack-pz/kali-linux-mcp/pkg/logger"

	"github.com/docker/docker/client"
	"golang.org/x/crypto/ssh"
)

func main() {
	cfg := config.Load()
	logger.Init()

	s := server.NewMCPServer("KaliLinux", "1.0.0")

	ctx := context.Background()
	exec := startConnection(cfg.Kali_url, ctx)
	// if exec is nil, go to panic. Fix this.
	if exec == nil {
		panic("failed to start connection")
	}

	container := di.NewContainer(exec)

	container.GetNmapHandler().Register(s)

	logger.Info("Server started")
	if err := server.NewStdioServer(s).Listen(ctx, os.Stdin, os.Stdout); err != nil {
		logger.Error(err)
	}
}

func startConnection(target string, ctx context.Context) executor.IExecutor {
	u, err := url.Parse(target)
	if err != nil {
		logger.Error(err)
	}

	switch u.Scheme {

	case "docker":
		cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
		if err != nil {
			logger.Error(err)
			return nil
		}
		return executor.NewDockerExecutor(cli, ctx, u.Host)

	case "ssh":
		cfg := &ssh.ClientConfig{
			User: "root",
			Auth: []ssh.AuthMethod{
				ssh.Password("password"),
			},
		}
		clientSSH, err := ssh.Dial("tcp", "localhost:22", cfg)
		if err != nil {
			logger.Error(err)
		}
		return executor.NewSSHExecutor(clientSSH)

	case "local":
		clientLocal, err := executor.NewTerminalClient()
		if err != nil {
			logger.Error(err)
		}
		return clientLocal

	default:
		logger.Error("unknown target scheme")
		return nil
	}
}
