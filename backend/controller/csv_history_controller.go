package controller

import (
	"monelog/controller/csv_history"
	"monelog/usecase"

	"github.com/labstack/echo/v4"
)

type ICSVHistoryController interface {
	GetAllCSVHistories(c echo.Context) error
	GetCSVHistoryById(c echo.Context) error
	GetCSVHistoriesByMonth(c echo.Context) error
	SaveCSVHistory(c echo.Context) error
	DeleteCSVHistory(c echo.Context) error
	DownloadCSVHistory(c echo.Context) error
}

type csvHistoryController struct {
	handler *csv_history.Handler
}

func NewCSVHistoryController(chu usecase.ICSVHistoryUsecase) ICSVHistoryController {
	return &csvHistoryController{
		handler: csv_history.NewHandler(chu),
	}
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
	return chc.handler.GetAllCSVHistories(c)
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
	return chc.handler.GetCSVHistoryById(c)
}

// GetCSVHistoriesByMonth 月別のCSV履歴を取得
// @Summary 月別のCSV履歴一覧を取得
// @Description 指定された年月のCSV履歴を取得する
// @Tags csv-histories
// @Accept json
// @Produce json
// @Param year query int true "年 (例: 2023)"
// @Param month query int true "月 (1-12)"
// @Success 200 {array} model.CSVHistoryResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /csv-histories/by-month [get]
func (chc *csvHistoryController) GetCSVHistoriesByMonth(c echo.Context) error {
	return chc.handler.GetCSVHistoriesByMonth(c)
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
// @Param year formData int true "年 (例: 2023)"
// @Param month formData int true "月 (1-12)"
// @Success 201 {object} model.CSVHistoryResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /csv-histories [post]
func (chc *csvHistoryController) SaveCSVHistory(c echo.Context) error {
	return chc.handler.SaveCSVHistory(c)
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
	return chc.handler.DeleteCSVHistory(c)
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
	return chc.handler.DownloadCSVHistory(c)
}
