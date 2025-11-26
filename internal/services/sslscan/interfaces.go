package sslscan

type SSLScanRepository interface {
	EnumerateCipherHTTPS(host string) (string, error)
	TargetServicePort(host string, port string) (string, error)
	ExtractCertificate(host string) (string, error)
	CheckObsoleteSSLProtocol(host string) (string, error)
	VerifyTLSVersion(host string) (string, error)
	DisplayOnlyAcceptedCiphers(host string) (string, error)
	SpecifySNIVirtualHost(domain string, host string) (string, error)
	TestForHeartbleedBug(host string) (string, error)
	ScanSMTPWithSTARTTLS(host string) (string, error)
}
