package routes

import (
	"monelog/controller"
	"os"

	"github.com/labstack/echo/v4"
)

// SetupDevRoutes は開発環境限定のルートを設定します
func SetupDevRoutes(e *echo.Echo, dcsc controller.IDevCardStatementController) {
	// 開発環境かどうかチェック
	env := os.Getenv("APP_ENV")
	if env != "development" && env != "dev" && env != "" {
		return // 開発環境でなければルートを設定しない
	}

	// 開発用ルートグループ
	dev := e.Group("/dev")
	
	// カード明細関連
	devCardStatements := dev.Group("/card-statements")
	devCardStatements.POST("/delete-all", dcsc.DeleteAllCardStatements)
}
