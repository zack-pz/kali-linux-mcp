package sslyze

import (
	"github.com/zack-pz/kali-linux-mcp/pkg/executor"
	"github.com/zack-pz/kali-linux-mcp/pkg/logger"
)

type sslyzeRepository struct {
	exec executor.IExecutor
}

func NewSslyzeRepository(exec executor.IExecutor) *sslyzeRepository {
	return &sslyzeRepository{exec: exec}
}

func (s *sslyzeRepository) PerformComprehensiveAnalysis(host string) (string, error) {
	return s.runSslyze(host)
}

func (s *sslyzeRepository) TargetSpecificPort(host string, port string) (string, error) {
	return s.runSslyze(host + ":" + port)
}

func (s *sslyzeRepository) ExportScanResultsToJSON(host string, file string) (string, error) {
	return s.runSslyze("--json_out=" + file + " " + host)
}

func (s *sslyzeRepository) ScanSMTPWithSTARTTLS(host string) (string, error) {
	return s.runSslyze("--starttls=smtp " + host)
}

func (s *sslyzeRepository) CheckHeartbleed(host string) (string, error) {
	return s.runSslyze("--heartbleed " + host)
}

func (s *sslyzeRepository) TestRobotVulnerability(host string) (string, error) {
	return s.runSslyze("--robot " + host)
}

func (s *sslyzeRepository) CheckOpenSSLCCSInjection(host string) (string, error) {
	return s.runSslyze("--openssl_ccs " + host)
}

func (s *sslyzeRepository) ScanMultipleTargetsFromFile(file string) (string, error) {
	return s.runSslyze("--targets_in=" + file)
}

func (s *sslyzeRepository) IncreaseConnectionTimeout(host string, timeout string) (string, error) {
	return s.runSslyze("--timeout=" + timeout + " " + host)
}

func (s *sslyzeRepository) CheckTLS13EarlyData(host string) (string, error) {
	return s.runSslyze("--early_data " + host)
}

func (s *sslyzeRepository) runSslyze(args string) (string, error) {
	result, err := s.exec.Run("sslyze " + args)
	if err != nil {
		logger.Error(err)
		return "", err
	}
	return result, nil
}
