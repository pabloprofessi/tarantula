package router

import (
	"fmt"
	"github.com/pprofessi/response_writer"
	"net/http"
)

func Router(w http.ResponseWriter, r *http.Request) {

	var s string
	s = check_path(r.URL.Path[:])

	response_writer.Response_writer(w, s)

}

func check_path(url_path string) string {

	var routable string = ""
	db := get_db()
	//db.Create(&RoutableKeyWord{KeyWord: url_path[1:]})
	var rkw RoutableKeyWord
	//db.Where("key_word = ?", url_path[1:]).First(&rkw)
	db.First(&rkw)
	fmt.Println(rkw)
	fmt.Println(rkw.KeyWord)
	return routable
}
