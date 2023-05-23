package v1

import (
	"pan/dao"
	"pan/pkg/errcode"
	"pan/pkg/response"

	"github.com/gin-gonic/gin"
)

func TestUserName() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := dao.NewUser()
		user.UserName = ctx.PostForm("username")
		response := response.NewRespponse(ctx)
		if user.HaveTheUserName() {
			response.ToErrorResponse(errcode.HaveTheUser)
		} else {
			response.ToResponse(errcode.Success)
		}
	}
}
