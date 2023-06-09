package v2

import (
	"fmt"
	"pan/pkg/app"
	"pan/pkg/errcode"
	"pan/utils"

	"github.com/gin-gonic/gin"
)

//重命名文件或者文件夹
func UpdateFileName() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response := app.NewRespponse(ctx)

		userid := ctx.GetUint("UserID")

		oldname := ctx.PostForm("oldname")
		newname := ctx.PostForm("newname")
		floderpath := ctx.PostForm("floderpath")

		oldpath := fmt.Sprintf("storage/%d/%s/%s", userid, floderpath, oldname)
		newpath := fmt.Sprintf("storage/%d/%s/%s", userid, floderpath, newname)
		fmt.Println(oldpath, newpath)

		err := utils.RenameFile(oldpath, newpath)
		if err != nil {
			response.ToErrorResponse(errcode.RenameErrer)
			return
		}

		response.ToSuccessResponse("Rename success")
	}
}
