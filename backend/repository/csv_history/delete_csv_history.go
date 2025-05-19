package csv_history

import (
	"fmt"
	"monelog/model"
)

// DeleteCSVHistory は指定されたIDのCSV履歴を削除します
func (chr *csvHistoryRepository) DeleteCSVHistory(userId uint, csvHistoryId uint) error {
	result := chr.db.Where("user_id=? AND id=?", userId, csvHistoryId).Delete(&model.CSVHistory{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("CSV history not found or not authorized to delete")
	}
	return nil
}