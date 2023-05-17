package global

import (
	"pan/config"

	"github.com/spf13/viper"
)

type PanServer struct {
	ConfigViper *viper.Viper
	Config      config.Configuration
}

var Panserver = new(PanServer)
