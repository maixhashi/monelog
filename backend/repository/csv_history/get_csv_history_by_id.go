package csv_history

import (
	"fmt"
	"monelog/model"
)

// GetCSVHistoryById は指定されたIDのCSV履歴を取得します
func (chr *csvHistoryRepository) GetCSVHistoryById(userId uint, csvHistoryId uint) (model.CSVHistory, error) {
	var csvHistory model.CSVHistory
	result := chr.db.Where("user_id=?", userId).First(&csvHistory, csvHistoryId)
	if result.Error != nil {
		return model.CSVHistory{}, result.Error
	}
	if result.RowsAffected == 0 {
		return model.CSVHistory{}, fmt.Errorf("CSV history not found")
	}
	return csvHistory, nil
}