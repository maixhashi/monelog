package csv_history

import (
	"monelog/model"
)

// CreateCSVHistory は新しいCSV履歴を作成します
func (chr *csvHistoryRepository) CreateCSVHistory(csvHistory *model.CSVHistory) error {
	return chr.db.Create(csvHistory).Error
}