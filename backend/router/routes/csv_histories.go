package routes

import (
	"monelog/controller"
	"monelog/utils/middleware"
	"github.com/labstack/echo/v4"
)

// SetupCSVHistoryRoutes はCSV履歴関連のルートを設定します
func SetupCSVHistoryRoutes(e *echo.Echo, chc controller.ICSVHistoryController) {
	ch := e.Group("/csv-histories")
	ch.Use(middleware.GetJWTMiddleware())
	ch.GET("", chc.GetAllCSVHistories)
	ch.GET("/:csvHistoryId", chc.GetCSVHistoryById)
	ch.GET("/:csvHistoryId/download", chc.DownloadCSVHistory)
	ch.POST("", chc.SaveCSVHistory)
	ch.DELETE("/:csvHistoryId", chc.DeleteCSVHistory)
}
