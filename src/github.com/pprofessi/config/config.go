package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
)

type ConfigType struct {
	LogOutputPath     string
	LogLevel          string
	LogFormat         string
	DbConectionString string
	FromDomain        string
	ToDomain          string
}

var Config ConfigType

func init() {
	conf_path := "/app/tarantula/config/tarantula_config_dev.conf"

	if os.Getenv("ENV") == "prod" {
		conf_path = "/app/tarantula/config/tarantula_config_prod.conf"
	}
	_, err := toml.DecodeFile(conf_path, &Config)

	if err != nil {
		fmt.Println(err)
		return
	}

	Set_logger()

	return
}
