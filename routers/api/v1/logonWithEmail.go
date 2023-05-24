package v1

import (
	"pan/dao"
	"pan/global"
	"pan/pkg/errcode"
	"pan/pkg/response"

	"github.com/gin-gonic/gin"
)

//这个接口用于用户用户使用邮箱和邮箱验证码登录
func LoginWithEmail() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := dao.NewUser()
		user.Email = ctx.PostForm("email")
		code := ctx.PostForm("emailcode")
		c, ok := global.GetEmailCodeFromMap(user.Email)
		response := response.NewRespponse(ctx)
		if user.HaveTheEmail() || ok || c == code {
			response.ToErrorResponse(errcode.Success)
		} else {
			err := errcode.InbalidParams
			err.WithDetails("没有注册此邮箱")
			response.ToErrorResponse(err)
		}
	}
}
