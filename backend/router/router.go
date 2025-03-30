package router

import (
	"monelog/controller"
	"github.com/labstack/echo/v4"
	"monelog/router/routes"
)

func NewRouter(
	uc controller.IUserController,
	tc controller.ITaskController) *echo.Echo {
	
	e := echo.New()
	
	// ミドルウェアの設定
	setupMiddleware(e)
	
	// 各種ルートの設定
	routes.SetupAuthRoutes(e, uc)
	routes.SetupTaskRoutes(e, tc)
	
	return e
}