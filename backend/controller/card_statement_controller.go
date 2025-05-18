package controller

import (
	"monelog/controller/card_statement"
	"monelog/usecase"

	"github.com/labstack/echo/v4"
)

type ICardStatementController interface {
	GetAllCardStatements(c echo.Context) error
	GetCardStatementById(c echo.Context) error
	GetCardStatementsByMonth(c echo.Context) error
	UploadCSV(c echo.Context) error
	PreviewCSV(c echo.Context) error
	SaveCardStatements(c echo.Context) error
}

type cardStatementController struct {
	handler *card_statement.Handler
}

func NewCardStatementController(csu usecase.ICardStatementUsecase, chu usecase.ICSVHistoryUsecase) ICardStatementController {
	return &cardStatementController{
		handler: card_statement.NewHandler(csu, chu),
	}
}

// GetAllCardStatements ユーザーのすべてのカード明細を取得
// @Summary ユーザーのカード明細一覧を取得
// @Description ログインユーザーのすべてのカード明細を取得する
// @Tags card-statements
// @Accept json
// @Produce json
// @Success 200 {array} model.CardStatementResponse
// @Failure 500 {object} map[string]string
// @Router /card-statements [get]
func (csc *cardStatementController) GetAllCardStatements(c echo.Context) error {
	return csc.handler.GetAllCardStatements(c)
}

// GetCardStatementById 指定されたIDのカード明細を取得
// @Summary 特定のカード明細を取得
// @Description 指定されたIDのカード明細を取得する
// @Tags card-statements
// @Accept json
// @Produce json
// @Param cardStatementId path int true "カード明細ID"
// @Success 200 {object} model.CardStatementResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /card-statements/{cardStatementId} [get]
func (csc *cardStatementController) GetCardStatementById(c echo.Context) error {
	return csc.handler.GetCardStatementById(c)
}

// UploadCSV CSVファイルをアップロードして解析
// @Summary CSVファイルをアップロードして解析
// @Description カード明細のCSVファイルをアップロードして解析する
// @Tags card-statements
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "CSVファイル"
// @Param card_type formData string true "カード種類 (rakuten, mufg, epos)"
// @Param year formData int true "年 (例: 2023)"
// @Param month formData int true "月 (1-12)"
// @Success 201 {array} model.CardStatementResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /card-statements/upload [post]
func (csc *cardStatementController) UploadCSV(c echo.Context) error {
	return csc.handler.UploadCSV(c)
}

// PreviewCSV CSVファイルをアップロードしてプレビュー（保存なし）
// @Summary CSVファイルをアップロードしてプレビュー
// @Description カード明細のCSVファイルをアップロードして解析するが、DBには保存しない
// @Tags card-statements
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "CSVファイル"
// @Param card_type formData string true "カード種類 (rakuten, mufg, epos)"
// @Param year formData int false "年 (例: 2023)"
// @Param month formData int false "月 (1-12)"
// @Success 200 {array} model.CardStatementResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /card-statements/preview [post]
func (csc *cardStatementController) PreviewCSV(c echo.Context) error {
	return csc.handler.PreviewCSV(c)
}

// SaveCardStatements プレビューしたカード明細をDBに保存
// @Summary プレビューしたカード明細を保存
// @Description プレビューしたカード明細データをデータベースに保存する
// @Tags card-statements
// @Accept json
// @Produce json
// @Param request body model.CardStatementSaveRequest true "保存するカード明細データ"
// @Success 201 {array} model.CardStatementResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /card-statements/save [post]
func (csc *cardStatementController) SaveCardStatements(c echo.Context) error {
	return csc.handler.SaveCardStatements(c)
}

// GetCardStatementsByMonth 支払月ごとのカード明細を取得
// @Summary 支払月ごとのカード明細を取得
// @Description 指定された年月の支払いに関するカード明細を取得する
// @Tags card-statements
// @Accept json
// @Produce json
// @Param year query int true "年 (例: 2023)"
// @Param month query int true "月 (1-12)"
// @Success 200 {array} model.CardStatementResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /card-statements/by-month [get]
func (csc *cardStatementController) GetCardStatementsByMonth(c echo.Context) error {
	return csc.handler.GetCardStatementsByMonth(c)
}