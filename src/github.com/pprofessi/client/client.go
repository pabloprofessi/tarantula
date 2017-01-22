package client

import (
	"github.com/pprofessi/config"
	"io"
	"net/http"
	"net/url"
)

var tr = &http.Transport{}
var client = &http.Client{Transport: tr}

func ClientBackendRequester(w http.ResponseWriter, r *http.Request, destinyRouteString string) {
	destinyRouteStringParsed, err := url.Parse(destinyRouteString)
	r.URL = destinyRouteStringParsed
	if err != nil {
		config.LOG.Errorf("failed parse destinyRouteString")
		config.LOG.Errorf("Writing response: %s", err)
	}

	r.Host = config.Config.ProxyDNS
	config.LOG.Infof("Redirecting request to: ", r.URL)
	r.RequestURI = ""

	res, err := client.Do(r)
	config.LOG.Debugf("BE ressponse: %#v", res)
	if err != nil {
		config.LOG.Errorf("failed to send request to BE")
		config.LOG.Errorf("Writing response: %s", err)
	}
	defer res.Body.Close()
	for k, v := range res.Header {
		w.Header()[k] = v
	}
	w.WriteHeader(res.StatusCode)
	io.Copy(w, res.Body)
	return
}
