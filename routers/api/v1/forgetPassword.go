package v1

import (
	"pan/models"
	"pan/pkg/app"
	"pan/pkg/email"
	"pan/pkg/errcode"
	"pan/utils"

	"github.com/gin-gonic/gin"
)

func ForgetPassword() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		response := app.NewRespponse(ctx)

		e := ctx.PostForm("email")
		code := ctx.PostForm("code")
		newPassword := ctx.PostForm("newpassword")
		newPassword = utils.StringMD5(newPassword)
		user := models.NewUser()
		user.Email = e

		if email.TestTheEmailCode(e, code) {
			err := user.UpdatePassword(newPassword)
			if err != nil {
				response.ToErrorResponse(errcode.InbalidParams.WithDetails("更新密码失败"))
			}
			data := gin.H{
				"code":    0,
				"message": "Replace password success!",
			}
			response.ToResponse(data)
		} else {
			response.ToErrorResponse(errcode.TheEmailCodeError.WithDetails("验证码校验不正确"))
			return
		}
	}
}
