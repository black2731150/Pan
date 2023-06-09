package v2

import (
	"fmt"
	"pan/pkg/app"
	"pan/pkg/errcode"
	"pan/utils"

	"github.com/gin-gonic/gin"
)

//创建新建文件夹
func CreateNewFolder() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response := app.NewRespponse(ctx)

		userid := ctx.GetUint("UserID")

		floderpath := ctx.DefaultQuery("floderpath", "")
		newflodername := ctx.DefaultQuery("newflodername", "")
		if newflodername == "" {
			newflodername = "新建文件夹"
		}

		path := fmt.Sprintf("storage/%d/%s/%s", userid, floderpath, newflodername)

		err := utils.MakeDir(path)

		if err != nil {
			response.ToErrorResponse(errcode.CreateDirError.WithDetails("创建文件夹失败"))
			return
		}

		response.ToSuccessResponse("Create floder success!")
	}
}
