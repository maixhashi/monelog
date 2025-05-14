package routes

import (
	"monelog/controller"
	"monelog/utils/middleware"
	"github.com/labstack/echo/v4"
)

// SetupAuthRoutes は認証関連のルートを設定します
func SetupAuthRoutes(e *echo.Echo, uc controller.IUserController) {
	// 認証不要のルート
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.GET("/csrf-token", uc.CsrfToken)
	e.GET("/auth-verify", uc.VerifyAuth) // JWT ミドルウェアを適用しない
	
	// 認証が必要なルート
	auth := e.Group("")
	auth.Use(middleware.GetJWTMiddleware())
	auth.POST("/logout", uc.LogOut)
}
