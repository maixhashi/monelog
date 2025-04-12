package card_statement_test

import (
	"fmt"
	"monelog/controller"
	"monelog/model"
	"monelog/repository"
	"monelog/testutils"
	"monelog/usecase"
	"monelog/validator"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// テスト用の共通変数
var (
	cardStatementDB          *gorm.DB
	cardStatementRepo        repository.ICardStatementRepository
	cardStatementValidator   validator.ICardStatementValidator
	cardStatementUsecase     usecase.ICardStatementUsecase
	cardStatementController  controller.ICardStatementController
	cardStatementTestUser    model.User
	cardStatementOtherUser   model.User
	nonExistentCardStatementID uint = 9999
)

// テストセットアップ関数
func setupCardStatementControllerTest() {
	// テストごとにデータベースをクリーンアップ
	if cardStatementDB != nil {
		testutils.CleanupTestDB(cardStatementDB)
	} else {
		// 初回のみデータベース接続を作成
		cardStatementDB = testutils.SetupTestDB()
		cardStatementRepo = repository.NewCardStatementRepository(cardStatementDB)
		cardStatementValidator = validator.NewCardStatementValidator()
		cardStatementUsecase = usecase.NewCardStatementUsecase(cardStatementRepo, cardStatementValidator)
		cardStatementController = controller.NewCardStatementController(cardStatementUsecase)
	}
	
	// テストユーザーを作成
	cardStatementTestUser = testutils.CreateTestUser(cardStatementDB)
	cardStatementOtherUser = testutils.CreateOtherUser(cardStatementDB)
}

// JWT認証をモックするヘルパー関数
func setupEchoWithJWT(userId uint) (*echo.Echo, echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	
	// JWTクレームを設定
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = float64(userId)
	
	// コンテキストにトークンを設定
	c.Set("user", token)
	
	return e, c, rec
}

// リクエストボディ付きのコンテキストを設定するヘルパー関数
func setupEchoWithJWTAndBody(userId uint, method, path, body string) (*echo.Echo, echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(method, path, strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	
	// JWTクレームを設定
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = float64(userId)
	c.Set("user", token)
	
	return e, c, rec
}

// CardStatementIDパラメータを持つリクエストコンテキストを設定するヘルパー関数
func setupEchoWithCardStatementId(userId uint, cardStatementId uint, method, path, body string) (*echo.Echo, echo.Context, *httptest.ResponseRecorder) {
	e, c, rec := setupEchoWithJWTAndBody(userId, method, path, body)
	c.SetParamNames("cardStatementId")
	c.SetParamValues(fmt.Sprintf("%d", cardStatementId))
	return e, c, rec
}
