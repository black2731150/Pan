package v2

import (
	"pan/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

//创建新建文件夹
func CreateNewFolder() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userid := ctx.GetUint("UserID")
		floderpath := ctx.DefaultQuery("floderpath", "新建文件夹")

		path := "storage" + "/" + strconv.Itoa(int(userid)) + "/" + floderpath
		utils.MakeDir(path)
	}
}
