package routers

import (
	"pan/dao"
	"pan/pkg/errcode"
	"pan/pkg/response"

	"github.com/gin-gonic/gin"
)

//登录api
func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := dao.NewUser()
		user.UserName = ctx.PostForm("username")
		user.Password = ctx.PostForm("password")

		response := response.NewRespponse(ctx)

		if user.HaveUser() {
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
