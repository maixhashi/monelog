package csv_history

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetAllCSVHistories(c echo.Context) error {
	userId, err := getUserIdFromToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "認証に失敗しました")
	}
	
	csvHistoriesRes, err := h.chu.GetAllCSVHistories(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, csvHistoriesRes)
}