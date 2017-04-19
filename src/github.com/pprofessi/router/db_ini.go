package router

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pprofessi/config"
)

type RouteToHost struct {
	gorm.Model
	SourceHost  string
	DestinyHost string
	RouteToUris []RouteToUri
}

type RouteToUri struct {
	gorm.Model
	SourceUri     string
	DestinyUri    string
	RouteToHostID uint
}

var kwdb *gorm.DB = nil

func db_init() {
	var err error
	kwdb, err = gorm.Open("mysql", config.Config.DbConectionString)
	if err != nil {
		config.LOG.Errorf("Failed to connect to database, err: %s", err)
	}

	kwdb.AutoMigrate(&RouteToHost{})
	kwdb.AutoMigrate(&RouteToUri{})
}

func get_db() *gorm.DB {
	if kwdb == nil {
		db_init()
		config.LOG.Debugf("Gorm DB initialized!")
	}
	return kwdb
}
