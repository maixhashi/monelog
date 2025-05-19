package csv_history

import (
	"monelog/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GetCSVHistoriesByMonth は指定された年月のCSV履歴を取得します
func (h *Handler) GetCSVHistoriesByMonth(c echo.Context) error {
	// ユーザーIDを取得
	userId, err := getUserIdFromToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "認証に失敗しました")
	}

	// クエリパラメータから年と月を取得
	yearStr := c.QueryParam("year")
	monthStr := c.QueryParam("month")

	// 年と月のパラメータが提供されているか確認
	if yearStr == "" || monthStr == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "年と月のパラメータが必要です",
		})
	}

	// 年を整数に変換
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "無効な年パラメータです",
		})
	}

	// 月を整数に変換
	month, err := strconv.Atoi(monthStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "無効な月パラメータです",
		})
	}

	// 月の範囲を検証（1〜12）
	if month < 1 || month > 12 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "月は1から12の間である必要があります",
		})
	}

	// ユースケースを呼び出して月別のCSV履歴を取得
	csvHistories, err := h.chu.GetCSVHistoriesByMonth(userId, year, month)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "CSV履歴の取得に失敗しました",
		})
	}

	// CSV履歴のレスポンスを作成
	var responses []model.CSVHistoryResponse
	for _, csvHistory := range csvHistories {
		responses = append(responses, model.CSVHistoryResponse{
			ID:        csvHistory.ID,
			FileName:  csvHistory.FileName,
			CardType:  csvHistory.CardType,
			Year:      csvHistory.Year,
			Month:     csvHistory.Month,
			CreatedAt: csvHistory.CreatedAt,
			UpdatedAt: csvHistory.UpdatedAt,
		})
	}

	return c.JSON(http.StatusOK, responses)
}