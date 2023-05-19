package global

import (
	"pan/config"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

//pan服务全局变量
type PanServer struct {
	ConfigViper *viper.Viper
	Config      config.Configuration
	DB          *gorm.DB
}

var Panserver = new(PanServer)
