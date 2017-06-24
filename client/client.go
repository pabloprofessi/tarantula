package client

import (
	"crypto/tls"
	"github.com/tarantula/config"
	"io"
	"net/http"
	"time"
)

var timeout = time.Duration(300 * time.Second)

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

	res, err := client.Do(r)
	//config.LOG.Debugf("BE response: %#v", res)
	if err != nil {
		config.LOG.Errorf("failed to send request to BE")
		config.LOG.Errorf("Writing response error: %s", err)
	}
	defer res.Body.Close()
	config.LOG.Debugf("HEADERS:\n")
	for k, v := range res.Header {
		config.LOG.Debugf("%s : %s", k, v)
		w.Header()[k] = v
	}
	w.WriteHeader(res.StatusCode)
	io.Copy(w, res.Body)
	return
}
