package nmap

type NmapRepository interface {
	Hi(name string) (string, error)
	DiscoverLiveHosts(command string) (string, error)
	TreatHostsAsOnline(command string) (string, error)
	StealthSYNScan(command string) (string, error)
	FullConnectScan(command string) (string, error)
	UDPServicesScan(command string) (string, error)
	ScanFullPortRange(command string) (string, error)
	ServiceVersionDetection(command string) (string, error)
	OperatingSystemFingerprinting(command string) (string, error)
	EnableAggressiveOptions(command string) (string, error)
	ScanDefaultScripts(command string) (string, error)
	ScanForVulnerabilities(command string) (string, error)
	HighSpeedTiming(command string) (string, error)
	FirewallPacketFragmentation(command string) (string, error)
	PoofSourceIP(command string) (string, error)
}
