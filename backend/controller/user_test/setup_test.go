package user_test

import (
	"monelog/controller"
	"monelog/repository"
	"monelog/testutils"
	"monelog/usecase"
	"monelog/validator"
	"net/http/httptest"
	"strings"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// テスト用の共通変数
var (
	userDb          *gorm.DB
	userRepo        repository.IUserRepository
	userValidator   validator.IUserValidator
	userUsecase     usecase.IUserUsecase
	userController  controller.IUserController
)

// テストセットアップ関数
func setupUserControllerTest() {
	// テストごとにデータベースをクリーンアップ
	if userDb != nil {
		testutils.CleanupTestDB(userDb)
	} else {
		// 初回のみデータベース接続を作成
		userDb = testutils.SetupTestDB()
		userRepo = repository.NewUserRepository(userDb)
		userValidator = validator.NewUserValidator()
		userUsecase = usecase.NewUserUsecase(userRepo, userValidator)
		userController = controller.NewUserController(userUsecase)
	}

	// 既存のテストユーザーを明示的に削除（念のため）
	userDb.Exec("DELETE FROM users WHERE email LIKE 'test%@example.com'")
	userDb.Exec("DELETE FROM users WHERE email = 'logintest@example.com'")
	userDb.Exec("DELETE FROM users WHERE email = 'duplicate@example.com'")
}

// リクエストボディ付きのコンテキストを設定するヘルパー関数
func setupEchoContextWithBody(method, path, body string) (*echo.Echo, echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(method, path, strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	return e, c, rec
}

// CSRFトークンを持つコンテキストをセットアップするヘルパー関数
func setupEchoContextWithCSRF(method, path string) (*echo.Echo, echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(method, path, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	// CSRFトークンをモック
	c.Set("csrf", "test-csrf-token")
	return e, c, rec
}
