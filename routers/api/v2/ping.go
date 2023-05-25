package v2

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//cookie测试路由
func Ping() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"code": "200",
			"ping": "pong",
		})
	}
}
