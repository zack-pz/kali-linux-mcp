package main

import (
	"context"
	"log"

	"security-tools/internal/di"
	"security-tools/pkg/config"

	"github.com/docker/docker/client"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"golang.org/x/crypto/ssh"
)

func main() {
	cfg := config.Load()

	ctx := context.Background()

	exec := startConnection(cfg, ctx)
	// if exec is nil, go to panic. Fix this.
	if exec == nil {
		panic("failed to start connection")
	}

	container := di.NewContainer(exec)

	// Create a server with a single tool.
	server := mcp.NewServer(&mcp.Implementation{Name: "server", Version: "v1.0.0"}, nil)
	container.GetNmapHandler().Register(server)

	if err := server.Run(ctx, &mcp.StdioTransport{}); err != nil {
		log.Fatal(err)
	}
}

func startConnection(cfg *config.Config, ctx context.Context) config.IExecutor {
	switch cfg.Environment {

	case "docker":
		cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
		if err != nil {
			log.Fatal(err)
		}
		return config.NewDockerExecutor(cli, ctx, "")

	case "ssh":
		cfg := &ssh.ClientConfig{
			User: "root",
			Auth: []ssh.AuthMethod{
				ssh.Password("password"),
			},
		}
		clientSSH, err := ssh.Dial("tcp", "localhost:22", cfg)
		if err != nil {
			log.Fatal(err)
		}
		return config.NewSSHExecutor(clientSSH)

	case "local":
		clientLocal, err := config.NewTerminalClient()
		if err != nil {
			log.Fatal(err)
		}
		return clientLocal

	default:
		return nil
	}
}
