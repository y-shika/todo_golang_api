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
		User:   os.Getenv("DB_USER"),
		Passwd: os.Getenv("DB_PASSWORD"),
		Net:    "tcp",
		DBName: "todo_golang_api",
	}

	db, err := sql.Open("mysql", mysqlConf.FormatDSN())
	if err != nil {
		return nil, err
	}
	return &MySQLClient{db}, nil
}
