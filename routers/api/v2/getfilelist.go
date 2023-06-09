package v2

import (
	"fmt"
	"pan/pkg/app"
	"pan/pkg/errcode"
	"pan/utils"

	"github.com/gin-gonic/gin"
)

//获取一个文件夹里面的文件列表信息
func GetFileList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response := app.NewRespponse(ctx)

		userid := ctx.GetUint("UserID")

		folderpath := ctx.Query("floderpath")

		path := fmt.Sprintf("storage/%d/%s", userid, folderpath)

		fileinfos, err := utils.GetFilesInfoFromFolder(path)
		if err != nil {
			response.ToErrorResponse(errcode.CanNotFindFolder.WithDetails("文件夹遍历失败"))
			return
		}

		data := gin.H{
			"code":      0,
			"message":   "Get file informations success",
			"fileinfos": fileinfos,
		}
		response.ToResponse(data)
	}
}
