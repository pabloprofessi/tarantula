package router

import (
	"fmt"
	"github.com/pprofessi/client"
	"github.com/pprofessi/response_writer"
	"net/http"
)

func Router(w http.ResponseWriter, r *http.Request) {
	rkw := redirectable(r.URL.Path[1:])
	if rkw.KeyWord != "" {
		client.ClientBackendRequester(w, r, rkw.DestinyRouteString)
	}
	response_writer.Response_writer(w, rkw.DestinyRouteString)

}

func redirectable(url_path string) RoutableKeyWord {
	var rkw RoutableKeyWord
	db := get_db()
	db.Where("key_word = ?", url_path).First(&rkw)
	fmt.Println(rkw.KeyWord)
	return rkw
}
