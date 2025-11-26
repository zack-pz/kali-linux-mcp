package sqlmap

import (
	"github.com/zack-pz/kali-linux-mcp/pkg/executor"
	"github.com/zack-pz/kali-linux-mcp/pkg/logger"
)

type sqlmapRepository struct {
	exec executor.IExecutor
}

func NewSqlmapRepository(exec executor.IExecutor) *sqlmapRepository {
	return &sqlmapRepository{exec: exec}
}

func (s *sqlmapRepository) TestSingleURL(url string) (string, error) {
	return s.runSqlmap("-u \"" + url + "\"")
}

func (s *sqlmapRepository) LoadRequestFromFile(file string) (string, error) {
	return s.runSqlmap("-r " + file)
}

func (s *sqlmapRepository) EnumerateDatabases(url string) (string, error) {
	return s.runSqlmap("-u \"" + url + "\" --dbs")
}

func (s *sqlmapRepository) ListTables(url string, db string) (string, error) {
	return s.runSqlmap("-u \"" + url + "\" -D " + db + " --tables")
}

func (s *sqlmapRepository) DumpTable(url string, db string, table string) (string, error) {
	return s.runSqlmap("-u \"" + url + "\" -D " + db + " -T " + table + " --dump")
}

func (s *sqlmapRepository) DumpAll(url string, db string) (string, error) {
	return s.runSqlmap("-u \"" + url + "\" -D " + db + " --dump-all")
}

func (s *sqlmapRepository) RunBatch(url string) (string, error) {
	return s.runSqlmap("-u \"" + url + "\" --batch")
}

func (s *sqlmapRepository) IncreaseRiskLevel(url string, level string, risk string) (string, error) {
	return s.runSqlmap("-u \"" + url + "\" --level=" + level + " --risk=" + risk)
}

func (s *sqlmapRepository) InjectCookie(url string, cookie string) (string, error) {
	return s.runSqlmap("-u \"" + url + "\" --cookie=\"" + cookie + "\"")
}

func (s *sqlmapRepository) SpawnOSShell(url string) (string, error) {
	return s.runSqlmap("-u \"" + url + "\" --os-shell")
}

func (s *sqlmapRepository) CheckTor(url string) (string, error) {
	return s.runSqlmap("-u \"" + url + "\" --tor --check-tor")
}

func (s *sqlmapRepository) FlushSession(url string) (string, error) {
	return s.runSqlmap("-u \"" + url + "\" --flush-session")
}

func (s *sqlmapRepository) runSqlmap(args string) (string, error) {
	result, err := s.exec.Run("sqlmap " + args)
	if err != nil {
		logger.Error(err)
		return "", err
	}
	return result, nil
}
