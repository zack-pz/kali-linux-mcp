package executor

type IExecutor interface {
	Run(command string) (string, error)
}
