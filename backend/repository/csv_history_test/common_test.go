package csv_history_test

import (
	"monelog/model"
	"testing"
)

// CSV履歴の存在を確認するヘルパー関数
func assertCSVHistoryExists(t *testing.T, csvHistoryId uint, expectedFileName string, expectedUserId uint) bool {
	var csvHistory model.CSVHistory
	result := csvHistoryDB.First(&csvHistory, csvHistoryId)
	
	if result.Error != nil {
		t.Errorf("CSV履歴(ID=%d)がデータベースに存在しません: %v", csvHistoryId, result.Error)
		return false
	}
	
	if csvHistory.FileName != expectedFileName {
		t.Errorf("CSV履歴のファイル名が一致しません: got=%s, want=%s", csvHistory.FileName, expectedFileName)
		return false
	}
	
	if csvHistory.UserId != expectedUserId {
		t.Errorf("CSV履歴のユーザーIDが一致しません: got=%d, want=%d", csvHistory.UserId, expectedUserId)
		return false
	}
	
	return true
}

// CSV履歴が存在しないことを確認するヘルパー関数
func assertCSVHistoryNotExists(t *testing.T, csvHistoryId uint) bool {
	var count int64
	csvHistoryDB.Model(&model.CSVHistory{}).Where("id = ?", csvHistoryId).Count(&count)
	
	if count != 0 {
		t.Errorf("CSV履歴(ID=%d)がデータベースに存在します", csvHistoryId)
		return false
	}
	
	return true
}

// CSV履歴の配列を検証するヘルパー関数
func validateCSVHistories(t *testing.T, histories []model.CSVHistory, expectedCount int) bool {
	if len(histories) != expectedCount {
		t.Errorf("CSV履歴の数が一致しません: got=%d, want=%d", len(histories), expectedCount)
		return false
	}
	
	for _, history := range histories {
		if history.ID == 0 || history.FileName == "" || history.UserId == 0 {
			t.Errorf("無効なCSV履歴: %+v", history)
			return false
		}
	}
	
	return true
}