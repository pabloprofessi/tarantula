package response_writer

import (
	"fmt"
	"github.com/tarantula/config"
	"net/http"
)

func Response_writer(w http.ResponseWriter, s string, r *http.Request) {

	if s == "pong" {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(200)
		fmt.Fprintf(w, "%s", s)
		return
	}

	target := "https://" + config.Config.ToDomain + "/page-not-found"
	if len(r.URL.RawQuery) > 0 {
		target += "?" + r.URL.RawQuery
	}
	http.Redirect(w, r, target, http.StatusPermanentRedirect)
	config.LOG.Infof("Not to be routed!")
}
