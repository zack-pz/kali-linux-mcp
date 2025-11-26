package mcp

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/zack-pz/kali-linux-mcp/internal/services/nmap"
	"github.com/zack-pz/kali-linux-mcp/pkg/logger"
)

type NmapHandler struct {
	repository nmap.NmapRepository
}

type Input struct {
	Ip string `json:"ip"`
}

func NewNmapHandler(repository nmap.NmapRepository) *NmapHandler {
	return &NmapHandler{repository: repository}
}

func (l *NmapHandler) DiscoverLiveHosts(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	var input Input
	args, err := json.Marshal(req.Params.Arguments)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to marshal arguments: %v", err)), nil
	}
	if err := json.Unmarshal(args, &input); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to unmarshal arguments: %v", err)), nil
	}

	response, err := l.repository.DiscoverLiveHosts(input.Ip)
	if err != nil {
		logger.Error(err)
		return mcp.NewToolResultError(err.Error()), nil
	}
	return mcp.NewToolResultText(response), nil
}

func (l *NmapHandler) TreatHostsAsOnline(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	var input Input
	args, err := json.Marshal(req.Params.Arguments)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to marshal arguments: %v", err)), nil
	}
	if err := json.Unmarshal(args, &input); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to unmarshal arguments: %v", err)), nil
	}

	response, err := l.repository.TreatHostsAsOnline(input.Ip)
	if err != nil {
		logger.Error(err)
		return mcp.NewToolResultError(err.Error()), nil
	}
	return mcp.NewToolResultText(response), nil
}

func (l *NmapHandler) StealthSYNScan(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	var input Input
	args, err := json.Marshal(req.Params.Arguments)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to marshal arguments: %v", err)), nil
	}
	if err := json.Unmarshal(args, &input); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to unmarshal arguments: %v", err)), nil
	}

	response, err := l.repository.StealthSYNScan(input.Ip)
	if err != nil {
		logger.Error(err)
		return mcp.NewToolResultError(err.Error()), nil
	}
	return mcp.NewToolResultText(response), nil
}

func (l *NmapHandler) FullConnectScan(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	var input Input
	args, err := json.Marshal(req.Params.Arguments)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to marshal arguments: %v", err)), nil
	}
	if err := json.Unmarshal(args, &input); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to unmarshal arguments: %v", err)), nil
	}

	response, err := l.repository.FullConnectScan(input.Ip)
	if err != nil {
		logger.Error(err)
		return mcp.NewToolResultError(err.Error()), nil
	}
	return mcp.NewToolResultText(response), nil
}

func (l *NmapHandler) UDPServicesScan(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	var input Input
	args, err := json.Marshal(req.Params.Arguments)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to marshal arguments: %v", err)), nil
	}
	if err := json.Unmarshal(args, &input); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to unmarshal arguments: %v", err)), nil
	}

	response, err := l.repository.UDPServicesScan(input.Ip)
	if err != nil {
		logger.Error(err)
		return mcp.NewToolResultError(err.Error()), nil
	}
	return mcp.NewToolResultText(response), nil
}

func (l *NmapHandler) ScanFullPortRange(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	var input Input
	args, err := json.Marshal(req.Params.Arguments)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to marshal arguments: %v", err)), nil
	}
	if err := json.Unmarshal(args, &input); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to unmarshal arguments: %v", err)), nil
	}

	response, err := l.repository.ScanFullPortRange(input.Ip)
	if err != nil {
		logger.Error(err)
		return mcp.NewToolResultError(err.Error()), nil
	}
	return mcp.NewToolResultText(response), nil
}

func (l *NmapHandler) ServiceVersionDetection(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	var input Input
	args, err := json.Marshal(req.Params.Arguments)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to marshal arguments: %v", err)), nil
	}
	if err := json.Unmarshal(args, &input); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to unmarshal arguments: %v", err)), nil
	}

	response, err := l.repository.ServiceVersionDetection(input.Ip)
	if err != nil {
		logger.Error(err)
		return mcp.NewToolResultError(err.Error()), nil
	}
	return mcp.NewToolResultText(response), nil
}

