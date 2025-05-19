package csv_history_test

import (
	"mime/multipart"
	"monelog/model"
	"monelog/repository"
	"monelog/testutils"
	"monelog/usecase"
	"monelog/validator"
	"os"
	"testing"

	"gorm.io/gorm"
)

// テスト用の共通変数
var (
	db                  *gorm.DB
	csvHistoryRepo      repository.ICSVHistoryRepository
	csvHistoryValidator validator.ICSVHistoryValidator
	csvHistoryUsecase   usecase.ICSVHistoryUsecase
	testUser            model.User
	otherUser           model.User
)

const nonExistentCSVHistoryID uint = 9999

// テスト前の共通セットアップ
func setupCSVHistoryUsecaseTest() {
	// テストごとにデータベースをクリーンアップ
	if db != nil {
		testutils.CleanupTestDB(db)
	} else {
		// 初回のみデータベース接続を作成
		db = testutils.SetupTestDB()
		csvHistoryRepo = repository.NewCSVHistoryRepository(db)
		csvHistoryValidator = validator.NewCSVHistoryValidator()
		csvHistoryUsecase = usecase.NewCSVHistoryUsecase(csvHistoryRepo, csvHistoryValidator)
	}
	
	// テストユーザーを作成
	testUser = testutils.CreateTestUser(db)
	
	// 別のテストユーザーを作成
	otherUser = testutils.CreateOtherUser(db)
}

// テスト用のCSV履歴を作成
func createTestCSVHistory(t *testing.T, fileName string, userId uint, year int, month int) model.CSVHistory {
	csvHistory := model.CSVHistory{
		FileName: fileName,
		CardType: "楽天カード",
		FileData: []byte("テスト用CSVデータ"),
		UserId:   userId,
		Year:     year,
		Month:    month,
	}
	
	result := db.Create(&csvHistory)
	if result.Error != nil {
		t.Fatalf("テストCSV履歴の作成に失敗しました: %v", result.Error)
	}
	
	return csvHistory
}

// テスト用のCSV履歴保存リクエストを作成
func createTestCSVHistorySaveRequest(userId uint, year int, month int) model.CSVHistorySaveRequest {
	return model.CSVHistorySaveRequest{
		FileName: "test.csv",
		CardType: "楽天カード",
		UserId:   userId,
		Year:     year,
		Month:    month,
	}
}

// テスト用のマルチパートファイルを作成
func createTestMultipartFile(t *testing.T) (*multipart.FileHeader, func()) {
	// テスト用の一時ファイルを作成
	tempFile, err := os.CreateTemp("", "test-*.csv")
	if err != nil {
		t.Fatalf("テスト用一時ファイルの作成に失敗しました: %v", err)
	}
	
	// テストデータを書き込む
	_, err = tempFile.WriteString("テスト用CSVデータ")
	if err != nil {
		t.Fatalf("テスト用ファイルへの書き込みに失敗しました: %v", err)
	}
	
	// ファイルを閉じる
	tempFile.Close()
	
	// マルチパートファイルヘッダを作成（実際のテストでは使用しない模擬オブジェクト）
	fileHeader := &multipart.FileHeader{
		Filename: "test.csv",
		Size:     int64(len("テスト用CSVデータ")),
	}
	
	// クリーンアップ関数を返す
	cleanup := func() {
		os.Remove(tempFile.Name())
	}
	
	return fileHeader, cleanup
}