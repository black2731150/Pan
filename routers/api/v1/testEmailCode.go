package v1

import (
	"pan/global"
	"pan/pkg/errcode"
	"pan/pkg/response"

	"github.com/gin-gonic/gin"
)

//这个接口用于前端检测验证码是否正确
func TestEmailCode() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		code := ctx.PostForm("code")
		email := ctx.PostForm("email")
		response := response.NewRespponse(ctx)

		if RealCode, ok := global.GetEmailCodeFromMap(email); ok && RealCode == code {
			response.ToErrorResponse(errcode.Success)
		} else {
			response.ToErrorResponse(errcode.InbalidParams)
		}
	}
}
