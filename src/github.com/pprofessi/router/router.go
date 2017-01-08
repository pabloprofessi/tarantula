package router

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pprofessi/response_writer"
	"net/http"
)

type RoutableKeyWord struct {
	gorm.Model
	KeyWord string
}

func Router(w http.ResponseWriter, r *http.Request) {

	var s string
	s = check_path(r.URL.Path[:])

	response_writer.Response_writer(w, s)

}

func check_path(url_path string) string {

	var routable string = ""
	db, err := gorm.Open("mysql", "tarantula:tarantula@tcp(db:3306)/tarantula")
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	db.AutoMigrate(&RoutableKeyWord{})
	//db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&RoutableKeyWord{})
	db.Create(&RoutableKeyWord{KeyWord: url_path})

	return routable
}
