package router

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pprofessi/config"
	//º"fmt"
)

var urlsdb *sql.DB

func db_init() {
	var err error
	urlsdb, err = sql.Open("mysql", config.Config.DbConectionString)

	if err != nil {
		config.LOG.Errorf("Failed to connect to database, err: %s", err)
	}

}

func get_db() *sql.DB {
	if urlsdb == nil {
		db_init()
		config.LOG.Debugf("Gorm DB initialized!")
	}
	return urlsdb
}
