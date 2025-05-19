package csv_history

import (
	"monelog/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *Handler) SaveCSVHistory(c echo.Context) error {
	userId, err := getUserIdFromToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "認証に失敗しました")
	}
	
	// ファイルの取得
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ファイルが見つかりません"})
	}
	
	// ファイル名の取得
	fileName := c.FormValue("file_name")
	if fileName == "" {
		// ファイル名が指定されていない場合は、アップロードされたファイル名を使用
		fileName = file.Filename
	}
	
	// カード種類の取得
	cardType := c.FormValue("card_type")
	if cardType == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "card_type is required"})
	}
	
	// 年月の取得
	yearStr := c.FormValue("year")
	monthStr := c.FormValue("month")
	
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid year format"})
	}
	
	month, err := strconv.Atoi(monthStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid month format"})
	}
	
	request := model.CSVHistorySaveRequest{
		FileName: fileName,
		CardType: cardType,
		Year:     year,
		Month:    month,
		UserId:   userId,
	}
	
	// CSV履歴の保存
	csvHistoryRes, err := h.chu.SaveCSVHistory(file, request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	
	return c.JSON(http.StatusCreated, csvHistoryRes)
}