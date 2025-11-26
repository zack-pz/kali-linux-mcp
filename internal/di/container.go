package di

import (
	"github.com/zack-pz/kali-linux-mcp/internal/mcp"
	"github.com/zack-pz/kali-linux-mcp/internal/services/nmap"
	"github.com/zack-pz/kali-linux-mcp/internal/services/sslscan"
	"github.com/zack-pz/kali-linux-mcp/pkg/executor"
)

type Container struct {
	// dockerExecutor *config.DockerExecutor
	// sshExecutor    *ssh.Client
	// localExecutor  *config.TerminalClient
	exec executor.IExecutor

	// Repository
	nmapRepository    nmap.NmapRepository
	sslScanRepository sslscan.SSLScanRepository

	// Handler
	nmapHandler    *mcp.NmapHandler
	sslScanHandler *mcp.SSLScanHandler
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
	c.nmapRepository = nmap.NewNmapRepository(c.exec)
	c.sslScanRepository = sslscan.NewSSLScanRepository(c.exec)
}

func (c *Container) setupHandlers() {
	c.nmapHandler = mcp.NewNmapHandler(c.nmapRepository)
	c.sslScanHandler = mcp.NewSSLScanHandler(c.sslScanRepository)
}

// Getters
func (c *Container) GetNmapHandler() *mcp.NmapHandler {
	return c.nmapHandler
}

func (c *Container) GetSSLScanHandler() *mcp.SSLScanHandler {
	return c.sslScanHandler
}
