package sslscan

import (
	"github.com/zack-pz/kali-linux-mcp/pkg/executor"
	"github.com/zack-pz/kali-linux-mcp/pkg/logger"
)

type sslScanRepository struct {
	exec executor.IExecutor
}

func NewSSLScanRepository(exec executor.IExecutor) *sslScanRepository {
	return &sslScanRepository{exec: exec}
}

func (s *sslScanRepository) EnumerateCipherHTTPS(host string) (string, error) {
	return s.runSSLScan(host)
}

func (s *sslScanRepository) TargetServicePort(host string, port string) (string, error) {
	return s.runSSLScan(host + ":" + port)
}

func (s *sslScanRepository) ExtractCertificate(host string) (string, error) {
	return s.runSSLScan("--show-certificate " + host)
}

func (s *sslScanRepository) CheckObsoleteSSLProtocol(host string) (string, error) {
	return s.runSSLScan("--ssl2 --ssl3 " + host)
}

func (s *sslScanRepository) VerifyTLSVersion(host string) (string, error) {
	return s.runSSLScan("--tls12 --tls13 " + host)
}

func (s *sslScanRepository) DisplayOnlyAcceptedCiphers(host string) (string, error) {
	return s.runSSLScan("--no-failed " + host)
}

func (s *sslScanRepository) SpecifySNIVirtualHost(domain string, host string) (string, error) {
	return s.runSSLScan("--sni-name=" + domain + " " + host)
}

func (s *sslScanRepository) TestForHeartbleedBug(host string) (string, error) {
	return s.runSSLScan("--heartbleed " + host)
}

func (s *sslScanRepository) ScanSMTPWithSTARTTLS(host string) (string, error) {
	return s.runSSLScan("--starttls-smtp " + host)
}

func (s *sslScanRepository) runSSLScan(args string) (string, error) {
	result, err := s.exec.Run("sslscan " + args)
	if err != nil {
		logger.Error(err)
		return "", err
	}
	return result, nil
}
