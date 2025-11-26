package mcp

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/zack-pz/kali-linux-mcp/internal/services/sslyze"
	"github.com/zack-pz/kali-linux-mcp/pkg/logger"
)

type SslyzeHandler struct {
	repository sslyze.SslyzeRepository
}

type SslyzeHostInput struct {
	Host string `json:"host"`
}

type SslyzeTargetInput struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type SslyzeJsonInput struct {
	Host string `json:"host"`
	File string `json:"file"`
}

type SslyzeFileInput struct {
	File string `json:"file"`
}

type SslyzeTimeoutInput struct {
	Host    string `json:"host"`
	Timeout string `json:"timeout"`
}

type SslyzeOutput struct {
	Response string `json:"response"`
}

func NewSslyzeHandler(repository sslyze.SslyzeRepository) *SslyzeHandler {
	return &SslyzeHandler{repository: repository}
}

func (h *SslyzeHandler) PerformComprehensiveAnalysis(ctx context.Context, req *mcp.CallToolRequest, input SslyzeHostInput) (
	*mcp.CallToolResult,
	SslyzeOutput,
	error,
) {
	result, err := h.repository.PerformComprehensiveAnalysis(input.Host)
	if err != nil {
		logger.Error(err)
		return nil, SslyzeOutput{}, err
	}
	return &mcp.CallToolResult{}, SslyzeOutput{Response: result}, nil
}

func (h *SslyzeHandler) TargetSpecificPort(ctx context.Context, req *mcp.CallToolRequest, input SslyzeTargetInput) (
	*mcp.CallToolResult,
	SslyzeOutput,
	error,
) {
	result, err := h.repository.TargetSpecificPort(input.Host, input.Port)
	if err != nil {
		logger.Error(err)
		return nil, SslyzeOutput{}, err
	}
	return &mcp.CallToolResult{}, SslyzeOutput{Response: result}, nil
}

func (h *SslyzeHandler) ExportScanResultsToJSON(ctx context.Context, req *mcp.CallToolRequest, input SslyzeJsonInput) (
	*mcp.CallToolResult,
	SslyzeOutput,
	error,
) {
	result, err := h.repository.ExportScanResultsToJSON(input.Host, input.File)
	if err != nil {
		logger.Error(err)
		return nil, SslyzeOutput{}, err
	}
	return &mcp.CallToolResult{}, SslyzeOutput{Response: result}, nil
}

func (h *SslyzeHandler) ScanSMTPWithSTARTTLS(ctx context.Context, req *mcp.CallToolRequest, input SslyzeHostInput) (
	*mcp.CallToolResult,
	SslyzeOutput,
	error,
) {
	result, err := h.repository.ScanSMTPWithSTARTTLS(input.Host)
	if err != nil {
		logger.Error(err)
		return nil, SslyzeOutput{}, err
	}
	return &mcp.CallToolResult{}, SslyzeOutput{Response: result}, nil
}

func (h *SslyzeHandler) CheckHeartbleed(ctx context.Context, req *mcp.CallToolRequest, input SslyzeHostInput) (
	*mcp.CallToolResult,
	SslyzeOutput,
	error,
) {
	result, err := h.repository.CheckHeartbleed(input.Host)
	if err != nil {
		logger.Error(err)
		return nil, SslyzeOutput{}, err
	}
	return &mcp.CallToolResult{}, SslyzeOutput{Response: result}, nil
}

func (h *SslyzeHandler) TestRobotVulnerability(ctx context.Context, req *mcp.CallToolRequest, input SslyzeHostInput) (
	*mcp.CallToolResult,
	SslyzeOutput,
	error,
) {
	result, err := h.repository.TestRobotVulnerability(input.Host)
	if err != nil {
		logger.Error(err)
		return nil, SslyzeOutput{}, err
	}
	return &mcp.CallToolResult{}, SslyzeOutput{Response: result}, nil
}

