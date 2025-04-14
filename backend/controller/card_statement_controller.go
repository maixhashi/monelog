package controller

import (
	"monelog/model"
	"monelog/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ICardStatementController interface {
	GetAllCardStatements(c echo.Context) error
	GetCardStatementById(c echo.Context) error
	UploadCSV(c echo.Context) error
	PreviewCSV(c echo.Context) error
	SaveCardStatements(c echo.Context) error
}

type cardStatementController struct {
	csu usecase.ICardStatementUsecase
}

func NewCardStatementController(csu usecase.ICardStatementUsecase) ICardStatementController {
	return &cardStatementController{csu}
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
	
	request := model.CardStatementRequest{
		CardType: cardType,
		UserId:   userId,
	}
	
	// CSVの処理
	cardStatementsRes, err := csc.csu.ProcessCSV(file, request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
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
	
	request := model.CardStatementPreviewRequest{
		CardType: cardType,
		UserId:   userId,
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
	
	// カード明細の保存
	cardStatementsRes, err := csc.csu.SaveCardStatements(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	
	return c.JSON(http.StatusCreated, cardStatementsRes)
}
