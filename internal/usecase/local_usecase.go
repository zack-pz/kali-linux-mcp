package usecase

type LocalNmapUseCase interface {
	Hi(name string) (string, error)
}
