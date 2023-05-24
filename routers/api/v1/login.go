package v1

import (
	"fmt"
	"pan/dao"
	"pan/pkg/app"
	"pan/pkg/errcode"
	"pan/utils"

	"github.com/gin-gonic/gin"
)

//这个接口用于用户使用用户名和密码登录
func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println(ctx.Request)
		user := dao.NewUser()
		user.UserName = ctx.PostForm("username")
		user.Password = utils.StringMD5(ctx.PostForm("password"))

		response := app.NewRespponse(ctx)

		if user.UserNameLogin() {
			data := gin.H{
				"code":    0,
				"details": "Login Success",
			}
			response.ToResponse(data)
		} else {
			err := errcode.InbalidParams
			response.ToErrorResponse(err)
		}
	}
}
