package csv_history

import (
	"monelog/model"
)

// GetCSVHistoriesByMonth は指定された年月のCSV履歴を取得します
func (chr *csvHistoryRepository) GetCSVHistoriesByMonth(userId uint, year int, month int) ([]model.CSVHistory, error) {
	var csvHistories []model.CSVHistory
	if err := chr.db.Where("user_id=? AND year=? AND month=?", userId, year, month).
		Order("created_at DESC").Find(&csvHistories).Error; err != nil {
		return nil, err
	}
	return csvHistories, nil
}