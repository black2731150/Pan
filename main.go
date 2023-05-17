package main

import (
	"fmt"
	"pan/global"
	"pan/initialize"
)

func main() {
	//初始化配置文件
	initialize.IniterlizeConfig()
	fmt.Println(global.Panserver.Config)

	//启动服务
	initialize.RunServer()
}
