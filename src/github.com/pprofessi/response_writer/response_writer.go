package response_writer

import (
	"fmt"
	"github.com/pprofessi/config"
	"net/http"
	"os"
)

func Response_writer(w http.ResponseWriter, s string) {

	if s == "" {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(404)
		fmt.Fprintf(w, "Url: %s \n", "Source not found!")
		fmt.Println("404 Source not found!")
	} else {
		//fmt.Fprintf(w, "Url: %s \n", s)
		fmt.Fprintf(os.Stdout, "Proxy to, url file: %s%s \n", config.Config.PathToRedirect, s)
	}

}
