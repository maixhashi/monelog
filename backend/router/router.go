package router

import (
	"monelog/controller"
	"github.com/labstack/echo/v4"
	"monelog/router/routes"
)

func NewRouter(
	uc controller.IUserController,
	tc controller.ITaskController,
	csc controller.ICardStatementController,
	dcsc controller.IDevCardStatementController,
	chc controller.ICSVHistoryController) *echo.Echo {
	
	e := echo.New()
	
	// ミドルウェアの設定
	setupMiddleware(e)
	
	// 各種ルートの設定
	routes.SetupAuthRoutes(e, uc)
	routes.SetupTaskRoutes(e, tc)
	routes.SetupCardStatementRoutes(e, csc) // カード明細ルートの設定
	routes.SetupDevRoutes(e, dcsc)
	routes.SetupCSVHistoryRoutes(e, chc) // CSV履歴ルートの設定
	
	return e
}