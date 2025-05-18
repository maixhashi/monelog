package card_statement

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

// getUserIdFromToken はJWTトークンからユーザーIDを取得する関数
func getUserIdFromToken(c echo.Context) (uint, error) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := uint(claims["user_id"].(float64))
	return userId, nil
}