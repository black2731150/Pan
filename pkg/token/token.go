package token

import (
	"pan/common"
	"pan/global"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	AppKey    string
	AppSecret string
	jwt.StandardClaims
}

func GetJWTSecret() []byte {
	return []byte(global.Panserver.Config.JWTconfig.Secret)
}

func GenerateToken(appKey, appSecret string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(global.Panserver.Config.JWTconfig.Expire))
	claims := Claims{
		AppKey:    common.StringMD5(appKey),
		AppSecret: common.StringMD5(appSecret),
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

	if tokenClaims != nil {
		claims, ok := tokenClaims.Claims.(*Claims)
		if ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
