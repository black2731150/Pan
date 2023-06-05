package routers

import (
	"net/http"
	"pan/pkg/middlieware"
	v1 "pan/routers/api/v1"
	v2 "pan/routers/api/v2"

	"github.com/gin-gonic/gin"
)

//根路由管理
func SetRootGroupRouters(router *gin.RouterGroup) {

	//中间件
	router.Use(middlieware.Cors())

	//静态文件
	router.StaticFS("/web", http.Dir("./web"))
	router.StaticFile("/", "./web/index.html")
	router.StaticFile("/requests", "./request.log")

	//APIV1 路由组
	apiV1Group := router.Group("/api/v1")
	setAPIV1GroupRouters(apiV1Group)

	//APIV2 路由组
	apiV2Geoup := router.Group("/api/v2")
	apiV2Geoup.Use(middlieware.JWT())
	setAPIV2GroupRouters(apiV2Geoup)

	//user 路由组
	userGroup := router.Group("/user")
	setUserGroupRouters(userGroup)
}

//api/v1 路由管理
func setAPIV1GroupRouters(router *gin.RouterGroup) {
	router.Use(middlieware.RequestShow())
	RegitsterNewRouter(router, "POST", "/login", v1.Login())
	RegitsterNewRouter(router, "POST", "/register", v1.Register())
	RegitsterNewRouter(router, "POST", "/testUserName", v1.TestUserName())
	RegitsterNewRouter(router, "POST", "/testEmailCode", v1.TestEmailCode())
	RegitsterNewRouter(router, "POST", "/loginWithEmail", v1.LoginWithEmail())
	RegitsterNewRouter(router, "POST", "/sendEmailCode", v1.SendEmailCode())
	RegitsterNewRouter(router, "POST", "/forgetpasswd", v1.ForgetPassword())
}

//api/v2 路由管理
func setAPIV2GroupRouters(router *gin.RouterGroup) {
	router.Use(middlieware.RequestShow())
	RegitsterNewRouter(router, "GET", "/ping", v2.Ping())
	RegitsterNewRouter(router, "POST", "/download", v2.Download())
	RegitsterNewRouter(router, "POST", "/upload", v2.Upload())
	RegitsterNewRouter(router, "GET", "/getfilelist", v2.GetFileList())
	RegitsterNewRouter(router, "DELETE", "/deletefile", v2.DeleteFile())
	RegitsterNewRouter(router, "PATCH", "/updatefilename", v2.UpdateFileName())
	RegitsterNewRouter(router, "POST", "/replacepasswd", v2.ReplacePassword())
	RegitsterNewRouter(router, "GET", "/createnewfolder", v2.CreateNewFolder())
}

//user 路由管理
func setUserGroupRouters(router *gin.RouterGroup) {
	router.GET("/:id", nil)
}

//自动注册OPTIONS
func RegitsterNewRouter(router *gin.RouterGroup, method string, path string, handleFunc gin.HandlerFunc) {
	router.Handle(method, path, handleFunc)
	router.OPTIONS(path)
}
