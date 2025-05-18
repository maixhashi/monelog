package card_statement

import (
	"log"
	"monelog/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *Handler) UploadCSV(c echo.Context) error {
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
	
	// 年月の取得（CSVHistoryにも保存するため）
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
	
	request := model.CardStatementRequest{
		CardType: cardType,
		UserId:   userId,
		Year:     year,
		Month:    month,
	}
	
	// CSVの処理
	cardStatementsRes, err := h.csu.ProcessCSV(file, request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	
	// CSV履歴も保存する
	csvHistoryRequest := model.CSVHistorySaveRequest{
		FileName: file.Filename,
		CardType: cardType,
		Year:     year,
		Month:    month,
		UserId:   userId,
	}
	
	// 注入されたCSV履歴ユースケースを使用
	_, err = h.chu.SaveCSVHistory(file, csvHistoryRequest)
	if err != nil {
		// CSV履歴の保存に失敗しても、カード明細の処理は続行
		// エラーログを出力するなどの対応が望ましい
		log.Printf("Failed to save CSV history: %v", err)
	}
	
	return c.JSON(http.StatusCreated, cardStatementsRes)
}