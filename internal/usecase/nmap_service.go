package usecase

import "github.com/zack-pz/kali-linux-mcp/internal/repository"

type localNmapService struct {
	repo repository.LocalNmapRepository
}

func NewLocalNmapService(repo repository.LocalNmapRepository) LocalNmapUseCase {
	return &localNmapService{repo: repo}
}

func (s *localNmapService) Hi(name string) (string, error) {
	return s.repo.Hi(name)
}
