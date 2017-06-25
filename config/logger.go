package config

import (
	"github.com/op/go-logging"
	//"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var LOG = logging.MustGetLogger("config")

func Set_logger() {

	f, err := os.OpenFile(Config.LogOutputPath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	var format = logging.MustStringFormatter(Config.LogFormat)
	backend := logging.NewLogBackend(f, "", 0)
	backend2Formatter := logging.NewBackendFormatter(backend, format)
	logging.SetLevel(logging.GetLevel(Config.LogLevel), "config")
	logging.SetBackend(backend2Formatter)
}

func PrettyRequestLoger(r *http.Request, prefix string) {

	headers_line := ""

	for k, v := range r.Header {
		headers_line = "[ " + k + " : " + strings.Join(v[:], ",") + " ]" + headers_line
	}

	LOG.Infof("%s > RemoteAddr: %s, URL: %s, Proto: %s", prefix, r.RemoteAddr, r.URL, r.Proto)
	LOG.Infof("%s > HEADERS : %s", prefix, headers_line)
	//body, _ := ioutil.ReadAll(r.Body)
	//LOG.Infof("%s > BODY: %s", prefix, body)

}

func PrettyResponseLoger(r *http.Response, prefix string) {

	headers_line := ""

	for k, v := range r.Header {
		headers_line = "[ " + k + " : " + strings.Join(v[:], ",") + " ]" + headers_line
	}
	LOG.Infof("%s > Status: %s, Proto: %s", prefix, r.Status, r.Proto)
	LOG.Infof("%s > HEADERS : %s", prefix, headers_line)
	//body, _ := ioutil.ReadAll(r.Body)
	//LOG.Infof("%s > BODY: %s", prefix, body)

}

//config.LOG.Debug("debug")
//config.LOG.Info("info")
//config.LOG.Notice("notice")
//config.LOG.Warning("warning")
//config.LOG.Error("err")
//config.LOG.Critical("crit")
