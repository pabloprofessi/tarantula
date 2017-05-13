package response_writer

import (
	"fmt"
	"github.com/pprofessi/config"
	"net/http"
)

func Response_writer(w http.ResponseWriter, s string) {

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	if s == "pong" {
		w.WriteHeader(200)
		fmt.Fprintf(w, "%s", s)
		return

	}

	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(404)
	fmt.Fprintf(w, "%s", s)
	config.LOG.Infof("404 Source not found!")
}
