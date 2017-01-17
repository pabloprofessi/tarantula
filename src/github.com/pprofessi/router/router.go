package router

import (
	"fmt"
	"github.com/pprofessi/client"
	"github.com/pprofessi/response_writer"
	"net/http"
)

func Router(w http.ResponseWriter, r *http.Request) {

	var destinyRouteString string
	destinyRouteString = redirectable(r.URL.Path[1:])
	if destinyRouteString != "" {
		w = client.ClientBackendRequester(r, destinyRouteString)
	}
	response_writer.Response_writer(w, destinyRouteString)

}

func redirectable(url_path string) string {

	var rkw RoutableKeyWord
	db := get_db()
	db.Where("key_word = ?", url_path).First(&rkw)
	fmt.Println(rkw.KeyWord)
	return rkw.KeyWord
}