func (l *NmapHandler) OperatingSystemFingerprinting(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	var input Input
	args, err := json.Marshal(req.Params.Arguments)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to marshal arguments: %v", err)), nil
	}
	if err := json.Unmarshal(args, &input); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to unmarshal arguments: %v", err)), nil
	}

	response, err := l.repository.OperatingSystemFingerprinting(input.Ip)
	if err != nil {
		logger.Error(err)
		return mcp.NewToolResultError(err.Error()), nil
	}
	return mcp.NewToolResultText(response), nil
}

func (l *NmapHandler) EnableAggressiveOptions(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	var input Input
	args, err := json.Marshal(req.Params.Arguments)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to marshal arguments: %v", err)), nil
	}
	if err := json.Unmarshal(args, &input); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to unmarshal arguments: %v", err)), nil
	}

	response, err := l.repository.EnableAggressiveOptions(input.Ip)
	if err != nil {
		logger.Error(err)
		return mcp.NewToolResultError(err.Error()), nil
	}
	return mcp.NewToolResultText(response), nil
}

func (l *NmapHandler) ScanDefaultScripts(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	var input Input
	args, err := json.Marshal(req.Params.Arguments)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to marshal arguments: %v", err)), nil
	}
	if err := json.Unmarshal(args, &input); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to unmarshal arguments: %v", err)), nil
	}

	response, err := l.repository.ScanDefaultScripts(input.Ip)
	if err != nil {
		logger.Error(err)
		return mcp.NewToolResultError(err.Error()), nil
	}
	return mcp.NewToolResultText(response), nil
}

func (l *NmapHandler) ScanForVulnerabilities(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	var input Input
	args, err := json.Marshal(req.Params.Arguments)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to marshal arguments: %v", err)), nil
	}
	if err := json.Unmarshal(args, &input); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to unmarshal arguments: %v", err)), nil
	}

	response, err := l.repository.ScanForVulnerabilities(input.Ip)
	if err != nil {
		logger.Error(err)
		return mcp.NewToolResultError(err.Error()), nil
	}
	return mcp.NewToolResultText(response), nil
}

func (l *NmapHandler) HighSpeedTiming(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	var input Input
	args, err := json.Marshal(req.Params.Arguments)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to marshal arguments: %v", err)), nil
	}
	if err := json.Unmarshal(args, &input); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to unmarshal arguments: %v", err)), nil
	}

	response, err := l.repository.HighSpeedTiming(input.Ip)
	if err != nil {
		logger.Error(err)
		return mcp.NewToolResultError(err.Error()), nil
	}
	return mcp.NewToolResultText(response), nil
}

func (l *NmapHandler) FirewallPacketFragmentation(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	var input Input
	args, err := json.Marshal(req.Params.Arguments)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to marshal arguments: %v", err)), nil
	}
	if err := json.Unmarshal(args, &input); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to unmarshal arguments: %v", err)), nil
	}

	response, err := l.repository.FirewallPacketFragmentation(input.Ip)
	if err != nil {
		logger.Error(err)
		return mcp.NewToolResultError(err.Error()), nil
	}
	return mcp.NewToolResultText(response), nil
}

func (l *NmapHandler) PoofSourceIP(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	var input Input
	args, err := json.Marshal(req.Params.Arguments)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to marshal arguments: %v", err)), nil
	}
	if err := json.Unmarshal(args, &input); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to unmarshal arguments: %v", err)), nil
	}

	response, err := l.repository.PoofSourceIP(input.Ip)
	if err != nil {
		logger.Error(err)
		return mcp.NewToolResultError(err.Error()), nil
	}
	return mcp.NewToolResultText(response), nil
}

