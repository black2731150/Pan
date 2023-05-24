package global

import (
	"pan/config"
	"pan/pkg/logger"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

//pan服务全局变量
type PanServer struct {
	ConfigViper *viper.Viper
	Config      config.Configuration
	DB          *gorm.DB
	Logger      *logger.Logger
}

var Panserver = new(PanServer)
