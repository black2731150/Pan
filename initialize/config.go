package initialize

import (
	"fmt"
	"os"
	"pan/global"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func IniterlizeConfig() *viper.Viper {
	//配置文件路径
	config := "config.yaml"
	if configEnv := os.Getenv("VIPER_CONFIG"); configEnv != "" {
		config = configEnv
	}

	//初始化viper
	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("Read config faild: %s", err))
	}
	fmt.Println("host", v.GetString("pan.port"))

	//监听配置文件
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("Config file changed: ", in.Name)
		//重载配置
		if err := v.Unmarshal(&global.Panserver.Config); err != nil {
			fmt.Println(err)
		}
	})

	//配置文件赋值给全局变量
	if err := v.Unmarshal(&global.Panserver.Config); err != nil {
		panic(fmt.Sprintf("Unmarshal config faild: %v", err))
	}

	fmt.Println("Config initialize SUCCESS!")
	return v
}
