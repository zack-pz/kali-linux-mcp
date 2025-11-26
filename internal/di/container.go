package di

import (
	"github.com/zack-pz/kali-linux-mcp/internal/mcp"
	"github.com/zack-pz/kali-linux-mcp/internal/services/nmap"
	"github.com/zack-pz/kali-linux-mcp/internal/services/sqlmap"
	"github.com/zack-pz/kali-linux-mcp/internal/services/sslscan"
	"github.com/zack-pz/kali-linux-mcp/internal/services/sslyze"
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
	sqlmapRepository  sqlmap.SqlmapRepository
	sslyzeRepository  sslyze.SslyzeRepository

	// Handler
	nmapHandler    *mcp.NmapHandler
	sslScanHandler *mcp.SSLScanHandler
	sqlmapHandler  *mcp.SqlmapHandler
	sslyzeHandler  *mcp.SslyzeHandler
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
	c.sqlmapRepository = sqlmap.NewSqlmapRepository(c.exec)
	c.sslyzeRepository = sslyze.NewSslyzeRepository(c.exec)
}

func (c *Container) setupHandlers() {
	c.nmapHandler = mcp.NewNmapHandler(c.nmapRepository)
	c.sslScanHandler = mcp.NewSSLScanHandler(c.sslScanRepository)
	c.sqlmapHandler = mcp.NewSqlmapHandler(c.sqlmapRepository)
	c.sslyzeHandler = mcp.NewSslyzeHandler(c.sslyzeRepository)
}

// Getters
func (c *Container) GetNmapHandler() *mcp.NmapHandler {
	return c.nmapHandler
}

func (c *Container) GetSSLScanHandler() *mcp.SSLScanHandler {
	return c.sslScanHandler
}

func (c *Container) GetSqlmapHandler() *mcp.SqlmapHandler {
	return c.sqlmapHandler
}

func (c *Container) GetSslyzeHandler() *mcp.SslyzeHandler {
	return c.sslyzeHandler
}
