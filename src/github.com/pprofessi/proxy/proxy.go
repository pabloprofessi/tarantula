package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var tr = &http.Transport{}
var client = &http.Client{Transport: tr}

func proxy_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Url: %s \n", r.URL.Path[:])

	server_url, _ := url.Parse("http://127.0.0.1:9900" + r.URL.String())

	r.Host = "localhost"
	r.RequestURI = ""
	r.URL = server_url

	res, _ := client.Do(r)
	defer res.Body.Close()
	for k, v := range res.Header {
		w.Header()[k] = v
	}
	w.WriteHeader(res.StatusCode)
	io.Copy(w, res.Body)

}

func main() {

	http.HandleFunc("/", proxy_handler)
	http.ListenAndServe(":9999", nil)
}
