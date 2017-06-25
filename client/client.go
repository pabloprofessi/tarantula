package client

import (
	"bytes"
	"compress/gzip"
	"crypto/tls"
	"github.com/tarantula/config"
	"io"
	"io/ioutil"
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

	config.PrettyRequestLoger(r, "cli request")
	res, err := client.Do(r)
	config.PrettyResponseLoger(res, "cli response")
	//config.LOG.Debugf("BE response: %#v", res)
	if err != nil {
		config.LOG.Errorf("failed to send request to BE")
		config.LOG.Errorf("Writing response error: %s", err)
	}
	defer res.Body.Close()

	var body []byte

	w.WriteHeader(res.StatusCode)
	if res.Header.Get("Content-Encoding") == "gzip" {
		reader, err := gzip.NewReader(res.Body)
		if err != nil {
			config.LOG.Errorf("reading body to be gziped replaced")
			config.LOG.Errorf("Error: %s", err)
		}
		body, err = ioutil.ReadAll(reader)
	} else {
		body, err = ioutil.ReadAll(res.Body)
	}

	if err != nil {
		config.LOG.Errorf("error casting body to byte[]")
		config.LOG.Errorf("Error: %s", err)
	}
	//config.LOG.Debugf("body: %s", body)
	body = bytes.Replace(body, []byte(`<meta name="robots" content="noindex,follow"/>`), []byte(""), -1)
	//config.LOG.Debugf("body_replaced: %s", body)
	buf := bytes.NewBuffer(body)

	for k, v := range res.Header {
		if (k != "Accept-Encoding") && (k != "Content-Encoding") {
			w.Header()[k] = v
			config.LOG.Debugf("REWIRTED headers %s : %s", k, v)
		}

	}
	io.Copy(w, buf)
	return
}
