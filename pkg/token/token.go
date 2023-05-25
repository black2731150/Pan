package token

import (
	"pan/global"
	"pan/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string
	Email    string
	UserID   uint
	jwt.StandardClaims
}

func GetJWTSecret() []byte {
	return []byte(global.Panserver.Config.JWTconfig.Secret)
}

func GenerateToken(username, email string, userID uint) (string, error) {
	nowTime := time.Now()
	// fmt.Println("nowTime:", nowTime)
	expireTime := nowTime.Add(7 * 24 * time.Hour)
	// fmt.Println("expireTime", expireTime)
	claims := Claims{
		Username: utils.StringMD5(username),
		Email:    utils.StringMD5(email),
		UserID:   userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    global.Panserver.Config.JWTconfig.Issure,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(GetJWTSecret())
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})
	// fmt.Println("tokenClaims:", tokenClaims)

	if tokenClaims != nil {
		claims, ok := tokenClaims.Claims.(*Claims)
		// fmt.Println("claims:", claims)
		if ok && tokenClaims.Valid {
			// fmt.Println("return success!")
			return claims, nil
		}
	}
	// fmt.Println("err:", err)
	return nil, err
}
