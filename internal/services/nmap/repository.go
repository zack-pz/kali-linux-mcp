package nmap

import (
	"github.com/zack-pz/kali-linux-mcp/pkg/executor"
	"github.com/zack-pz/kali-linux-mcp/pkg/logger"
)

type nmapRepository struct {
	exec executor.IExecutor
}

func NewNmapRepository(exec executor.IExecutor) *nmapRepository {
	return &nmapRepository{exec: exec}
}

func (n *nmapRepository) Hi(name string) (string, error) {
	name = "'" + name + "'"
	result, err := n.exec.Run("echo 'hi ' + " + name)
	if err != nil {
		logger.Error(err)
		return "", err
	}
	return result, nil
}

func (n *nmapRepository) DiscoverLiveHosts(ip string) (string, error) {
	return n.runNmap("-sn " + ip)
}

func (n *nmapRepository) TreatHostsAsOnline(ip string) (string, error) {
	return n.runNmap("-Pn " + ip)
}

func (n *nmapRepository) StealthSYNScan(ip string) (string, error) {
	return n.runNmap("-sS " + ip)
}

func (n *nmapRepository) FullConnectScan(ip string) (string, error) {
	return n.runNmap("-sT " + ip)
}

func (n *nmapRepository) UDPServicesScan(ip string) (string, error) {
	return n.runNmap("-sU " + ip)
}

func (n *nmapRepository) ScanFullPortRange(ip string) (string, error) {
	return n.runNmap("-p- " + ip)
}

func (n *nmapRepository) ServiceVersionDetection(ip string) (string, error) {
	return n.runNmap("-sV " + ip)
}

func (n *nmapRepository) OperatingSystemFingerprinting(ip string) (string, error) {
	return n.runNmap("-O " + ip)
}

func (n *nmapRepository) EnableAggressiveOptions(ip string) (string, error) {
	return n.runNmap("-A " + ip)
}

func (n *nmapRepository) ScanDefaultScripts(ip string) (string, error) {
	return n.runNmap("-sC " + ip)
}

func (n *nmapRepository) ScanForVulnerabilities(ip string) (string, error) {
	return n.runNmap("--script vuln " + ip)
}

func (n *nmapRepository) HighSpeedTiming(ip string) (string, error) {
	return n.runNmap("-T4 " + ip)
}

func (n *nmapRepository) FirewallPacketFragmentation(ip string) (string, error) {
	return n.runNmap("-f " + ip)
}

func (n *nmapRepository) PoofSourceIP(ip string) (string, error) {
	return n.runNmap("-S " + ip)
}

func (n *nmapRepository) runNmap(args string) (string, error) {
	result, err := n.exec.Run("nmap " + args)
	if err != nil {
		logger.Error(err)
		return "", err
	}
	return result, nil
}
