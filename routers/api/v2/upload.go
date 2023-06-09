package v2

import (
	"fmt"
	"pan/pkg/app"
	"pan/pkg/errcode"

	"github.com/gin-gonic/gin"
)

//上传文件
func Upload() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response := app.NewRespponse(ctx)

		userid := ctx.GetUint("UserID")

		file, err := ctx.FormFile("file")
		floderpath := ctx.PostForm("floderpath")

		if err != nil {
			response.ToErrorResponse(errcode.FaildUploadFile.WithDetails("上传文件接受参数失败"))
			return
		}

		floderpath = fmt.Sprintf("storage/%d/%s/%s", userid, floderpath, file.Filename)

		err = ctx.SaveUploadedFile(file, floderpath)
		if err != nil {
			response.ToErrorResponse(errcode.FaildUploadFile.WithDetails("保存上传的文件失败"))
			return
		}

		response.ToSuccessResponse("Upload File Success!")
	}
}
