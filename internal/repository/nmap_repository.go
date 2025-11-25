package repository

import (
	"github.com/zack-pz/kali-linux-mcp/pkg/executor"
	"github.com/zack-pz/kali-linux-mcp/pkg/logger"
)

type localNmapRepository struct {
	exec executor.IExecutor
}

func NewLocalNmapRepository(exec executor.IExecutor) *localNmapRepository {
	return &localNmapRepository{exec: exec}
}

func (l *localNmapRepository) Hi(name string) (string, error) {
	name = "'" + name + "'"
	result, err := l.exec.Run("echo 'hi ' + " + name)
	if err != nil {
		logger.Error(err)
		return "", err
	}
	return result, nil
}
