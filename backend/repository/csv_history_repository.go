package repository

import (
	"monelog/model"
	"monelog/repository/csv_history"

	"gorm.io/gorm"
)

// ICSVHistoryRepository はCSV履歴リポジトリのインターフェースを定義します
type ICSVHistoryRepository interface {
	GetAllCSVHistories(userId uint) ([]model.CSVHistory, error)
	GetCSVHistoryById(userId uint, csvHistoryId uint) (model.CSVHistory, error)
	GetCSVHistoriesByMonth(userId uint, year int, month int) ([]model.CSVHistory, error)
	CreateCSVHistory(csvHistory *model.CSVHistory) error
	DeleteCSVHistory(userId uint, csvHistoryId uint) error
}

// NewCSVHistoryRepository はCSV履歴リポジトリの新しいインスタンスを作成します
func NewCSVHistoryRepository(db *gorm.DB) ICSVHistoryRepository {
	return csv_history.NewCSVHistoryRepository(db)
}
