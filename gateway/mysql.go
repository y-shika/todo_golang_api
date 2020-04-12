package gateway

import (
	"database/sql"
	"os"

	"github.com/go-sql-driver/mysql"
)

// MySQLClient is a sql.DB wrapper.
type MySQLClient struct {
	*sql.DB
}

// NewMySQLClient is a new MySQL client.
func NewMySQLClient() (*MySQLClient, error) {
	mysqlConf := mysql.Config{
		User:   "root",
		Passwd: os.Getenv("MYSQL_ROOT_PASSWORD"),
		Net:    "tcp",
		// ${MySQLコンテナ名}:${ポート番号}
		Addr:                 "mysql_db:3306",
		DBName:               "todo_golang_api",
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", mysqlConf.FormatDSN())
	if err != nil {
		return nil, err
	}
	return &MySQLClient{db}, nil
}
