package routers

import "github.com/gin-gonic/gin"

func SetRootGroupRouters(router *gin.RouterGroup) {
	//ping测试路由
	router.GET("/ping", ping())
}