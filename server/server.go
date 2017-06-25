package server

import (
	"github.com/tarantula/config"
	"github.com/tarantula/router"
	"net/http"
)

func proxy(w http.ResponseWriter, r *http.Request) {
	router.Router(w, r)

}

func redirect(w http.ResponseWriter, r *http.Request) {

	target := "https://" + config.Config.FromDomain + r.URL.Path
	if len(r.URL.RawQuery) > 0 {
		target += "?" + r.URL.RawQuery
	}
	http.Redirect(w, r, target, http.StatusPermanentRedirect)
	config.LOG.Debugf("redirected!")
}

func InitServer() {

	mux := http.NewServeMux()
	proxy_handler := http.HandlerFunc(proxy)
	mux.Handle("/", proxy_handler)

	s := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: proxy_handler,
	}

	go s.ListenAndServe()

	redirect_mux := http.NewServeMux()
	redirect_handler := http.HandlerFunc(redirect)
	redirect_mux.Handle("/", redirect_handler)

	redirect_server := &http.Server{
		Addr:    "0.0.0.0:8081",
		Handler: redirect_handler,
	}
	redirect_server.ListenAndServe()
}
