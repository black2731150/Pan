package v1

import (
	"pan/dao"
	"pan/pkg/errcode"
	"pan/pkg/response"

	"github.com/gin-gonic/gin"
)

//这个接口用于前端检测这个用户名是否存在
func TestUserName() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := dao.NewUser()
		user.UserName = ctx.PostForm("username")
		// fmt.Println("The user name is : ", user.UserName)
		response := response.NewRespponse(ctx)
		if user.HaveTheUserName() {
			response.ToErrorResponse(errcode.HaveTheUser)
		} else {
			response.ToErrorResponse(errcode.Success)
		}
	}
}
