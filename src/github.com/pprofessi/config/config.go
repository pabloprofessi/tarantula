package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
)

type ConfigType struct {
	PathToRedirect string
	ProxyDNS       string
}

var Config ConfigType

func init() {

	//setar con variable de entorno
	if _, err := toml.DecodeFile(os.Getenv("TARANTULA_CONF"), &Config); err != nil {
		fmt.Println(err)
		return
	}

	return
}
