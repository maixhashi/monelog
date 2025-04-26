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
	
	// 一覧取得
	ch.GET("", chc.GetAllCSVHistories)
	
	// 月別取得（新機能）- 個別取得より先に定義する必要がある
	ch.GET("/by-month", chc.GetCSVHistoriesByMonth)
	
	// 個別取得
	ch.GET("/:csvHistoryId", chc.GetCSVHistoryById)
	
	// ファイルダウンロード
	ch.GET("/:csvHistoryId/download", chc.DownloadCSVHistory)
	
	// CSV履歴保存
	ch.POST("", chc.SaveCSVHistory)
	
	// CSV履歴削除
	ch.DELETE("/:csvHistoryId", chc.DeleteCSVHistory)
}
