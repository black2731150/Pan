package v2

import (
	"fmt"
	"pan/pkg/app"
	"pan/pkg/errcode"
	"pan/utils"

	"github.com/gin-gonic/gin"
)

//删除文件
func DeleteFile() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		response := app.NewRespponse(ctx)

		userid := ctx.GetUint("UserID")

		floderpath := ctx.DefaultPostForm("floderpath", "")
		filename := ctx.DefaultPostForm("filename", "")
		if filename == "" {
			response.ToErrorResponse(errcode.CanNotFindFile.WithDetails("没有指定文件"))
			return
		}

		path := fmt.Sprintf("storage/%d/%s/%s", userid, floderpath, filename)
		err := utils.RemoveFile(path)
		if err != nil {
			response.ToErrorResponse(errcode.DeleteFileError)
			return
		}

		response.ToSuccessResponse("Delete success!")
	}
}
