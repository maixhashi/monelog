package task_test

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
	taskDB          *gorm.DB
	taskRepo        repository.ITaskRepository
	taskValidator   validator.ITaskValidator
	taskUsecase     usecase.ITaskUsecase
	taskController  controller.ITaskController
	taskTestUser    model.User
	taskOtherUser   model.User
	nonExistentTaskID uint = 9999
)

// テストセットアップ関数
func setupTaskControllerTest() {
	// テストごとにデータベースをクリーンアップ
	if taskDB != nil {
		testutils.CleanupTestDB(taskDB)
	} else {
		// 初回のみデータベース接続を作成
		taskDB = testutils.SetupTestDB()
		taskRepo = repository.NewTaskRepository(taskDB)
		taskValidator = validator.NewTaskValidator()
		taskUsecase = usecase.NewTaskUsecase(taskRepo, taskValidator)
		taskController = controller.NewTaskController(taskUsecase)
	}
	
	// テストユーザーを作成
	taskTestUser = testutils.CreateTestUser(taskDB)
	taskOtherUser = testutils.CreateOtherUser(taskDB)
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

// TaskIDパラメータを持つリクエストコンテキストを設定するヘルパー関数
func setupEchoWithTaskId(userId uint, taskId uint, method, path, body string) (*echo.Echo, echo.Context, *httptest.ResponseRecorder) {
	e, c, rec := setupEchoWithJWTAndBody(userId, method, path, body)
	c.SetParamNames("taskId")
	c.SetParamValues(fmt.Sprintf("%d", taskId))
	return e, c, rec
}