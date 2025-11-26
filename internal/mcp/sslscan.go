package mcp

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/zack-pz/kali-linux-mcp/internal/services/sslscan"
	"github.com/zack-pz/kali-linux-mcp/pkg/logger"
)

type SSLScanHandler struct {
	repository sslscan.SSLScanRepository
}

type SSLInput struct {
	Host string `json:"host"`
}

type SSLPortInput struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type SSLSNIInput struct {
	Domain string `json:"domain"`
	Host   string `json:"host"`
}

type SSLOutput struct {
	Response string `json:"response"`
}

func NewSSLScanHandler(repository sslscan.SSLScanRepository) *SSLScanHandler {
	return &SSLScanHandler{repository: repository}
}

func (l *SSLScanHandler) EnumerateCipherHTTPS(ctx context.Context, req *mcp.CallToolRequest, input SSLInput) (
	*mcp.CallToolResult,
	SSLOutput,
	error,
) {
	result, err := l.repository.EnumerateCipherHTTPS(input.Host)
	if err != nil {
		logger.Error(err)
		return nil, SSLOutput{}, err
	}
	return &mcp.CallToolResult{}, SSLOutput{Response: result}, nil
}

func (l *SSLScanHandler) TargetServicePort(ctx context.Context, req *mcp.CallToolRequest, input SSLPortInput) (
	*mcp.CallToolResult,
	SSLOutput,
	error,
) {
	result, err := l.repository.TargetServicePort(input.Host, input.Port)
	if err != nil {
		logger.Error(err)
		return nil, SSLOutput{}, err
	}
	return &mcp.CallToolResult{}, SSLOutput{Response: result}, nil
}

func (l *SSLScanHandler) ExtractCertificate(ctx context.Context, req *mcp.CallToolRequest, input SSLInput) (
	*mcp.CallToolResult,
	SSLOutput,
	error,
) {
	result, err := l.repository.ExtractCertificate(input.Host)
	if err != nil {
		logger.Error(err)
		return nil, SSLOutput{}, err
	}
	return &mcp.CallToolResult{}, SSLOutput{Response: result}, nil
}

func (l *SSLScanHandler) CheckObsoleteSSLProtocol(ctx context.Context, req *mcp.CallToolRequest, input SSLInput) (
	*mcp.CallToolResult,
	SSLOutput,
	error,
) {
	result, err := l.repository.CheckObsoleteSSLProtocol(input.Host)
	if err != nil {
		logger.Error(err)
		return nil, SSLOutput{}, err
	}
	return &mcp.CallToolResult{}, SSLOutput{Response: result}, nil
}

func (l *SSLScanHandler) VerifyTLSVersion(ctx context.Context, req *mcp.CallToolRequest, input SSLInput) (
	*mcp.CallToolResult,
	SSLOutput,
	error,
) {
	result, err := l.repository.VerifyTLSVersion(input.Host)
	if err != nil {
		logger.Error(err)
		return nil, SSLOutput{}, err
	}
	return &mcp.CallToolResult{}, SSLOutput{Response: result}, nil
}

func (l *SSLScanHandler) DisplayOnlyAcceptedCiphers(ctx context.Context, req *mcp.CallToolRequest, input SSLInput) (
	*mcp.CallToolResult,
	SSLOutput,
	error,
) {
	result, err := l.repository.DisplayOnlyAcceptedCiphers(input.Host)
	if err != nil {
		logger.Error(err)
		return nil, SSLOutput{}, err
	}
	return &mcp.CallToolResult{}, SSLOutput{Response: result}, nil
}

func (l *SSLScanHandler) SpecifySNIVirtualHost(ctx context.Context, req *mcp.CallToolRequest, input SSLSNIInput) (
	*mcp.CallToolResult,
	SSLOutput,
	error,
) {
	result, err := l.repository.SpecifySNIVirtualHost(input.Domain, input.Host)
	if err != nil {
		logger.Error(err)
		return nil, SSLOutput{}, err
	}
	return &mcp.CallToolResult{}, SSLOutput{Response: result}, nil
}

func (l *SSLScanHandler) TestForHeartbleedBug(ctx context.Context, req *mcp.CallToolRequest, input SSLInput) (
	*mcp.CallToolResult,
	SSLOutput,
	error,
) {
	result, err := l.repository.TestForHeartbleedBug(input.Host)
	if err != nil {
		logger.Error(err)
		return nil, SSLOutput{}, err
	}
	return &mcp.CallToolResult{}, SSLOutput{Response: result}, nil
}

func (l *SSLScanHandler) ScanSMTPWithSTARTTLS(ctx context.Context, req *mcp.CallToolRequest, input SSLInput) (
	*mcp.CallToolResult,
	SSLOutput,
	error,
) {
	result, err := l.repository.ScanSMTPWithSTARTTLS(input.Host)
	if err != nil {
		logger.Error(err)
		return nil, SSLOutput{}, err
	}
	return &mcp.CallToolResult{}, SSLOutput{Response: result}, nil
}

func (l *SSLScanHandler) Register(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{Name: "SSLScanEnumerateCipherHTTPS", Description: "Enumerates supported ciphers for HTTPS on the target host."}, l.EnumerateCipherHTTPS)
	mcp.AddTool(server, &mcp.Tool{Name: "SSLScanTargetServicePort", Description: "Scans a specific port on the target host for SSL/TLS services."}, l.TargetServicePort)
	mcp.AddTool(server, &mcp.Tool{Name: "SSLScanExtractCertificate", Description: "Extracts the SSL certificate from the target host."}, l.ExtractCertificate)
	mcp.AddTool(server, &mcp.Tool{Name: "SSLScanCheckObsoleteSSLProtocol", Description: "Checks if the target host supports obsolete SSL protocols."}, l.CheckObsoleteSSLProtocol)
	mcp.AddTool(server, &mcp.Tool{Name: "SSLScanVerifyTLSVersion", Description: "Verifies the TLS versions supported by the target host."}, l.VerifyTLSVersion)
	mcp.AddTool(server, &mcp.Tool{Name: "SSLScanDisplayOnlyAcceptedCiphers", Description: "Displays only the ciphers accepted by the target host."}, l.DisplayOnlyAcceptedCiphers)
	mcp.AddTool(server, &mcp.Tool{Name: "SSLScanSpecifySNIVirtualHost", Description: "Specifies an SNI virtual host for the scan."}, l.SpecifySNIVirtualHost)
	mcp.AddTool(server, &mcp.Tool{Name: "SSLScanTestForHeartbleedBug", Description: "Tests the target host for the Heartbleed vulnerability."}, l.TestForHeartbleedBug)
	mcp.AddTool(server, &mcp.Tool{Name: "SSLScanScanSMTPWithSTARTTLS", Description: "Scans an SMTP server using STARTTLS."}, l.ScanSMTPWithSTARTTLS)
}
