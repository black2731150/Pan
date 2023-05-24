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

	//初始化日志
	initialize.InitLog()

	//初始化数据库
	initialize.InitDB()
	defer func() {
		if global.Panserver.DB != nil {
			db, _ := global.Panserver.DB.DB()
			db.Close()
		}
	}()

	//启动服务
	initialize.RunServer()
}
