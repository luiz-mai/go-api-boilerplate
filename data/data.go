package data

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/luiz-mai/go-api-boilerplate/config"
)

type Conn struct {
	db   *sql.DB
	todo todoRepo
}

func Connect(cfg config.Config) (*Conn, error) {
	db, err := getDB(cfg)
	if err != nil {
		return nil, err
	}

	conn := new(Conn)
	conn.db = db
	conn.todo = newTodoRepo(db)

	return conn, nil
}

func getDB(cfg config.Config) (*sql.DB, error) {

	mysqlConf := mysql.NewConfig()
	mysqlConf.Net = "tcp"
	mysqlConf.Addr = fmt.Sprintf("%s:%d", cfg.DB.Host, cfg.DB.Port)
	mysqlConf.DBName = cfg.DB.Name
	mysqlConf.User = cfg.DB.User
	mysqlConf.Passwd = cfg.DB.Pass
	mysqlConf.ParseTime = true
	mysqlConf.Params = map[string]string{
		"transaction_isolation": "'READ-COMMITTED'",
		//In case you are using a MySQL version prior to 5.7.20, use tx_isolation instead of transaction_isolation
	}

	db, err := sql.Open("mysql", mysqlConf.FormatDSN())
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil

}

func (c Conn) ToDo() todoRepo {
	return c.todo
}
