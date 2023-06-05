package v1

import (
	"pan/models"
	"pan/pkg/app"
	"pan/pkg/errcode"
	"pan/pkg/token"
	"pan/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

//这个接口用于用户使用用户名和密码登录
func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := models.NewUser()
		user.UserName = ctx.PostForm("username")
		user.Password = utils.StringMD5(ctx.PostForm("password"))

		response := app.NewRespponse(ctx)

		if user.UserNameLogin() {
			err := user.GetEmailFromUserName()
			if err != nil {
				response.ToErrorResponse(errcode.ServerError)
				return
			}
			data := gin.H{
				"code":    0,
				"details": "Login Success",
			}
			err = user.GetUserIDFromUsername()
			if err != nil {
				response.ToErrorResponse(errcode.ServerError)
				return
			}
			token, err := token.GenerateToken(user.UserName, user.Email, user.ID)
			if err == nil {
				maxAge := 60 * 60 * 24 * 7
				tokenValue := "token=" + token + "; Path=/; Max-Age=" + strconv.Itoa(maxAge) + "; SameSite=None"
				ctx.Header("Set-Cookie", tokenValue)
				// ctx.SetCookie("token", token, 60*60*24*7, "/", "", false, true)

				data["userid"] = user.ID
				response.ToResponse(data)

			} else {
				response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
			}
		} else {
			err := errcode.InbalidParams
			err.WithDetails("用户名或者密码错误")
			response.ToErrorResponse(err)
		}
	}
}
