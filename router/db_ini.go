package router

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tarantula/config"
)

var urlsdb *sql.DB = nil

func db_init() *sql.DB {

	db, err := sql.Open("mysql", config.Config.DbConectionString)

	if err != nil {
		config.LOG.Errorf("Failed to connect to database, err: %s", err)
	}

	return db
}

func get_db() *sql.DB {

	if urlsdb != nil {
		return urlsdb
	}

	urlsdb = db_init()
	config.LOG.Debugf("DB initialized!")
	return urlsdb
}
