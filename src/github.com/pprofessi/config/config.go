package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
)

type ConfigType struct {
	ProxyDNS          string
	LogOutputPath     string
	LogLevel          string
	LogFormat         string
	DbConectionString string
	FromDomain        string
	ToDomain          string
}

var Config ConfigType

func init() {
	if _, err := toml.DecodeFile(os.Getenv("TARANTULA_CONF"), &Config); err != nil {
		fmt.Println(err)
		return
	}
	Set_logger()

	return
}
