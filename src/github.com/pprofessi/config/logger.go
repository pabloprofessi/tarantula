package config

import (
	"github.com/op/go-logging"
	"os"
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

// hacer import de: "github.com/pprofessi/config"
//config.LOG.Debug("debug")
//config.LOG.Info("info")
//config.LOG.Notice("notice")
//config.LOG.Warning("warning")
//config.LOG.Error("err")
//config.LOG.Critical("crit")
