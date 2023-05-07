package middleware

import (
	"miniproject_golang/constants"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(userId int, username string) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userId
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.JWT_SECRET))
}
