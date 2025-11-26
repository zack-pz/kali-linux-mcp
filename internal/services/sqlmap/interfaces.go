package sqlmap

type SqlmapRepository interface {
	TestSingleURL(url string) (string, error)
	LoadRequestFromFile(file string) (string, error)
	EnumerateDatabases(url string) (string, error)
	ListTables(url string, db string) (string, error)
	DumpTable(url string, db string, table string) (string, error)
	DumpAll(url string, db string) (string, error)
	RunBatch(url string) (string, error)
	IncreaseRiskLevel(url string, level string, risk string) (string, error)
	InjectCookie(url string, cookie string) (string, error)
	SpawnOSShell(url string) (string, error)
	CheckTor(url string) (string, error)
	FlushSession(url string) (string, error)
}
