package v2

import (
	"fmt"
	"io"
	"pan/pkg/app"
	"pan/pkg/errcode"
	"pan/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

//下载文件
func Download() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response := app.NewRespponse(ctx)
		userid := ctx.GetUint("UserID")
		fmt.Println(userid)
		filepath := ctx.DefaultQuery("filepath", "")
		if filepath == "" {
			response.ToErrorResponse(errcode.CanNotFindFile.WithDetails("没有指定文件"))
			return
		}

		path := "./storage" + "/" + strconv.Itoa(int(userid)) + "/" + filepath
		fmt.Println(path)

		file, err := utils.OpenFile(path)
		if err != nil {
			response.ToErrorResponse(errcode.CanNotFindFile.WithDetails("打开需要下载的文件失败"))
			return
		}
		defer file.Close()

		chuck := 1024 * 4

		ctx.Header("Content-Type", "application/octet-stream")

		for {
			buf := make([]byte, chuck)
			n, err := file.Read(buf)
			if err == io.EOF {
				break
			}
			if err != nil {
				response.ToErrorResponse(errcode.DownloadFileError.WithDetails("下载文件过程出错"))
				return
			}
			ctx.Writer.Write(buf[:n])
		}
	}
}
