package errcode

import (
	"fmt"
	"net/http"
)

var (
	Success                   = NewError(0, "成功")
	ServerError               = NewError(10000000, "服务器内部错误")
	InbalidParams             = NewError(10000001, "入参错误")
	NotFound                  = NewError(10000002, "找不到")
	UnauthorizedAuthNotExist  = NewError(10000003, "鉴权失败,找不到对应的AppKey和AppSecret")
	UnauthorizedTokenError    = NewError(10000004, "鉴权失败,Token错误")
	UnauthorizedTokenTimeOut  = NewError(10000005, "鉴权失败,Token超时")
	UnauthorizedTokenGenerate = NewError(10000006, "鉴权失败,Token生成失败")
	TooManyRequests           = NewError(10000007, "请求过多")
	HaveTheUser               = NewError(10000008, "用户名已经存在")
	TheEmailCodeError         = NewError(10000009, "验证码错误")
)

var (
	FaildUploadFile   = NewError(20000001, "上传文件失败")
	CanNotFindFile    = NewError(20000002, "找不到文件")
	DownloadFileError = NewError(20000003, "下载文件错误")
	RenameErrer       = NewError(20000004, "更改文件名错误")
	DeleteFileError   = NewError(20000005, "删除文件错误")
	CanNotFindFolder  = NewError(20000006, "找不到文件夹")
)

type Error struct {
	code    int      //`json:"code"`
	message string   //`json:"message"`
	details []string //`json:"details"`
}

var codes = map[int]string{}

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已经存在,请更换一个", code))
	}
	codes[code] = msg
	return &Error{
		code:    code,
		message: msg,
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("错误码: %d,错误信息 : %s", e.Code(), e.Msg())
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Msg() string {
	return e.message
}

func (e *Error) Msgf(args []interface{}) string {
	return fmt.Sprintf(e.message, args...)
}

func (e *Error) Details() []string {
	return e.details
}

func (e *Error) WithDetails(details ...string) *Error {
	newError := *e
	newError.details = []string{}
	newError.details = append(newError.details, details...)

	return &newError
}

func (e *Error) StatusCode() int {
	switch e.Code() {
	case Success.Code():
		return http.StatusOK
	case ServerError.Code():
		return http.StatusInternalServerError
	case InbalidParams.Code():
		return http.StatusBadRequest
	case UnauthorizedAuthNotExist.Code():
		fallthrough
	case UnauthorizedTokenError.Code():
		fallthrough
	case UnauthorizedTokenGenerate.Code():
		fallthrough
	case UnauthorizedTokenTimeOut.Code():
		fallthrough
	case TooManyRequests.Code():
		return http.StatusTooManyRequests
	}
	return http.StatusInternalServerError
}
