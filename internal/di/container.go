package di

import (
	"github.com/zack-pz/kali-linux-mcp/internal/handler/mcp"
	"github.com/zack-pz/kali-linux-mcp/internal/repository"
	"github.com/zack-pz/kali-linux-mcp/internal/usecase"
	"github.com/zack-pz/kali-linux-mcp/pkg/executor"
)

type Container struct {
	// dockerExecutor *config.DockerExecutor
	// sshExecutor    *ssh.Client
	// localExecutor  *config.TerminalClient
	exec executor.IExecutor

	// Repository
	localNmapRepository repository.LocalNmapRepository

	// UseCase
	localNmapUseCase usecase.LocalNmapUseCase

	// Handler
	localNmapHandler *mcp.LocalHandler
}

func NewContainer(exec executor.IExecutor) *Container {
	c := &Container{
		exec: exec,
	}
	c.setupRepositories()
	c.setupUseCases()
	c.setupHandlers()
	return c
}

func (c *Container) setupRepositories() {
	c.localNmapRepository = repository.NewLocalNmapRepository(c.exec)
}

func (c *Container) setupUseCases() {
	c.localNmapUseCase = usecase.NewLocalNmapService(c.localNmapRepository)
}

func (c *Container) setupHandlers() {
	c.localNmapHandler = mcp.NewLocalHandler(c.localNmapUseCase)
}

// Getters
func (c *Container) GetNmapHandler() *mcp.LocalHandler {
	return c.localNmapHandler
}
