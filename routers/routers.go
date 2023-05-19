package routers

import "github.com/gin-gonic/gin"

//跟路由管理
func SetRootGroupRouters(router *gin.RouterGroup) {
	//ping测试路由
	router.GET("/ping", ping())

	//静态文件
	router.StaticFile("/", "/home/PAN/web/mypan/src/index.js")

	//API路由组
	apiV1Group := router.Group("/api/v1")
	setAPIGroupRouters(apiV1Group)
}

//api路由管理
func setAPIGroupRouters(router *gin.RouterGroup) {
	router.POST("/login", Login())
	router.POST("/register", Register())
}
