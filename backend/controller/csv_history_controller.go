package controller

import (
	"fmt"
	"monelog/model"
	"monelog/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ICSVHistoryController interface {
	GetAllCSVHistories(c echo.Context) error
	GetCSVHistoryById(c echo.Context) error
	SaveCSVHistory(c echo.Context) error
	DeleteCSVHistory(c echo.Context) error
	DownloadCSVHistory(c echo.Context) error
}

type csvHistoryController struct {
	chu usecase.ICSVHistoryUsecase
}

func NewCSVHistoryController(chu usecase.ICSVHistoryUsecase) ICSVHistoryController {
	return &csvHistoryController{chu}
}

// GetAllCSVHistories ユーザーのすべてのCSV履歴を取得
// @Summary ユーザーのCSV履歴一覧を取得
// @Description ログインユーザーのすべてのCSV履歴を取得する
// @Tags csv-histories
// @Accept json
// @Produce json
// @Success 200 {array} model.CSVHistoryResponse
// @Failure 500 {object} map[string]string
// @Router /csv-histories [get]
func (chc *csvHistoryController) GetAllCSVHistories(c echo.Context) error {
	userId, err := getUserIdFromToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "認証に失敗しました")
	}
	
	csvHistoriesRes, err := chc.chu.GetAllCSVHistories(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, csvHistoriesRes)
}

// GetCSVHistoryById 指定されたIDのCSV履歴を取得
// @Summary 特定のCSV履歴を取得
// @Description 指定されたIDのCSV履歴を取得する
// @Tags csv-histories
// @Accept json
// @Produce json
// @Param csvHistoryId path int true "CSV履歴ID"
// @Success 200 {object} model.CSVHistoryDetailResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /csv-histories/{csvHistoryId} [get]
func (chc *csvHistoryController) GetCSVHistoryById(c echo.Context) error {
	userId, err := getUserIdFromToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "認証に失敗しました")
	}
	
	id := c.Param("csvHistoryId")
	csvHistoryId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid CSV history ID"})
	}
	
	csvHistoryRes, err := chc.chu.GetCSVHistoryById(userId, uint(csvHistoryId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, csvHistoryRes)
}

// SaveCSVHistory CSVファイルを履歴として保存
// @Summary CSVファイルを履歴として保存
// @Description カード明細のCSVファイルを履歴として保存する
// @Tags csv-histories
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "CSVファイル"
// @Param file_name formData string true "ファイル名"
// @Param card_type formData string true "カード種類 (rakuten, mufg, epos)"
// @Success 201 {object} model.CSVHistoryResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /csv-histories [post]
func (chc *csvHistoryController) SaveCSVHistory(c echo.Context) error {
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
	
	request := model.CSVHistorySaveRequest{
		FileName: fileName,
		CardType: cardType,
		UserId:   userId,
	}
	
	// CSV履歴の保存
	csvHistoryRes, err := chc.chu.SaveCSVHistory(file, request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	
	return c.JSON(http.StatusCreated, csvHistoryRes)
}

// DeleteCSVHistory 指定されたIDのCSV履歴を削除
// @Summary CSV履歴を削除
// @Description 指定されたIDのCSV履歴を削除する
// @Tags csv-histories
// @Accept json
// @Produce json
// @Param csvHistoryId path int true "CSV履歴ID"
// @Success 204 {string} string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /csv-histories/{csvHistoryId} [delete]
func (chc *csvHistoryController) DeleteCSVHistory(c echo.Context) error {
	userId, err := getUserIdFromToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "認証に失敗しました")
	}
	
	id := c.Param("csvHistoryId")
	csvHistoryId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid CSV history ID"})
	}
	
	if err := chc.chu.DeleteCSVHistory(userId, uint(csvHistoryId)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	
	return c.NoContent(http.StatusNoContent)
}

// DownloadCSVHistory 指定されたIDのCSV履歴からCSVファイルをダウンロード
// @Summary CSV履歴からCSVファイルをダウンロード
// @Description 指定されたIDのCSV履歴からCSVファイルをダウンロードする
// @Tags csv-histories
// @Accept json
// @Produce text/csv
// @Param csvHistoryId path int true "CSV履歴ID"
// @Success 200 {file} file
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /csv-histories/{csvHistoryId}/download [get]
func (chc *csvHistoryController) DownloadCSVHistory(c echo.Context) error {
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
    csvHistoryDetail, err := chc.chu.GetCSVHistoryById(userId, uint(csvHistoryId))
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    
    // Content-Dispositionヘッダーを設定してファイルとしてダウンロードさせる
    c.Response().Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", csvHistoryDetail.FileName))
    c.Response().Header().Set("Content-Type", "text/csv")
    
    // ファイルデータを返す
    return c.Blob(http.StatusOK, "text/csv", csvHistoryDetail.FileData)
}
