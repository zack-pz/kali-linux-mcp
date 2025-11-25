package repository

type LocalNmapRepository interface {
	Hi(name string) (string, error)
}
