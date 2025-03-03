package mysql

type SQLite struct{}

func NewSQLite() *SQLite {
	return &SQLite{}
}

func (s *SQLite) Open() error   { return nil }
func (s *SQLite) Close() error  { return nil }
func (s *SQLite) MethodFirst()  {}
func (s *SQLite) MethodSecond() {}
