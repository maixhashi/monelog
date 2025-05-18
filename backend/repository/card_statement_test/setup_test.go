package card_statement_test

import (
	"monelog/model"
	"monelog/repository"  // 変更: repositoryパッケージをインポート
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type testDB struct {
	db                     *gorm.DB
	cardStatementRepository repository.ICardStatementRepository  // 変更: インターフェース型を使用
}

func setupTestDB(t *testing.T) *testDB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}

	// マイグレーション
	err = db.AutoMigrate(&model.CardStatement{})
	if err != nil {
		t.Fatalf("failed to migrate database: %v", err)
	}

	return &testDB{
		db:                     db,
		cardStatementRepository: repository.NewCardStatementRepository(db),  // 変更: 新しいファクトリ関数を使用
	}
}

func (tdb *testDB) cleanup(t *testing.T) {
	sqlDB, err := tdb.db.DB()
	if err != nil {
		t.Fatalf("failed to get database: %v", err)
	}
	sqlDB.Close()
}
