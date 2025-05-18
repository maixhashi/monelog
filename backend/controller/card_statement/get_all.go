package card_statement

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetAllCardStatements(c echo.Context) error {
	userId, err := getUserIdFromToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "認証に失敗しました")
	}
	
	cardStatementsRes, err := h.csu.GetAllCardStatements(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, cardStatementsRes)
}