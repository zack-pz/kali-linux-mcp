package di

import (
	"github.com/zack-pz/kali-linux-mcp/internal/mcp"
	"github.com/zack-pz/kali-linux-mcp/internal/services/nmap"
	"github.com/zack-pz/kali-linux-mcp/pkg/executor"
)

type Container struct {
	// dockerExecutor *config.DockerExecutor
	// sshExecutor    *ssh.Client
	// localExecutor  *config.TerminalClient
	exec executor.IExecutor

	// Repository
	localNmapRepository nmap.NmapRepository

	// Handler
	localNmapHandler *mcp.NmapHandler
}

func NewContainer(exec executor.IExecutor) *Container {
	c := &Container{
		exec: exec,
	}
	c.setupRepositories()
	c.setupHandlers()
	return c
}

func (c *Container) setupRepositories() {
	c.localNmapRepository = nmap.NewNmapRepository(c.exec)
}

func (c *Container) setupHandlers() {
	c.localNmapHandler = mcp.NewNmapHandler(c.localNmapRepository)
}

// Getters
func (c *Container) GetNmapHandler() *mcp.NmapHandler {
	return c.localNmapHandler
}
