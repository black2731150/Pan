package routers

import (
	"net/http"
	"pan/pkg/middlieware"
	v1 "pan/routers/api/v1"

	"github.com/gin-gonic/gin"
)

//跟路由管理
func SetRootGroupRouters(router *gin.RouterGroup) {

	//中间件
	router.Use(middlieware.Cors())

	//ping测试路由
	router.GET("/ping", ping())

	//静态文件
	router.StaticFS("/web", http.Dir("./web"))
	router.StaticFile("/", "./web/index.html")

	//API路由组
	apiV1Group := router.Group("/api/v1")
	setAPIGroupRouters(apiV1Group)
}

//api路由管理
func setAPIGroupRouters(router *gin.RouterGroup) {
	RegitsterNewRouter(router, "POST", "/login", v1.Login())
	RegitsterNewRouter(router, "POST", "/register", v1.Register())
	RegitsterNewRouter(router, "POST", "/testUserName", v1.TestUserName())
	RegitsterNewRouter(router, "POST", "/testEmailCode", v1.TestEmailCode())
	RegitsterNewRouter(router, "POST", "/loginWithEmail", v1.LoginWithEmail())
}

//自动注册OPTIONS
func RegitsterNewRouter(router *gin.RouterGroup, method string, path string, handleFunc gin.HandlerFunc) {
	router.Handle(method, path, handleFunc)
	router.OPTIONS(path)
}
