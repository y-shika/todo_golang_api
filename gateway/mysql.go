package gateway

import (
	"database/sql"
	"os"

	"github.com/go-sql-driver/mysql"

	"github.com/y-shika/todo_golang_api/config"
)

// MySQLClient is a sql.DB wrapper.
type MySQLClient struct {
	*sql.DB
}

// NewMySQLClient is a new MySQL client.
func NewMySQLClient() (*MySQLClient, error) {
	// TODO: heroku用とローカル用のクライアントを作成する
	mysqlConf := mysql.Config{}

	// TODO: DBの接続先が増えたらこの部分の条件分岐も増やす
	if config.IsHeroku() || config.IsConnectedHerokuDB() {
		mysqlConf = herokuMySQLConf()
	} else {
		mysqlConf = localMySQLConf()
	}

	db, err := sql.Open("mysql", mysqlConf.FormatDSN())
	if err != nil {
		return nil, err
	}

	return &MySQLClient{db}, nil
}

func herokuMySQLConf() mysql.Config {
	conf := mysql.Config{
		User:                 os.Getenv("HEROKU_DB_USERNAME"),
		Passwd:               os.Getenv("HEROKU_DB_PASSWORD"),
		Net:                  "tcp",
		Addr:                 os.Getenv("HEROKU_DB_HOST"),
		DBName:               os.Getenv("HEROKU_DB_NAME"),
		AllowNativePasswords: true,
	}

	return conf
}

func localMySQLConf() mysql.Config {
	conf := mysql.Config{
		User:                 "root",
		Passwd:               os.Getenv("LOCAL_DB_ROOT_PASSWORD"),
		Net:                  "tcp",
		Addr:                 os.Getenv("LOCAL_DB_HOST"),
		DBName:               os.Getenv("LOCAL_DB_NAME"),
		AllowNativePasswords: true,
	}

	return conf
}
