package v2

import (
	"pan/utils"

	"github.com/gin-gonic/gin"
)

//创建新建文件夹
func CreateNewFolder() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		folderpath := ctx.DefaultQuery("folderpath", "storage/13/新建文件夹")
		utils.MakeDir(folderpath)
	}
}
