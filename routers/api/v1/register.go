package v1

import (
	"pan/global"
	"pan/models"
	"pan/pkg/app"
	"pan/pkg/errcode"
	"pan/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

//注册api
func Register() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response := app.NewRespponse(ctx)
		user := models.NewUser()
		//获取用户名并验证
		user.UserName = ctx.PostForm("username")

		if len(user.UserName) == 0 {
			err := errcode.InbalidParams
			err = err.WithDetails("用户名不能为空")
			//fmt.Println(err)
			response.ToErrorResponse(err)
			return
		}

		//获取密码并验证
		user.Password = ctx.PostForm("password")
		if len(user.Password) < 6 || len(user.Password) > 20 {
			err := errcode.InbalidParams
			err = err.WithDetails("密码长度不合规范")
			//fmt.Println(err)
			response.ToErrorResponse(err)
			return
		}
		user.Password = utils.StringMD5(user.Password)

		//验证这个用户名是否存在
		if user.HaveTheUserName() {
			err := errcode.InbalidParams
			err = err.WithDetails("用户名已经被注册")
			//fmt.Println(err)
			response.ToErrorResponse(err)
			return
		}

		//验证邮箱是否为空
		user.Email = ctx.PostForm("email")
		if len(user.Email) == 0 {
			err := errcode.InbalidParams
			err = err.WithDetails("邮箱账户不能为空")
			//fmt.Println(err)
			response.ToErrorResponse(err)
			return
		}

		//验证验证码是是否正确
		code := ctx.PostForm("code")
		realcode, ok := global.GetEmailCodeFromMap(user.Email)
		if ok || code == realcode {
		} else {
			err := errcode.InbalidParams
			err = err.WithDetails("验证码错误或过期")
			//fmt.Println(err)
			response.ToErrorResponse(err)
			return
		}

		//注册用户
		user.RegisterNewUser()
		user.GetUserIDFromUsername()
		utils.MakeDir("./storage/" + strconv.Itoa(int(user.ID)))

		//返回成功消息响应消息
		data := gin.H{
			"code":    0,
			"message": "Register Success!",
		}
		response.ToResponse(data)
	}
}
