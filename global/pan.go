package global

import (
	"pan/config"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type PanServer struct {
	ConfigViper *viper.Viper
	Config      config.Configuration
	DB          *gorm.DB
}

var Panserver = new(PanServer)
