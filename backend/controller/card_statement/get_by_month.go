package card_statement

import (
	"monelog/dto"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetCardStatementsByMonth(c echo.Context) error {
	userId, err := getUserIdFromToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "認証に失敗しました")
	}
	
	// クエリパラメータの取得
	yearStr := c.QueryParam("year")
	monthStr := c.QueryParam("month")
	
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid year format"})
	}
	
	month, err := strconv.Atoi(monthStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid month format"})
	}
	
	request := dto.CardStatementByMonthRequest{
		Year:   year,
		Month:  month,
		UserId: userId,
	}
	
	cardStatementsRes, err := h.csu.GetCardStatementsByMonth(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	
	return c.JSON(http.StatusOK, cardStatementsRes)
}