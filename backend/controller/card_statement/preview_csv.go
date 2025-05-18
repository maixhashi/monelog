package card_statement

import (
	"monelog/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *Handler) PreviewCSV(c echo.Context) error {
	userId, err := getUserIdFromToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "認証に失敗しました")
	}
	
	// ファイルの取得
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ファイルが見つかりません"})
	}
	
	// カード種類の取得
	cardType := c.FormValue("card_type")
	if cardType == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "card_type is required"})
	}
	
	// 年月の取得（プレビュー時は任意）
	yearStr := c.FormValue("year")
	monthStr := c.FormValue("month")
	
	// 年月の変換
	var year, month int
	if yearStr != "" {
		year, err = strconv.Atoi(yearStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid year format"})
		}
	}
	
	if monthStr != "" {
		month, err = strconv.Atoi(monthStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid month format"})
		}
	}
	
	request := model.CardStatementPreviewRequest{
		CardType: cardType,
		UserId:   userId,
		Year:     year,
		Month:    month,
	}
	
	// CSVのプレビュー処理
	cardStatementsRes, err := h.csu.PreviewCSV(file, request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	
	return c.JSON(http.StatusOK, cardStatementsRes)
}