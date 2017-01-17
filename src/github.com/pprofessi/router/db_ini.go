package router

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type RoutableKeyWord struct {
	gorm.Model
	KeyWord            string
	DestinyRouteString string
}

var kwdb *gorm.DB = nil

func db_init() {
	var err error
	kwdb, err = gorm.Open("mysql", "tarantula:tarantula@tcp(db:3306)/tarantula?parseTime=true")
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	kwdb.AutoMigrate(&RoutableKeyWord{})
}

func get_db() *gorm.DB {
	if kwdb == nil {
		db_init()
		fmt.Println("gorm DB initialized")
	}
	return kwdb
}
