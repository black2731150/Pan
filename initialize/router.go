package initialize

import (
	"pan/global"
	"pan/routers"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	rootGroup := router.Group("/")
	routers.SetRootGroupRouters(rootGroup)

	return router
}

func RunServer() {
	router := setupRouter()
	router.Run(global.Panserver.Config.Pan.Host + ":" + global.Panserver.Config.Pan.Port)
}
