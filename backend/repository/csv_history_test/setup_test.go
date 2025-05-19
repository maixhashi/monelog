package csv_history_test

import (
	"monelog/model"
	"monelog/repository"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	csvHistoryDB       *gorm.DB
	csvHistoryRepo     repository.ICSVHistoryRepository
	csvHistoryTestUser uint = 1
)

// テスト環境のセットアップ
func setupCSVHistoryTest(t *testing.T) {
	var err error
	csvHistoryDB, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}

	// マイグレーション
	err = csvHistoryDB.AutoMigrate(&model.CSVHistory{})
	if err != nil {
		t.Fatalf("failed to migrate database: %v", err)
	}

	// リポジトリの初期化
	csvHistoryRepo = repository.NewCSVHistoryRepository(csvHistoryDB)
}

// テスト環境のクリーンアップ
func cleanupCSVHistoryTest(t *testing.T) {
	sqlDB, err := csvHistoryDB.DB()
	if err != nil {
		t.Fatalf("failed to get database: %v", err)
	}
	sqlDB.Close()
}