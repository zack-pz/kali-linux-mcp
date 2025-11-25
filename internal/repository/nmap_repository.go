package repository

import (
	"github.com/zack-pz/kali-linux-mcp/pkg/config"
)

type localNmapRepository struct {
	exec config.IExecutor
}

func NewLocalNmapRepository(exec config.IExecutor) *localNmapRepository {
	return &localNmapRepository{exec: exec}
}

func (l *localNmapRepository) Hi(name string) (string, error) {
	name = "'" + name + "'"
	result, err := l.exec.Run("echo 'hi ' + " + name)
	if err != nil {
		return "", err
	}
	return result, nil
}
