package router

import (
	"github.com/pprofessi/client"
	"github.com/pprofessi/config"
	"github.com/pprofessi/response_writer"
	"net/http"
)

func Router(w http.ResponseWriter, r *http.Request) {
	config.LOG.Debugf("r.URL.String(): %s", r.URL.String())
	config.LOG.Debugf("r.URL.Path[1:]: %s", r.URL.Path[1:])
	config.LOG.Debugf("r.Host:  %s", r.Host)

	if r.URL.String() == "/" {
		client.ClientBackendRequester(w, r)
	} else {
		response_writer.Response_writer(w)
	}
	//rkw := redirectable(r.URL.Path[1:])
	//if rkw.KeyWord != "" {
	//	config.LOG.Debugf("url redirected: %s", rkw.DestinyRouteString)
	//	config.LOG.Debugf("Redirecting to: %s", rkw.KeyWord)
	//	client.ClientBackendRequester(w, r, rkw.DestinyRouteString)
	//} else {
	//}

}

func redirectableUri(urlPath string) RouteToKeyWord {
	var rtkw RouteToKeyWord
	db := get_db()
	db.Where("key_word = ?", urlPath).First(&rtkw)
	config.LOG.Debugf("Route to key-word obtained from DB query: %s", rtkw.KeyWord)
	return rtkw
}

func redirectableHost(sourceHost string, sourceScheme string) RouteToHost {
	var rth RouteToHost
	db := get_db()
	db.Where("source_host = ? AND source_scheme = ?", sourceHost, sourceScheme).First(&rth)
	//config.LOG.Debugf("Route to host obtained from DB query: %s", rth.sourceHost)
	return rth
}
