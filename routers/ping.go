package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//测试功能
func ping() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"code": "200",
			"ping": "pong",
		})
	}
}
