package csv_history

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *Handler) DownloadCSVHistory(c echo.Context) error {
    userId, err := getUserIdFromToken(c)
    if err != nil {
        return echo.NewHTTPError(http.StatusUnauthorized, "認証に失敗しました")
    }
    
    id := c.Param("csvHistoryId")
    csvHistoryId, err := strconv.ParseUint(id, 10, 32)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid CSV history ID"})
    }
    
    // CSV履歴の詳細を取得
    csvHistoryDetail, err := h.chu.GetCSVHistoryById(userId, uint(csvHistoryId))
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    
    // Content-Dispositionヘッダーを設定してファイルとしてダウンロードさせる
    c.Response().Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", csvHistoryDetail.FileName))
    c.Response().Header().Set("Content-Type", "text/csv")
    
    // ファイルデータを返す
    return c.Blob(http.StatusOK, "text/csv", csvHistoryDetail.FileData)
}