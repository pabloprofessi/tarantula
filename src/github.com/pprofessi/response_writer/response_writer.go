package response_writer

import (
	"fmt"
	"net/http"
)

func Response_writer(w http.ResponseWriter, s string) {

	if s == "" {
		fmt.Fprintf(w, "Url: %s \n", "source not found")
	} else {
		fmt.Fprintf(w, "Url: %s \n", s)
	}

}
