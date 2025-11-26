package sslyze

type SslyzeRepository interface {
	PerformComprehensiveAnalysis(host string) (string, error)
	TargetSpecificPort(host string, port string) (string, error)
	ExportScanResultsToJSON(host string, file string) (string, error)
	ScanSMTPWithSTARTTLS(host string) (string, error)
	CheckHeartbleed(host string) (string, error)
	TestRobotVulnerability(host string) (string, error)
	CheckOpenSSLCCSInjection(host string) (string, error)
	ScanMultipleTargetsFromFile(file string) (string, error)
	IncreaseConnectionTimeout(host string, timeout string) (string, error)
	CheckTLS13EarlyData(host string) (string, error)
}
