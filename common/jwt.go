package common

import "github.com/golang-jwt/jwt/v4"

func GetJwt(userId string, secretKey string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
