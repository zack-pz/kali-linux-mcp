package mcp

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/zack-pz/kali-linux-mcp/internal/services/sqlmap"
	"github.com/zack-pz/kali-linux-mcp/pkg/logger"
)

type SqlmapHandler struct {
	repository sqlmap.SqlmapRepository
}

type SqlmapURLInput struct {
	URL string `json:"url"`
}

type SqlmapFileInput struct {
	File string `json:"file"`
}

type SqlmapDBInput struct {
	URL string `json:"url"`
	DB  string `json:"db"`
}

type SqlmapTableInput struct {
	URL   string `json:"url"`
	DB    string `json:"db"`
	Table string `json:"table"`
}

type SqlmapRiskInput struct {
	URL   string `json:"url"`
	Level string `json:"level"`
	Risk  string `json:"risk"`
}

type SqlmapCookieInput struct {
	URL    string `json:"url"`
	Cookie string `json:"cookie"`
}

type SqlmapOutput struct {
	Response string `json:"response"`
}

func NewSqlmapHandler(repository sqlmap.SqlmapRepository) *SqlmapHandler {
	return &SqlmapHandler{repository: repository}
}

func (h *SqlmapHandler) TestSingleURL(ctx context.Context, req *mcp.CallToolRequest, input SqlmapURLInput) (
	*mcp.CallToolResult,
	SqlmapOutput,
	error,
) {
	result, err := h.repository.TestSingleURL(input.URL)
	if err != nil {
		logger.Error(err)
		return nil, SqlmapOutput{}, err
	}
	return &mcp.CallToolResult{}, SqlmapOutput{Response: result}, nil
}

func (h *SqlmapHandler) LoadRequestFromFile(ctx context.Context, req *mcp.CallToolRequest, input SqlmapFileInput) (
	*mcp.CallToolResult,
	SqlmapOutput,
	error,
) {
	result, err := h.repository.LoadRequestFromFile(input.File)
	if err != nil {
		logger.Error(err)
		return nil, SqlmapOutput{}, err
	}
	return &mcp.CallToolResult{}, SqlmapOutput{Response: result}, nil
}

func (h *SqlmapHandler) EnumerateDatabases(ctx context.Context, req *mcp.CallToolRequest, input SqlmapURLInput) (
	*mcp.CallToolResult,
	SqlmapOutput,
	error,
) {
	result, err := h.repository.EnumerateDatabases(input.URL)
	if err != nil {
		logger.Error(err)
		return nil, SqlmapOutput{}, err
	}
	return &mcp.CallToolResult{}, SqlmapOutput{Response: result}, nil
}

func (h *SqlmapHandler) ListTables(ctx context.Context, req *mcp.CallToolRequest, input SqlmapDBInput) (
	*mcp.CallToolResult,
	SqlmapOutput,
	error,
) {
	result, err := h.repository.ListTables(input.URL, input.DB)
	if err != nil {
		logger.Error(err)
		return nil, SqlmapOutput{}, err
	}
	return &mcp.CallToolResult{}, SqlmapOutput{Response: result}, nil
}

func (h *SqlmapHandler) DumpTable(ctx context.Context, req *mcp.CallToolRequest, input SqlmapTableInput) (
	*mcp.CallToolResult,
	SqlmapOutput,
	error,
) {
	result, err := h.repository.DumpTable(input.URL, input.DB, input.Table)
	if err != nil {
		logger.Error(err)
		return nil, SqlmapOutput{}, err
	}
	return &mcp.CallToolResult{}, SqlmapOutput{Response: result}, nil
}

func (h *SqlmapHandler) DumpAll(ctx context.Context, req *mcp.CallToolRequest, input SqlmapDBInput) (
	*mcp.CallToolResult,
	SqlmapOutput,
	error,
) {
	result, err := h.repository.DumpAll(input.URL, input.DB)
	if err != nil {
		logger.Error(err)
		return nil, SqlmapOutput{}, err
	}
	return &mcp.CallToolResult{}, SqlmapOutput{Response: result}, nil
}

func (h *SqlmapHandler) RunBatch(ctx context.Context, req *mcp.CallToolRequest, input SqlmapURLInput) (
	*mcp.CallToolResult,
	SqlmapOutput,
	error,
) {
	result, err := h.repository.RunBatch(input.URL)
	if err != nil {
		logger.Error(err)
		return nil, SqlmapOutput{}, err
	}
	return &mcp.CallToolResult{}, SqlmapOutput{Response: result}, nil
}

