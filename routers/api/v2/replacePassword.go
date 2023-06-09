package v2

import (
	"pan/models"
	"pan/pkg/app"
	"pan/pkg/errcode"
	"pan/utils"

	"github.com/gin-gonic/gin"
)

//重置密码
func ReplacePassword() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response := app.NewRespponse(ctx)
		// oldpassword := ctx.PostForm("oldpassword")
		// oldpassword = utils.StringMD5(oldpassword)

		newpassword := ctx.PostForm("newpassword")
		newpassword = utils.StringMD5(newpassword)

		user := models.NewUser()
		ID, ok := ctx.Get("UserID")
		if ok {
			user.ID = ID.(uint)
		} else {
			response.ToErrorResponse(errcode.InbalidParams.WithDetails("UserID获取失败"))
			return
		}

		user.GetEmailAndUsernameFromID()
		err := user.UpdatePassword(newpassword)

		if err != nil {
			response.ToErrorResponse(errcode.ServerError.WithDetails("更新密码失败"))
			return
		} else {
			response.ToSuccessResponse("Replace password success!")
			return
		}
	}
}
