package v1

import (
	"pan/global"
	"pan/pkg/app"
	"pan/pkg/errcode"

	"github.com/gin-gonic/gin"
)

//这个接口用于前端检测验证码是否正确
func TestEmailCode() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		code := ctx.PostForm("code")
		email := ctx.PostForm("email")
		response := app.NewRespponse(ctx)

		if RealCode, ok := global.GetEmailCodeFromMap(email); ok && RealCode == code {
			response.ToErrorResponse(errcode.Success)
		} else {
			response.ToErrorResponse(errcode.InbalidParams)
		}
	}
}