func (h *SqlmapHandler) IncreaseRiskLevel(ctx context.Context, req *mcp.CallToolRequest, input SqlmapRiskInput) (
	*mcp.CallToolResult,
	SqlmapOutput,
	error,
) {
	result, err := h.repository.IncreaseRiskLevel(input.URL, input.Level, input.Risk)
	if err != nil {
		logger.Error(err)
		return nil, SqlmapOutput{}, err
	}
	return &mcp.CallToolResult{}, SqlmapOutput{Response: result}, nil
}

func (h *SqlmapHandler) InjectCookie(ctx context.Context, req *mcp.CallToolRequest, input SqlmapCookieInput) (
	*mcp.CallToolResult,
	SqlmapOutput,
	error,
) {
	result, err := h.repository.InjectCookie(input.URL, input.Cookie)
	if err != nil {
		logger.Error(err)
		return nil, SqlmapOutput{}, err
	}
	return &mcp.CallToolResult{}, SqlmapOutput{Response: result}, nil
}

func (h *SqlmapHandler) SpawnOSShell(ctx context.Context, req *mcp.CallToolRequest, input SqlmapURLInput) (
	*mcp.CallToolResult,
	SqlmapOutput,
	error,
) {
	result, err := h.repository.SpawnOSShell(input.URL)
	if err != nil {
		logger.Error(err)
		return nil, SqlmapOutput{}, err
	}
	return &mcp.CallToolResult{}, SqlmapOutput{Response: result}, nil
}

func (h *SqlmapHandler) CheckTor(ctx context.Context, req *mcp.CallToolRequest, input SqlmapURLInput) (
	*mcp.CallToolResult,
	SqlmapOutput,
	error,
) {
	result, err := h.repository.CheckTor(input.URL)
	if err != nil {
		logger.Error(err)
		return nil, SqlmapOutput{}, err
	}
	return &mcp.CallToolResult{}, SqlmapOutput{Response: result}, nil
}

func (h *SqlmapHandler) FlushSession(ctx context.Context, req *mcp.CallToolRequest, input SqlmapURLInput) (
	*mcp.CallToolResult,
	SqlmapOutput,
	error,
) {
	result, err := h.repository.FlushSession(input.URL)
	if err != nil {
		logger.Error(err)
		return nil, SqlmapOutput{}, err
	}
	return &mcp.CallToolResult{}, SqlmapOutput{Response: result}, nil
}

func (h *SqlmapHandler) Register(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{Name: "SqlmapTestSingleURL", Description: "Test single URL for injection points"}, h.TestSingleURL)
	mcp.AddTool(server, &mcp.Tool{Name: "SqlmapLoadRequestFromFile", Description: "Load HTTP request from saved file"}, h.LoadRequestFromFile)
	mcp.AddTool(server, &mcp.Tool{Name: "SqlmapEnumerateDatabases", Description: "Enumerate all available databases on server"}, h.EnumerateDatabases)
	mcp.AddTool(server, &mcp.Tool{Name: "SqlmapListTables", Description: "List tables within a specific database"}, h.ListTables)
	mcp.AddTool(server, &mcp.Tool{Name: "SqlmapDumpTable", Description: "Dump all data from specific table"}, h.DumpTable)
	mcp.AddTool(server, &mcp.Tool{Name: "SqlmapDumpAll", Description: "Dump entire database contents blindly"}, h.DumpAll)
	mcp.AddTool(server, &mcp.Tool{Name: "SqlmapRunBatch", Description: "Run non-interactively accepting default answers"}, h.RunBatch)
	mcp.AddTool(server, &mcp.Tool{Name: "SqlmapIncreaseRiskLevel", Description: "Increase scan depth and payload risk"}, h.IncreaseRiskLevel)
	mcp.AddTool(server, &mcp.Tool{Name: "SqlmapInjectCookie", Description: "Inject using authenticated session cookie"}, h.InjectCookie)
	mcp.AddTool(server, &mcp.Tool{Name: "SqlmapSpawnOSShell", Description: "Attempt to spawn interactive OS shell"}, h.SpawnOSShell)
	mcp.AddTool(server, &mcp.Tool{Name: "SqlmapCheckTor", Description: "Route traffic through Tor for anonymity"}, h.CheckTor)
	mcp.AddTool(server, &mcp.Tool{Name: "SqlmapFlushSession", Description: "Clear cached data for fresh scan"}, h.FlushSession)
}
