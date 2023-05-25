package middlieware

import (
	"pan/models"
	"pan/pkg/app"
	"pan/pkg/errcode"
	tok "pan/pkg/token"
	"pan/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var (
			token string
			ecode = errcode.Success
		)

		token, err := ctx.Cookie("token")
		if err != nil {
			ecode = errcode.UnauthorizedTokenError
		}

		if token == "" {
			ecode = errcode.UnauthorizedTokenError
		} else {
			claim, err := tok.ParseToken(token)
			// fmt.Println(claim)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					ecode = errcode.UnauthorizedTokenTimeOut
				default:
					ecode = errcode.UnauthorizedTokenError
				}
			}

			user := models.NewUser()
			user.ID = claim.UserID
			user.GetEmailAndUsernameFromID()
			ok := utils.StringMD5(user.UserName) == claim.Username && utils.StringMD5(user.Email) == claim.Email
			if !ok {
				ecode = errcode.UnauthorizedTokenError
			}
			ctx.Set("UserID", user.ID)
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
