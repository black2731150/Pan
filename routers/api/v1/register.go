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

//注册api
func Register() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response := response.NewRespponse(ctx)
		user := dao.NewUser()
		user.UserName = ctx.PostForm("username")

		if len(user.UserName) == 0 {
			err := errcode.InbalidParams
			err.WithDetails("用户名不能为空")
			response.ToErrorResponse(err)
			return
		}

		user.Password = ctx.PostForm("password")
		if len(user.Password) < 6 || len(user.Password) > 20 {
			err := errcode.InbalidParams
			err.WithDetails("密码长度不合规范")
			response.ToErrorResponse(err)
			return
		}
		user.Password = common.StringMD5(user.Password)

		if user.HaveTheUserName() {
			err := errcode.InbalidParams
			err.WithDetails("用户名已经被注册")
			response.ToErrorResponse(err)
			return
		}

		user.Email = ctx.PostForm("email")
		if len(user.Email) == 0 {
			err := errcode.InbalidParams
			err.WithDetails("邮箱账户不能为空")
			response.ToErrorResponse(err)
			return
		}

		code := common.GetRandSixCode()
		email.SendEmail(user.Email, "验证码", code)
		global.AddToMap(user.Email, code, 60*time.Second)

		user.Phonenum = ctx.PostForm("phonenum")

		user.RegisterNewUser()
		data := gin.H{
			"code":    0,
			"message": "Register Success!",
		}
		response.ToResponse(data)
	}
}
