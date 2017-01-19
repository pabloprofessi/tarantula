package client

import (
	"fmt"
	"github.com/pprofessi/config"
	"io"
	"net/http"
	"net/url"
)

var tr = &http.Transport{}
var client = &http.Client{Transport: tr}

func ClientBackendRequester(w http.ResponseWriter, r *http.Request, destinyRouteString string) {
	//TODO: cachiar el error
	destinyRouteStringParsed, err := url.Parse(destinyRouteString)
	r.URL = destinyRouteStringParsed
	if err != nil {
		fmt.Println(err)
		fmt.Println("failed parse destinyRouteString")
	}

	r.Host = config.Config.ProxyDNS //deberia el DNS del proxy
	fmt.Println(" ---- r.URL --->> ", r.URL)
	r.RequestURI = ""
	r.Host = config.Config.ProxyDNS

	res, err := client.Do(r)
	//fmt.Printf("%#v\n", res)
	if err != nil {
		fmt.Println(err)
		fmt.Println("failed to send request to BE")
	}
	defer res.Body.Close()
	for k, v := range res.Header {
		w.Header()[k] = v
	}
	w.WriteHeader(res.StatusCode)
	io.Copy(w, res.Body)
	return
}
