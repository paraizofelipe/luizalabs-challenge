package storage

import "github.com/jmoiron/sqlx"

type Postgres struct {
	db *sqlx.DB
}

func NewPostgres(db *sqlx.DB) *Postgres {
	return &Postgres{
		db: db,
	}
}

func (p Postgres) Exec(statement string, args ...interface{}) (err error) {
	var stmt *sqlx.Stmt
	if stmt, err = p.db.Preparex(statement); err != nil {
		return
	}
	if _, err = stmt.Exec(args...); err != nil {
		return
	}
	err = stmt.Close()
	return
}

func (p Postgres) Find(statement string, dest interface{}, args ...interface{}) (err error) {
	var stmt *sqlx.Stmt
	if stmt, err = p.db.Preparex(statement); err != nil {
		return
	}
	if err = stmt.Get(dest, args...); err != nil {
		return
	}
	err = stmt.Close()
	return
}

func (p Postgres) FindAll(statement string, dest interface{}, args ...interface{}) (err error) {
	var stmt *sqlx.Stmt
	if stmt, err = p.db.Preparex(statement); err != nil {
		return
	}
	if err = stmt.Select(dest, args...); err != nil {
		return
	}
	err = stmt.Close()
	return
}
