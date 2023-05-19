package routers

import (
	"pan/dao"

	"github.com/gin-gonic/gin"
)

//注册api
func Register() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := dao.NewUser()
		if user.UserName = ctx.PostForm("username"); len(user.UserName) == 0 {
			return
		}
		if user.Password = ctx.PostForm("password"); len(user.Password) < 6 || len(user.Password) > 20 {
			return
		}

		user.Email = ctx.PostForm("email")
		user.Phonenum = ctx.PostForm("phonenum")
		user.RegisterNewUser()
	}
}
