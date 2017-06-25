package client

import (
	"bytes"
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

	res, err := client.Do(r)
	//config.LOG.Debugf("BE response: %#v", res)
	if err != nil {
		config.LOG.Errorf("failed to send request to BE")
		config.LOG.Errorf("Writing response error: %s", err)
	}
	defer res.Body.Close()
	for k, v := range res.Header {
		//config.LOG.Debugf("%s : %s", k, v)
		w.Header()[k] = v
	}
	config.PrettyRequestLoger(r, "cli request")
	w.WriteHeader(res.StatusCode)
	body, _ := ioutil.ReadAll(res.Body)
	body = bytes.Replace(body, []byte(`<meta name="robots" content="noindex,follow"/>`), []byte(""), -1)
	buf := bytes.NewBuffer(body)
	io.Copy(w, buf)
	return
}
