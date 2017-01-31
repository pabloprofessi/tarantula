package client

import (
	"crypto/tls"
	"github.com/pprofessi/config"
	"io"
	"net/http"
	"net/url"
	//"time"
	//"reflect"
)

//var tr = &http.Transport{
//	TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
//	MaxIdleConnsPerHost: 65535,
//}

//var timeout = time.Duration(5 * time.Second)

//var client = &http.Client{Transport: tr, Timeout: timeout}
var tr = &http.Transport{
	TLSClientConfig: &tls.Config{RootCAs: nil},
}
var client = &http.Client{Transport: tr,
	CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	},
}

func ClientBackendRequester(w http.ResponseWriter, r *http.Request, destinyRouteString string) {

	//cert, err := tls.LoadX509KeyPair("/app/tarantula/src/github.com/pprofessi/server/server.crt", "/app/tarantula/src/github.com/pprofessi/server/server.key")
	//tlsConfig := &tls.Config{
	//	Certificates:       []tls.Certificate{cert},
	//	InsecureSkipVerify: true,
	//}

	destinyRouteURL, err := url.Parse(destinyRouteString)
	if err != nil {
		config.LOG.Errorf("failed parse destinyRouteString")
		config.LOG.Errorf("Writing response: %s", err)
	}
	r.URL = destinyRouteURL
	r.Host = destinyRouteURL.Host
	r.RequestURI = ""

	res, err := client.Do(r)

	config.LOG.Debugf("BE response: %#v", res)
	if err != nil {
		config.LOG.Errorf("failed to send request to BE")
		config.LOG.Errorf("Writing response: %s", err)
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
