package mcp

import (
	"context"

	"github.com/zack-pz/kali-linux-mcp/internal/services/nmap"
	"github.com/zack-pz/kali-linux-mcp/pkg/logger"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type NmapHandler struct {
	repository nmap.NmapRepository
}

type NmapInput struct {
	Ip string `json:"ip"`
}

type NmapOutput struct {
	Response string `json:"response"`
}

func NewNmapHandler(repository nmap.NmapRepository) *NmapHandler {
	return &NmapHandler{repository: repository}
}

func (l *NmapHandler) Scan(ctx context.Context, req *mcp.CallToolRequest, input NmapInput) (
	*mcp.CallToolResult,
	NmapOutput,
	error,
) {
	response, err := l.repository.DiscoverLiveHosts(input.Ip)
	if err != nil {
		logger.Error(err)
		return nil, NmapOutput{}, err
	}
	return nil, NmapOutput{Response: response}, nil
}

func (l *NmapHandler) DiscoverLiveHosts(ctx context.Context, req *mcp.CallToolRequest, input NmapInput) (
	*mcp.CallToolResult,
	NmapOutput,
	error,
) {
	response, err := l.repository.DiscoverLiveHosts(input.Ip)
	if err != nil {
		logger.Error(err)
		return nil, NmapOutput{}, err
	}
	return nil, NmapOutput{Response: response}, nil
}

func (l *NmapHandler) TreatHostsAsOnline(ctx context.Context, req *mcp.CallToolRequest, input NmapInput) (
	*mcp.CallToolResult,
	NmapOutput,
	error,
) {
	response, err := l.repository.TreatHostsAsOnline(input.Ip)
	if err != nil {
		logger.Error(err)
		return nil, NmapOutput{}, err
	}
	return nil, NmapOutput{Response: response}, nil
}

func (l *NmapHandler) StealthSYNScan(ctx context.Context, req *mcp.CallToolRequest, input NmapInput) (
	*mcp.CallToolResult,
	NmapOutput,
	error,
) {
	response, err := l.repository.StealthSYNScan(input.Ip)
	if err != nil {
		logger.Error(err)
		return nil, NmapOutput{}, err
	}
	return nil, NmapOutput{Response: response}, nil
}

func (l *NmapHandler) FullConnectScan(ctx context.Context, req *mcp.CallToolRequest, input NmapInput) (
	*mcp.CallToolResult,
	NmapOutput,
	error,
) {
	response, err := l.repository.FullConnectScan(input.Ip)
	if err != nil {
		logger.Error(err)
		return nil, NmapOutput{}, err
	}
	return nil, NmapOutput{Response: response}, nil
}

func (l *NmapHandler) UDPServicesScan(ctx context.Context, req *mcp.CallToolRequest, input NmapInput) (
	*mcp.CallToolResult,
	NmapOutput,
	error,
) {
	response, err := l.repository.UDPServicesScan(input.Ip)
	if err != nil {
		logger.Error(err)
		return nil, NmapOutput{}, err
	}
	return nil, NmapOutput{Response: response}, nil
}

func (l *NmapHandler) ScanFullPortRange(ctx context.Context, req *mcp.CallToolRequest, input NmapInput) (
	*mcp.CallToolResult,
	NmapOutput,
	error,
) {
	response, err := l.repository.ScanFullPortRange(input.Ip)
	if err != nil {
		logger.Error(err)
		return nil, NmapOutput{}, err
	}
	return nil, NmapOutput{Response: response}, nil
}

func (l *NmapHandler) ServiceVersionDetection(ctx context.Context, req *mcp.CallToolRequest, input NmapInput) (
	*mcp.CallToolResult,
	NmapOutput,
	error,
) {
	response, err := l.repository.ServiceVersionDetection(input.Ip)
	if err != nil {
		logger.Error(err)
		return nil, NmapOutput{}, err
	}
	return nil, NmapOutput{Response: response}, nil
}

func (l *NmapHandler) OperatingSystemFingerprinting(ctx context.Context, req *mcp.CallToolRequest, input NmapInput) (
	*mcp.CallToolResult,
	NmapOutput,
	error,
) {
	response, err := l.repository.OperatingSystemFingerprinting(input.Ip)
	if err != nil {
		logger.Error(err)
		return nil, NmapOutput{}, err
	}
	return nil, NmapOutput{Response: response}, nil
}

func (l *NmapHandler) EnableAggressiveOptions(ctx context.Context, req *mcp.CallToolRequest, input NmapInput) (
	*mcp.CallToolResult,
	NmapOutput,
	error,
) {
	response, err := l.repository.EnableAggressiveOptions(input.Ip)
	if err != nil {
		logger.Error(err)
		return nil, NmapOutput{}, err
	}
	return nil, NmapOutput{Response: response}, nil
}

func (l *NmapHandler) ScanDefaultScripts(ctx context.Context, req *mcp.CallToolRequest, input NmapInput) (
	*mcp.CallToolResult,
	NmapOutput,
	error,
) {
	response, err := l.repository.ScanDefaultScripts(input.Ip)
	if err != nil {
		logger.Error(err)
		return nil, NmapOutput{}, err
	}
	return nil, NmapOutput{Response: response}, nil
}

func (l *NmapHandler) ScanForVulnerabilities(ctx context.Context, req *mcp.CallToolRequest, input NmapInput) (
	*mcp.CallToolResult,
	NmapOutput,
	error,
) {
	response, err := l.repository.ScanForVulnerabilities(input.Ip)
	if err != nil {
		logger.Error(err)
		return nil, NmapOutput{}, err
	}
	return nil, NmapOutput{Response: response}, nil
}

func (l *NmapHandler) HighSpeedTiming(ctx context.Context, req *mcp.CallToolRequest, input NmapInput) (
	*mcp.CallToolResult,
	NmapOutput,
	error,
) {
	response, err := l.repository.HighSpeedTiming(input.Ip)
	if err != nil {
		logger.Error(err)
		return nil, NmapOutput{}, err
	}
	return nil, NmapOutput{Response: response}, nil
}

func (l *NmapHandler) FirewallPacketFragmentation(ctx context.Context, req *mcp.CallToolRequest, input NmapInput) (
	*mcp.CallToolResult,
	NmapOutput,
	error,
) {
	response, err := l.repository.FirewallPacketFragmentation(input.Ip)
	if err != nil {
		logger.Error(err)
		return nil, NmapOutput{}, err
	}
	return nil, NmapOutput{Response: response}, nil
}

func (l *NmapHandler) PoofSourceIP(ctx context.Context, req *mcp.CallToolRequest, input NmapInput) (
	*mcp.CallToolResult,
	NmapOutput,
	error,
) {
	response, err := l.repository.PoofSourceIP(input.Ip)
	if err != nil {
		logger.Error(err)
		return nil, NmapOutput{}, err
	}
	return nil, NmapOutput{Response: response}, nil
}

func (l *NmapHandler) Register(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{Name: "NmapDiscoverLiveHosts", Description: "Performs a 'Ping Sweep' to discover live hosts on the network without scanning their ports."}, l.DiscoverLiveHosts)
	mcp.AddTool(server, &mcp.Tool{Name: "NmapTreatHostsAsOnline", Description: "Skips the host discovery phase and scans the target assuming it is online, useful for bypassing ICMP blocks."}, l.TreatHostsAsOnline)
	mcp.AddTool(server, &mcp.Tool{Name: "NmapStealthSYNScan", Description: "Executes a TCP SYN scan (half-open) which is faster and stealthier because it does not complete the full TCP handshake."}, l.StealthSYNScan)
	mcp.AddTool(server, &mcp.Tool{Name: "NmapFullConnectScan", Description: "Performs a full TCP Connect scan, completing the handshake; this is the default mode for unprivileged users but is noisier."}, l.FullConnectScan)
	mcp.AddTool(server, &mcp.Tool{Name: "NmapUDPServicesScan", Description: "Scans for open UDP ports to identify services like DNS, DHCP, or SNMP."}, l.UDPServicesScan)
	mcp.AddTool(server, &mcp.Tool{Name: "NmapScanFullPortRange", Description: "Scans the entire range of 65,535 TCP ports."}, l.ScanFullPortRange)
	mcp.AddTool(server, &mcp.Tool{Name: "NmapServiceVersionDetection", Description: "Probes open ports to determine the exact service name and version number running on the target."}, l.ServiceVersionDetection)
	mcp.AddTool(server, &mcp.Tool{Name: "NmapOperatingSystemFingerprinting", Description: "Analyzes packet responses to guess the target's Operating System, looking for specific TCP/IP stack fingerprints."}, l.OperatingSystemFingerprinting)
	mcp.AddTool(server, &mcp.Tool{Name: "NmapEnableAggressiveOptions", Description: "Enables 'Aggressive Mode', which combines OS detection, version detection, script scanning, and traceroute into a single command."}, l.EnableAggressiveOptions)
	mcp.AddTool(server, &mcp.Tool{Name: "NmapScanDefaultScripts", Description: "Runs the default set of NSE (Nmap Scripting Engine) scripts to find common misconfigurations and information."}, l.ScanDefaultScripts)
	mcp.AddTool(server, &mcp.Tool{Name: "NmapScanForVulnerabilities", Description: "Executes a specific category of scripts designed to detect known security vulnerabilities and CVEs on the target."}, l.ScanForVulnerabilities)
	mcp.AddTool(server, &mcp.Tool{Name: "NmapHighSpeedTiming", Description: "Sets the timing template to 'Aggressive' (level 4), which speeds up the scan by reducing timeouts, recommended for reliable networks."}, l.HighSpeedTiming)
	mcp.AddTool(server, &mcp.Tool{Name: "NmapFirewallPacketFragmentation", Description: "Splits packets into smaller fragments to make it harder for packet filters and firewalls to detect the scan signature."}, l.FirewallPacketFragmentation)
	mcp.AddTool(server, &mcp.Tool{Name: "NmapPoofSourceIP", Description: "Spoofs the source IP address by adding decoy addresses to the scan, hiding your real IP among the fakes."}, l.PoofSourceIP)
}
