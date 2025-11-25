package mcp

import (
	"context"
	"security-tools/internal/usecase"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type LocalHandler struct {
	usecase usecase.LocalNmapUseCase
}

type GreetInput struct {
	Name string `json:"name"`
}

type GreetOutput struct {
	Greeting string `json:"greeting"`
}

func NewLocalHandler(usecase usecase.LocalNmapUseCase) *LocalHandler {
	return &LocalHandler{usecase: usecase}
}

func (l *LocalHandler) SayHi(ctx context.Context, req *mcp.CallToolRequest, input GreetInput) (
	*mcp.CallToolResult,
	GreetOutput,
	error,
) {
	greeting, err := l.usecase.Hi(input.Name)
	if err != nil {
		return nil, GreetOutput{}, err
	}
	return nil, GreetOutput{Greeting: greeting}, nil
}

func (l *LocalHandler) Register(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{Name: "greet", Description: "say hi"}, l.SayHi)
}
