package main

import (
	//"crypto/tls"
	"github.com/pprofessi/router"
	"net/http"
)

func proxy(w http.ResponseWriter, r *http.Request) {
	router.Router(w, r)

}

func main() {

	mux := http.NewServeMux()
	proxy_handler := http.HandlerFunc(proxy)
	mux.Handle("/", proxy_handler)

	s := &http.Server{
		Addr:    ":9999",
		Handler: proxy_handler,
	}

	s.ListenAndServe()
}
