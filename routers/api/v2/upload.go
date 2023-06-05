package v2

import (
	"fmt"
	"os"
	"pan/pkg/app"
	"pan/pkg/errcode"

	"github.com/gin-gonic/gin"
)

//上传文件
func Upload() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response := app.NewRespponse(ctx)
		file, err := ctx.FormFile("file")
		userid := ctx.GetUint("UserID")

		if err != nil {
			response.ToErrorResponse(errcode.FaildUploadFile.WithDetails("上传文件接受参数失败"))
			return
		}

		pwd, _ := os.Getwd()
		filepath := fmt.Sprintf("%s/storage/%d/%s", pwd, userid, file.Filename)
		ctx.SaveUploadedFile(file, filepath)
	}
}
