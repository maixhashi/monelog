package csv_history

import (
	"gorm.io/gorm"
)

type csvHistoryRepository struct {
	db *gorm.DB
}

// NewCSVHistoryRepository はCSV履歴リポジトリの新しいインスタンスを作成します
func NewCSVHistoryRepository(db *gorm.DB) *csvHistoryRepository {
	return &csvHistoryRepository{db}
}