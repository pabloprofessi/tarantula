package router

import (
	"github.com/pprofessi/client"
	"github.com/pprofessi/config"
	"github.com/pprofessi/response_writer"
	"net/http"
)

func Router(w http.ResponseWriter, r *http.Request) {
	rkw := redirectable(r.URL.Path[1:])
	if rkw.KeyWord != "" {
		config.LOG.Debugf("Redirecting: %s", rkw.KeyWord)
		client.ClientBackendRequester(w, r, rkw.DestinyRouteString)
	}
	config.LOG.Debugf("Writing response: %#v", w)
	response_writer.Response_writer(w, rkw.DestinyRouteString)

}

func redirectable(url_path string) RoutableKeyWord {
	var rkw RoutableKeyWord
	db := get_db()
	db.Where("key_word = ?", url_path).First(&rkw)
	config.LOG.Debugf("Routable key-word obtained from DB querys: %s", rkw.KeyWord)
	return rkw
}
