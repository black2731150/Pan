package v2

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"pan/pkg/app"
	"pan/pkg/errcode"

	"github.com/gin-gonic/gin"
)

//下载文件
// func Download() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		response := app.NewRespponse(ctx)

// 		userid := ctx.GetUint("UserID")

// 		floderpath := ctx.DefaultQuery("floderpath", "")
// 		filename := ctx.DefaultQuery("filename", "")
// 		if filename == "" {
// 			response.ToErrorResponse(errcode.CanNotFindFile.WithDetails("没有指定文件"))
// 			return
// 		}

// 		path := fmt.Sprintf("storage/%d/%s/%s", userid, floderpath, filename)

// 		file, err := utils.OpenFile(path)
// 		if err != nil {
// 			response.ToErrorResponse(errcode.CanNotFindFile.WithDetails("打开需要下载的文件失败"))
// 			return
// 		}
// 		// defer file.Close()

// 		fileInfo, err := file.Stat()
// 		if err != nil {
// 			response.ToErrorResponse(errcode.DownloadFileError.WithDetails("获取文件信息失败"))
// 			return
// 		}

// 		ctx.Header("Content-Type", "application/octet-stream")
// 		ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileInfo.Name()))
// 		ctx.Header("Content-Length", strconv.FormatInt(fileInfo.Size(), 10))

// 		file.Close()

// 		// chuck := 1024 * 4

// 		// for {
// 		// 	buf := make([]byte, chuck)
// 		// 	n, err := file.Read(buf)
// 		// 	if err == io.EOF {
// 		// 		break
// 		// 	}
// 		// 	if err != nil {
// 		// 		response.ToErrorResponse(errcode.DownloadFileError.WithDetails("下载文件过程出错"))
// 		// 		return
// 		// 	}
// 		// 	ctx.Writer.Write(buf[:n])
// 		// }

// 		ctx.File(path)
// 	}
// }

func Download() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response := app.NewRespponse(ctx)

		userid := ctx.GetUint("UserID")

		floderpath := ctx.DefaultQuery("floderpath", "")
		filename := ctx.DefaultQuery("filename", "")
		if filename == "" {
			response.ToErrorResponse(errcode.CanNotFindFile.WithDetails("没有指定文件"))
			return
		}

		path := fmt.Sprintf("storage/%d/%s/%s", userid, floderpath, filename)

		file, err := os.Open(path)
		if err != nil {
			response.ToErrorResponse(errcode.CanNotFindFile.WithDetails("打开需要下载的文件失败"))
			return
		}
		defer file.Close()

		fileContent, err := ioutil.ReadAll(file)
		if err != nil {
			response.ToErrorResponse(errcode.DownloadFileError.WithDetails("读取文件内容失败"))
			return
		}

		encodedContent := base64.StdEncoding.EncodeToString(fileContent)

		blob := map[string]string{
			"filename": filename,
			"content":  encodedContent,
		}

		ctx.JSON(http.StatusOK, blob)
	}
}
