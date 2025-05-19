package csv_history

import (
	"monelog/model"
)

// GetAllCSVHistories はユーザーのすべてのCSV履歴を取得します
func (chr *csvHistoryRepository) GetAllCSVHistories(userId uint) ([]model.CSVHistory, error) {
	var csvHistories []model.CSVHistory
	if err := chr.db.Where("user_id=?", userId).Order("created_at DESC").Find(&csvHistories).Error; err != nil {
		return nil, err
	}
	return csvHistories, nil
}