func (l *NmapHandler) Register(s *server.MCPServer) {
	s.AddTool(mcp.NewTool("DiscoverLiveHosts",
		mcp.WithDescription("Performs a 'Ping Sweep' to discover live hosts on the network without scanning their ports."),
		mcp.WithString("ip", mcp.Required(), mcp.Description("Target IP address or range")),
	), l.DiscoverLiveHosts)

	s.AddTool(mcp.NewTool("TreatHostsAsOnline",
		mcp.WithDescription("Skips the host discovery phase and scans the target assuming it is online, useful for bypassing ICMP blocks."),
		mcp.WithString("ip", mcp.Required(), mcp.Description("Target IP address or range")),
	), l.TreatHostsAsOnline)

	s.AddTool(mcp.NewTool("StealthSYNScan",
		mcp.WithDescription("Executes a TCP SYN scan (half-open) which is faster and stealthier because it does not complete the full TCP handshake."),
		mcp.WithString("ip", mcp.Required(), mcp.Description("Target IP address or range")),
	), l.StealthSYNScan)

	s.AddTool(mcp.NewTool("FullConnectScan",
		mcp.WithDescription("Performs a full TCP Connect scan, completing the handshake; this is the default mode for unprivileged users but is noisier."),
		mcp.WithString("ip", mcp.Required(), mcp.Description("Target IP address or range")),
	), l.FullConnectScan)

	s.AddTool(mcp.NewTool("UDPServicesScan",
		mcp.WithDescription("Scans for open UDP ports to identify services like DNS, DHCP, or SNMP."),
		mcp.WithString("ip", mcp.Required(), mcp.Description("Target IP address or range")),
	), l.UDPServicesScan)

	s.AddTool(mcp.NewTool("ScanFullPortRange",
		mcp.WithDescription("Scans the entire range of 65,535 TCP ports."),
		mcp.WithString("ip", mcp.Required(), mcp.Description("Target IP address or range")),
	), l.ScanFullPortRange)

	s.AddTool(mcp.NewTool("ServiceVersionDetection",
		mcp.WithDescription("Probes open ports to determine the exact service name and version number running on the target."),
		mcp.WithString("ip", mcp.Required(), mcp.Description("Target IP address or range")),
	), l.ServiceVersionDetection)

	s.AddTool(mcp.NewTool("OperatingSystemFingerprinting",
		mcp.WithDescription("Analyzes packet responses to guess the target's Operating System, looking for specific TCP/IP stack fingerprints."),
		mcp.WithString("ip", mcp.Required(), mcp.Description("Target IP address or range")),
	), l.OperatingSystemFingerprinting)

	s.AddTool(mcp.NewTool("EnableAggressiveOptions",
		mcp.WithDescription("Enables 'Aggressive Mode', which combines OS detection, version detection, script scanning, and traceroute into a single command."),
		mcp.WithString("ip", mcp.Required(), mcp.Description("Target IP address or range")),
	), l.EnableAggressiveOptions)

	s.AddTool(mcp.NewTool("ScanDefaultScripts",
		mcp.WithDescription("Runs the default set of NSE (Nmap Scripting Engine) scripts to find common misconfigurations and information."),
		mcp.WithString("ip", mcp.Required(), mcp.Description("Target IP address or range")),
	), l.ScanDefaultScripts)

	s.AddTool(mcp.NewTool("ScanForVulnerabilities",
		mcp.WithDescription("Executes a specific category of scripts designed to detect known security vulnerabilities and CVEs on the target."),
		mcp.WithString("ip", mcp.Required(), mcp.Description("Target IP address or range")),
	), l.ScanForVulnerabilities)

	s.AddTool(mcp.NewTool("HighSpeedTiming",
		mcp.WithDescription("Sets the timing template to 'Aggressive' (level 4), which speeds up the scan by reducing timeouts, recommended for reliable networks."),
		mcp.WithString("ip", mcp.Required(), mcp.Description("Target IP address or range")),
	), l.HighSpeedTiming)

	s.AddTool(mcp.NewTool("FirewallPacketFragmentation",
		mcp.WithDescription("Splits packets into smaller fragments to make it harder for packet filters and firewalls to detect the scan signature."),
		mcp.WithString("ip", mcp.Required(), mcp.Description("Target IP address or range")),
	), l.FirewallPacketFragmentation)

	s.AddTool(mcp.NewTool("PoofSourceIP",
		mcp.WithDescription("Spoofs the source IP address by adding decoy addresses to the scan, hiding your real IP among the fakes."),
		mcp.WithString("ip", mcp.Required(), mcp.Description("Target IP address or range")),
	), l.PoofSourceIP)
}
