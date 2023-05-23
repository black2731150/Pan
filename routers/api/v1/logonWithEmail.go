package v1

import (
	"pan/common"
	"pan/dao"
	"pan/global"
	"pan/pkg/email"
	"pan/pkg/errcode"
	"pan/pkg/response"
	"time"

	"github.com/gin-gonic/gin"
)

func LoginWithEmail() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := dao.NewUser()
		user.Email = ctx.PostForm("eamil")
		response := response.NewRespponse(ctx)
		if user.HaveTheEmail() {
			code := common.GetRandSixCode()
			email.SendEmail(user.Email, "验证码", code)
			global.AddToMap(user.Email, code, 60*time.Second)
			response.ToErrorResponse(errcode.Success)
		} else {
			err := errcode.InbalidParams
			err.WithDetails("没有注册此邮箱")
			response.ToErrorResponse(err)
		}
	}
}
