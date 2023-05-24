package v1

import (
	"pan/pkg/email"
	"pan/pkg/errcode"
	"pan/pkg/response"
	"pan/utils"

	"github.com/gin-gonic/gin"
)

//这个接口用于前端触发给用户发送验证码
func SendEmailCode() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		em := ctx.PostForm("email")
		code := utils.GetRandSixCode()
		err := email.SendEmail(em, "云盘验证码", code)
		response := response.NewRespponse(ctx)
		if err != nil {
			response.ToErrorResponse(errcode.ServerError)
		} else {
			response.ToErrorResponse(errcode.Success)
		}
	}
}
