package v2

import (
	"pan/pkg/app"
	"pan/pkg/errcode"
	"pan/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

//删除文件
func DeleteFile() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		response := app.NewRespponse(ctx)
		userid := ctx.GetUint("UserID")
		filepath := ctx.PostForm("filepath")

		path := "storage" + "/" + strconv.Itoa(int(userid)) + "/" + filepath
		err := utils.RemoveFile(path)
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