func (h *SslyzeHandler) CheckOpenSSLCCSInjection(ctx context.Context, req *mcp.CallToolRequest, input SslyzeHostInput) (
	*mcp.CallToolResult,
	SslyzeOutput,
	error,
) {
	result, err := h.repository.CheckOpenSSLCCSInjection(input.Host)
	if err != nil {
		logger.Error(err)
		return nil, SslyzeOutput{}, err
	}
	return &mcp.CallToolResult{}, SslyzeOutput{Response: result}, nil
}

func (h *SslyzeHandler) ScanMultipleTargetsFromFile(ctx context.Context, req *mcp.CallToolRequest, input SslyzeFileInput) (
	*mcp.CallToolResult,
	SslyzeOutput,
	error,
) {
	result, err := h.repository.ScanMultipleTargetsFromFile(input.File)
	if err != nil {
		logger.Error(err)
		return nil, SslyzeOutput{}, err
	}
	return &mcp.CallToolResult{}, SslyzeOutput{Response: result}, nil
}

func (h *SslyzeHandler) IncreaseConnectionTimeout(ctx context.Context, req *mcp.CallToolRequest, input SslyzeTimeoutInput) (
	*mcp.CallToolResult,
	SslyzeOutput,
	error,
) {
	result, err := h.repository.IncreaseConnectionTimeout(input.Host, input.Timeout)
	if err != nil {
		logger.Error(err)
		return nil, SslyzeOutput{}, err
	}
	return &mcp.CallToolResult{}, SslyzeOutput{Response: result}, nil
}

func (h *SslyzeHandler) CheckTLS13EarlyData(ctx context.Context, req *mcp.CallToolRequest, input SslyzeHostInput) (
	*mcp.CallToolResult,
	SslyzeOutput,
	error,
) {
	result, err := h.repository.CheckTLS13EarlyData(input.Host)
	if err != nil {
		logger.Error(err)
		return nil, SslyzeOutput{}, err
	}
	return &mcp.CallToolResult{}, SslyzeOutput{Response: result}, nil
}

func (h *SslyzeHandler) Register(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{Name: "SslyzePerformComprehensiveAnalysis", Description: "Perform comprehensive default SSL/TLS analysis"}, h.PerformComprehensiveAnalysis)
	mcp.AddTool(server, &mcp.Tool{Name: "SslyzeTargetSpecificPort", Description: "Target specific port for SSL analysis"}, h.TargetSpecificPort)
	mcp.AddTool(server, &mcp.Tool{Name: "SslyzeExportScanResultsToJSON", Description: "Export scan results to JSON file"}, h.ExportScanResultsToJSON)
	mcp.AddTool(server, &mcp.Tool{Name: "SslyzeScanSMTPWithSTARTTLS", Description: "Scan SMTP server supporting STARTTLS upgrade"}, h.ScanSMTPWithSTARTTLS)
	mcp.AddTool(server, &mcp.Tool{Name: "SslyzeCheckHeartbleed", Description: "Check specifically for Heartbleed vulnerability"}, h.CheckHeartbleed)
	mcp.AddTool(server, &mcp.Tool{Name: "SslyzeTestRobotVulnerability", Description: "Test for Bleichenbacher ROBOT vulnerability"}, h.TestRobotVulnerability)
	mcp.AddTool(server, &mcp.Tool{Name: "SslyzeCheckOpenSSLCCSInjection", Description: "Check for OpenSSL CCS Injection vulnerability"}, h.CheckOpenSSLCCSInjection)
	mcp.AddTool(server, &mcp.Tool{Name: "SslyzeScanMultipleTargetsFromFile", Description: "Scan multiple targets from text file"}, h.ScanMultipleTargetsFromFile)
	mcp.AddTool(server, &mcp.Tool{Name: "SslyzeIncreaseConnectionTimeout", Description: "Increase connection timeout for slow servers"}, h.IncreaseConnectionTimeout)
	mcp.AddTool(server, &mcp.Tool{Name: "SslyzeCheckTLS13EarlyData", Description: "Check support for TLS 1.3 Early Data"}, h.CheckTLS13EarlyData)
}
