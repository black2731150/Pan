package v2

import (
	"pan/pkg/app"
	"pan/pkg/errcode"
	"pan/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

//重命名文件或者文件夹
func UpdateFileName() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response := app.NewRespponse(ctx)
		newname := ctx.PostForm("newname")
		userid := ctx.GetUint("UserID")
		filepath := ctx.PostForm("filepath")

		path := "storage" + "/" + strconv.Itoa(int(userid)) + "/" + filepath

		err := utils.RenameFile(path, newname)
		if err != nil {
			response.ToErrorResponse(errcode.RenameErrer)
		}

		data := gin.H{
			"code":    0,
			"message": "Rename success",
		}
		response.ToResponse(data)
	}
}
