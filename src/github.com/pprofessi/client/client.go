package client

import (
	"crypto/tls"
	"github.com/pprofessi/config"
	"io"
	"net/http"
	"net/url"
	"time"
	//"reflect"
)

var timeout = time.Duration(5 * time.Second)

var tr = &http.Transport{
	TLSClientConfig:     &tls.Config{RootCAs: nil},
	MaxIdleConnsPerHost: 65535,
}
var client = &http.Client{Transport: tr,
	CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	},
	Timeout: timeout,
}

func ClientBackendRequester(w http.ResponseWriter, r *http.Request) {

	destinyRouteURL, err := url.Parse("http://www.garbarino.com")
	if err != nil {
		config.LOG.Errorf("failed parse destinyRouteString")
		config.LOG.Errorf("Error: %s", err)
	}
	r.URL = destinyRouteURL
	r.Host = destinyRouteURL.Host
	r.RequestURI = ""

	res, err := client.Do(r)
	config.LOG.Debugf("BE response: %#v", res)
	if err != nil {
		config.LOG.Errorf("failed to send request to BE")
		config.LOG.Errorf("Writing response error: %s", err)
	}
	defer res.Body.Close()
	for k, v := range res.Header {
		config.LOG.Debugf("header %s", v)
		w.Header()[k] = v
	}
	w.WriteHeader(res.StatusCode)
	io.Copy(w, res.Body)
	return
}
