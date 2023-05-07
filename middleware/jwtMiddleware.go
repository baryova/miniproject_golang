package middleware

import (
	"miniproject_golang/constants"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func CreateToken(userId int, username string) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userId
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.JWT_SECRET))
}

func GetUserIdByToken(c echo.Context) (int, error) {
	user := c.Get("user").(*jwt.Token)
	if !user.Valid {
		return 0, echo.NewHTTPError(401, "Unauthorized")
	}
	claims := user.Claims.(jwt.MapClaims)

	userId := int(claims["user_id"].(float64))

	return userId, nil
}
