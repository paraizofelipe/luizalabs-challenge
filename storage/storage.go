package storage

type PostgresStorage interface {
	Exec(statement string, args ...interface{}) error
	Find(statement string, dest interface{}, args ...interface{}) error
	FindAll(statement string, dest interface{}, args ...interface{}) error
}
