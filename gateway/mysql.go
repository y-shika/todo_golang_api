package gateway

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/y-shika/todo_golang_api/config"
)

// MySQLClient is a sql.DB wrapper.
type MySQLClient struct {
	*sql.DB
}

// NewMySQLClient is a new MySQL client.
func NewMySQLClient() (*MySQLClient, error) {
	mysqlConf := mysql.Config{}

	// TODO: DBの接続先が増えたらこの部分の条件分岐も増やす
	if config.IsHeroku() || config.IsConnectedHerokuDB() {
		log.Println("Connect to HerokuDB")
		mysqlConf = herokuMySQLConf()
	} else {
		log.Println("Connect to LocalDB")
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
		ParseTime:            true,
		Loc:                  time.UTC,
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
		ParseTime:            true,
		Loc:                  time.UTC,
	}

	return conf
}
