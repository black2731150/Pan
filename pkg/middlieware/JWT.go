package middlieware

import (
	"pan/pkg/app"
	"pan/pkg/errcode"
	tok "pan/pkg/token"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var (
			token string
			ecode = errcode.Success
		)

		if s, exist := ctx.GetQuery("token"); exist {
			token = s
		} else {
			token = ctx.GetHeader("token")
		}

		if token == "" {
			ecode = errcode.InbalidParams
		} else {
			_, err := tok.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					ecode = errcode.UnauthorizedTokenTimeOut
				default:
					ecode = errcode.UnauthorizedTokenError
				}
			}
		}

		if ecode != errcode.Success {
			respnse := app.NewRespponse(ctx)
			respnse.ToErrorResponse(ecode)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
