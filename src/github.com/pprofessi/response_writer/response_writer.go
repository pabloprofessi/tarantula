package response_writer

import (
	"fmt"
	"github.com/pprofessi/config"
	"net/http"
)

func Response_writer(w http.ResponseWriter) {

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(404)
	fmt.Fprintf(w, "Url: %s \n", "Source not found!")
	config.LOG.Infof("404 Source not found!")
}
