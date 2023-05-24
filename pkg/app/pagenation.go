package app

import (
	"pan/global"
	"pan/pkg/convert"

	"github.com/gin-gonic/gin"
)

func GetPage(ctx *gin.Context) int {
	page := convert.StrTo(ctx.Query("page")).MustInt()
	if page <= 0 {
		return 1
	}

	return page
}

func GtPageSize(ctx *gin.Context) int {
	pageSize := convert.StrTo(ctx.Query("page_size")).MustInt()
	if pageSize <= 0 {
		return global.Panserver.Config.Pan.Defaultpagepize
	}
	if pageSize > global.Panserver.Config.Pan.Maxpagesize {
		return global.Panserver.Config.Pan.Maxpagesize
	}

	return pageSize
}

func GetPageOffset(page, pagesize int) int {
	result := 0
	if page > 0 {
		result = (page - 1) * pagesize
	}
	return result
}
