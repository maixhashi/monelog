package csv_history_test

import (
	"encoding/json"
	"monelog/controller"
	"monelog/model"
	"monelog/repository"
	"monelog/testutils"
	"monelog/usecase"
	"monelog/validator"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// テスト用の共通変数
var (
	csvHistoryDB          *gorm.DB
	csvHistoryRepo        repository.ICSVHistoryRepository
	csvHistoryValidator   validator.ICSVHistoryValidator
	csvHistoryUsecase     usecase.ICSVHistoryUsecase
	csvHistoryController  controller.ICSVHistoryController
	csvHistoryTestUser    model.User
	csvHistoryOtherUser   model.User
	nonExistentCSVHistoryID uint = 9999
	e                     *echo.Echo
)

// テストセットアップ関数
func setupCSVHistoryControllerTest() {
	// テストごとにデータベースをクリーンアップ
	if csvHistoryDB != nil {
		testutils.CleanupTestDB(csvHistoryDB)
	} else {
		// 初回のみデータベース接続を作成
		csvHistoryDB = testutils.SetupTestDB()
		csvHistoryRepo = repository.NewCSVHistoryRepository(csvHistoryDB)
		csvHistoryValidator = validator.NewCSVHistoryValidator()
		csvHistoryUsecase = usecase.NewCSVHistoryUsecase(csvHistoryRepo, csvHistoryValidator)
		csvHistoryController = controller.NewCSVHistoryController(csvHistoryUsecase)
		e = echo.New()
	}
	
	// テストユーザーを作成
	csvHistoryTestUser = testutils.CreateTestUser(csvHistoryDB)
	csvHistoryOtherUser = testutils.CreateOtherUser(csvHistoryDB)
}

// テスト用のCSV履歴を作成するヘルパー関数
func createTestCSVHistory(t *testing.T, fileName string, userId uint, year int, month int) model.CSVHistory {
	csvHistory := model.CSVHistory{
		FileName: fileName,
		CardType: "rakuten",
		FileData: []byte("test,data\n1,2"),
		Year:     year,
		Month:    month,
		UserId:   userId,
	}
	
	result := csvHistoryDB.Create(&csvHistory)
	if result.Error != nil {
		t.Fatalf("テストCSV履歴の作成に失敗しました: %v", result.Error)
	}
	
	return csvHistory
}

// リクエストとレスポンスを作成するヘルパー関数
func createRequestResponse(method, path string) (*httptest.ResponseRecorder, echo.Context) {
	req := httptest.NewRequest(method, path, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	return rec, c
}

// ユーザーコンテキストを設定するヘルパー関数
func setUserContext(c echo.Context, user model.User) {
	// テスト環境では model.User 型でセットする
	c.Set("user", user)
}

// ステータスコードを検証するヘルパー関数
func assertStatusCode(t *testing.T, rec *httptest.ResponseRecorder, expectedCode int) {
	if rec.Code != expectedCode {
		t.Errorf("ステータスコードが一致しません: got=%d, want=%d", rec.Code, expectedCode)
	}
}

// エラーレスポンスを検証するヘルパー関数
func validateErrorResponse(t *testing.T, rec *httptest.ResponseRecorder, expectedCode int) {
	assertStatusCode(t, rec, expectedCode)
	
	// レスポンスが空の場合はスキップ
	if len(rec.Body.Bytes()) == 0 {
		return
	}
	
	// まず、文字列マップとしてパースを試みる
	var stringResponse map[string]string
	err1 := json.Unmarshal(rec.Body.Bytes(), &stringResponse)
	if err1 == nil {
		if _, exists := stringResponse["error"]; !exists {
			t.Errorf("エラーレスポンスにerrorフィールドがありません: %v", stringResponse)
		}
		return
	}
	
	// 次に、インターフェースマップとしてパースを試みる
	var interfaceResponse map[string]interface{}
	err2 := json.Unmarshal(rec.Body.Bytes(), &interfaceResponse)
	if err2 == nil {
		if _, exists := interfaceResponse["error"]; !exists {
			t.Errorf("エラーレスポンスにerrorフィールドがありません: %v", interfaceResponse)
		}
		return
	}
	
	// どちらの形式でもパースできない場合はエラーを報告
	t.Errorf("エラーレスポンスのJSONパースに失敗しました: %v, %v", err1, err2)
}

// レスポンスのJSONをパースするヘルパー関数
func parseJSONResponse(t *testing.T, rec *httptest.ResponseRecorder, v interface{}) {
	if err := json.Unmarshal(rec.Body.Bytes(), v); err != nil {
		t.Errorf("レスポンスのJSONパースに失敗しました: %v", err)
	}
}

// CSV履歴が存在しないことを確認するヘルパー関数
func assertCSVHistoryNotExists(t *testing.T, csvHistoryId uint) {
	var count int64
	csvHistoryDB.Model(&model.CSVHistory{}).Where("id = ?", csvHistoryId).Count(&count)
	
	if count != 0 {
		t.Errorf("CSV履歴(ID=%d)がデータベースに存在します", csvHistoryId)
	}
}

// CSV履歴レスポンスの検証ヘルパー関数
func validateCSVHistoryResponse(t *testing.T, response map[string]interface{}, expectedId uint, expectedFileName string) {
	id, ok := response["id"].(float64)
	if !ok || uint(id) != expectedId {
		t.Errorf("CSV履歴IDが一致しません: got=%v, want=%d", response["id"], expectedId)
	}
	
	fileName, ok := response["file_name"].(string)
	if !ok || fileName != expectedFileName {
		t.Errorf("CSV履歴のファイル名が一致しません: got=%v, want=%s", response["file_name"], expectedFileName)
	}
}

// CSV履歴レスポンスの配列を検証するヘルパー関数
func validateCSVHistoryResponses(t *testing.T, responses []map[string]interface{}, expectedCount int) {
	if len(responses) != expectedCount {
		t.Errorf("CSV履歴の数が一致しません: got=%d, want=%d", len(responses), expectedCount)
		return
	}
	
	for _, response := range responses {
		id, ok := response["id"].(float64)
		if !ok || id == 0 {
			t.Errorf("無効なCSV履歴ID: %v", response["id"])
		}
		
		fileName, ok := response["file_name"].(string)
		if !ok || fileName == "" {
			t.Errorf("無効なCSV履歴ファイル名: %v", response["file_name"])
		}
	}
}