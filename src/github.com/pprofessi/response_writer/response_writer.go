package response_writer

import (
	"fmt"
	"github.com/pprofessi/config"
	"net/http"
)

func Response_writer(w http.ResponseWriter, s string) {

	if s == "" {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(404)
		fmt.Fprintf(w, "Url: %s \n", "Source not found!")
		config.LOG.Infof("404 Source not found!")
	} else {
		config.LOG.Infof("Request to %s proxied!", s)
	}

}
