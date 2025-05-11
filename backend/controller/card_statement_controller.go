package controller

import (
	"log"
	"monelog/model"
	"monelog/usecase"
	"net/http"
	"strconv"

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
	csu usecase.ICardStatementUsecase
	chu usecase.ICSVHistoryUsecase
}

func NewCardStatementController(csu usecase.ICardStatementUsecase, chu usecase.ICSVHistoryUsecase) ICardStatementController {
	return &cardStatementController{csu, chu}
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
	userId, err := getUserIdFromToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "認証に失敗しました")
	}
	
	cardStatementsRes, err := csc.csu.GetAllCardStatements(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, cardStatementsRes)
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
	userId, err := getUserIdFromToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "認証に失敗しました")
	}
	
	id := c.Param("cardStatementId")
	cardStatementId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid card statement ID"})
	}
	
	cardStatementRes, err := csc.csu.GetCardStatementById(userId, uint(cardStatementId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, cardStatementRes)
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
	cardStatementsRes, err := csc.csu.ProcessCSV(file, request)
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
	_, err = csc.chu.SaveCSVHistory(file, csvHistoryRequest)
	if err != nil {
		// CSV履歴の保存に失敗しても、カード明細の処理は続行
		// エラーログを出力するなどの対応が望ましい
		log.Printf("Failed to save CSV history: %v", err)
	}
	
	return c.JSON(http.StatusCreated, cardStatementsRes)
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
	cardStatementsRes, err := csc.csu.PreviewCSV(file, request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	
	return c.JSON(http.StatusOK, cardStatementsRes)
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
	userId, err := getUserIdFromToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "認証に失敗しました")
	}
	
	var request model.CardStatementSaveRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request format"})
	}
	
	request.UserId = userId
	
	// リクエストの検証を強化
	if request.Year <= 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Year must be positive"})
	}
	
	if request.Month < 1 || request.Month > 12 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Month must be between 1 and 12"})
	}
	
	if len(request.CardStatements) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Card statements cannot be empty"})
	}
	
	// カード明細の保存
	cardStatementsRes, err := csc.csu.SaveCardStatements(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	
	return c.JSON(http.StatusCreated, cardStatementsRes)
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
	
	request := model.CardStatementByMonthRequest{
		Year:   year,
		Month:  month,
		UserId: userId,
	}
	
	cardStatementsRes, err := csc.csu.GetCardStatementsByMonth(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	
	return c.JSON(http.StatusOK, cardStatementsRes)
}