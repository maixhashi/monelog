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
	cs.GET("", csc.GetAllCardStatements)
	cs.GET("/:cardStatementId", csc.GetCardStatementById)
	cs.POST("/upload", csc.UploadCSV)
}
