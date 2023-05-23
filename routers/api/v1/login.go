package v1

import (
	"fmt"
	"pan/common"
	"pan/dao"
	"pan/pkg/errcode"
	"pan/pkg/response"

	"github.com/gin-gonic/gin"
)

//登录api
func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println(ctx.Request)
		user := dao.NewUser()
		user.UserName = ctx.PostForm("username")
		user.Password = ctx.PostForm("password")
		user.Password = common.StringMD5(user.Password)

		response := response.NewRespponse(ctx)

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
