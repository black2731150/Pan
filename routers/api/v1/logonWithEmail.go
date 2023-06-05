package v1

import (
	"pan/global"
	"pan/models"
	"pan/pkg/app"
	"pan/pkg/errcode"
	"pan/pkg/token"
	"strconv"

	"github.com/gin-gonic/gin"
)

//这个接口用于用户用户使用邮箱和邮箱验证码登录
func LoginWithEmail() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := models.NewUser()
		user.Email = ctx.PostForm("email")
		code := ctx.PostForm("emailcode")
		c, ok := global.GetEmailCodeFromMap(user.Email)
		response := app.NewRespponse(ctx)
		if user.HaveTheEmail() || ok || c == code {
			err := user.GetUserIDFromEmail()
			if err != nil {
				response.ToErrorResponse(errcode.ServerError)
				return
			}
			token, err := token.GenerateToken(user.UserName, user.Email, user.ID)
			if err != nil {
				data := gin.H{
					"code":    0,
					"details": "Login Success",
				}
				user.GetEmailFromUserName()
				maxAge := 60 * 60 * 24 * 7
				tokenValue := "token=" + token + "; Path=/; Max-Age=" + strconv.Itoa(maxAge) + "; SameSite=None"
				ctx.Header("Set-Cookie", tokenValue)
				// ctx.SetCookie("token", token, 60*60*24*7, "/", "", false, true)

				data["userid"] = user.ID
				response.ToResponse(data)
			} else {
				response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
			}
			response.ToErrorResponse(errcode.Success)
		} else {
			err := errcode.InbalidParams
			err.WithDetails("没有注册此邮箱")
			response.ToErrorResponse(err)
		}
	}
}
