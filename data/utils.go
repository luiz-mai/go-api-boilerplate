package data

import "database/sql"

type scanner interface {
	Scan(dest ...interface{}) error
}

type executor interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Query(query string, args ...interface{}) (*sql.Rows, error)
}
