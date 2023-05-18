package routers

import (
	"net/http"
	"pan/dao"

	"github.com/gin-gonic/gin"
)

func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := dao.NewUser()
		user.UserName = ctx.PostForm("username")
		user.Password = ctx.PostForm("password")

		if user.HaveUser() {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    100,
				"details": "Login Success",
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    300,
				"details": "Login Faild",
			})
		}
	}
}
