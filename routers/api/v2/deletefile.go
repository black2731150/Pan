package v2

import (
	"pan/pkg/app"
	"pan/pkg/errcode"
	"pan/utils"

	"github.com/gin-gonic/gin"
)

//删除文件
func DeleteFile() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		response := app.NewRespponse(ctx)

		filepath := ctx.PostForm("filepath")

		err := utils.RemoveFile(filepath)
		if err != nil {
			response.ToErrorResponse(errcode.DeleteFileError)
			return
		}

		data := gin.H{
			"code":    0,
			"message": "Delete success!",
		}
		response.ToResponse(data)
	}
}
