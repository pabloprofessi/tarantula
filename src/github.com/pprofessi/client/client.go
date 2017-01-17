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

func ClientBackendRequester(r *http.Request, destinyRouteString string) http.ResponseWriter {
	var w http.ResponseWriter
	backend_server_url, _ := url.Parse(destinyRouteString)

	r.Host = config.Config.ProxyDNS //deberia el DNS del proxy
	r.RequestURI = "*"              //no se si esto va a molestar
	r.URL = backend_server_url
	//TODO: cachiar el error
	res, err := client.Do(r)
	if err != nil {
		fmt.Println(err)
		panic("failed to send request to BE")
	}
	defer res.Body.Close()
	for k, v := range res.Header {
		w.Header()[k] = v
	}
	w.WriteHeader(res.StatusCode)
	io.Copy(w, res.Body)
	return w
}
