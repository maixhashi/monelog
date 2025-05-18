package card_statement

import (
	"monelog/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) SaveCardStatements(c echo.Context) error {
	userId, err := getUserIdFromToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "認証に失敗しました")
	}
	
	var request model.CardStatementSaveRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request format"})
	}
	
	request.UserId = userId
	
	// リクエストの検証を強化
	if request.Year <= 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Year must be positive"})
	}
	
	if request.Month < 1 || request.Month > 12 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Month must be between 1 and 12"})
	}
	
	if len(request.CardStatements) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Card statements cannot be empty"})
	}
	
	// カード明細の保存
	cardStatementsRes, err := h.csu.SaveCardStatements(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	
	return c.JSON(http.StatusCreated, cardStatementsRes)
}