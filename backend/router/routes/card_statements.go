package routes

import (
	"monelog/controller"
	"monelog/utils/middleware"
	"github.com/labstack/echo/v4"
)

// SetupCardStatementRoutes はカード明細関連のルートを設定します
func SetupCardStatementRoutes(e *echo.Echo, csc controller.ICardStatementController) {
	cs := e.Group("/card-statements")
	cs.Use(middleware.GetJWTMiddleware())
	
	// 一覧取得
	cs.GET("", csc.GetAllCardStatements)
	
	// 個別取得
	cs.GET("/:cardStatementId", csc.GetCardStatementById)
	
	// 支払月ごとの取得 - 新機能
	cs.GET("/by-month", csc.GetCardStatementsByMonth)
	
	// CSVアップロード（直接保存）- 既存機能
	cs.POST("/upload", csc.UploadCSV)
	
	// CSVプレビュー（保存なし）- 新機能
	cs.POST("/preview", csc.PreviewCSV)
	
	// プレビューデータ保存 - 新機能
	cs.POST("/save", csc.SaveCardStatements)
}