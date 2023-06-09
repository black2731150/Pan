package app

import (
	"net/http"
	"pan/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Ctx *gin.Context
}

type Pager struct {
	Page      int `json:"page"`
	PageSize  int `json:"page_size"`
	TotalRows int `json:"total_rows"`
}

func NewRespponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

func (r *Response) ToResponse(data interface{}) {
	if data == nil {
		r.Ctx.JSON(http.StatusOK, gin.H{})
	}
	r.Ctx.JSON(http.StatusOK, data)
}

func (r *Response) ToResponseList(list interface{}, totalRows int) {
	r.Ctx.JSON(http.StatusOK, gin.H{
		"list": list,
		"pager": Pager{
			Page:      GetPage(r.Ctx),
			PageSize:  GtPageSize(r.Ctx),
			TotalRows: totalRows,
		},
	})
}

func (r *Response) ToSuccessResponse(message string) {
	r.Ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": message,
	})
}

func (r *Response) ToErrorResponse(err *errcode.Error) {
	response := gin.H{
		"code":    err.Code(),
		"message": err.Msg(),
	}
	details := err.Details()
	if len(details) > 0 {
		response["details"] = details
	}

	r.Ctx.JSON(err.StatusCode(), response)
}
