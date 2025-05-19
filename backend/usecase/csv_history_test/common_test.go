package csv_history_test

import (
	"monelog/model"
	"testing"
)

// CSV履歴の存在を確認するヘルパー関数
func assertCSVHistoryExists(t *testing.T, csvHistoryId uint, expectedFileName string, expectedUserId uint) bool {
	var csvHistory model.CSVHistory
	result := db.First(&csvHistory, csvHistoryId)
	
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
	db.Model(&model.CSVHistory{}).Where("id = ?", csvHistoryId).Count(&count)
	
	if count != 0 {
		t.Errorf("CSV履歴(ID=%d)がデータベースに存在します", csvHistoryId)
		return false
	}
	
	return true
}

// CSV履歴レスポンスの検証ヘルパー関数
func validateCSVHistoryResponse(t *testing.T, csvHistory model.CSVHistoryResponse, expectedId uint, expectedFileName string) bool {
	if csvHistory.ID != expectedId {
		t.Errorf("CSV履歴IDが一致しません: got=%d, want=%d", csvHistory.ID, expectedId)
		return false
	}
	
	if csvHistory.FileName != expectedFileName {
		t.Errorf("CSV履歴のファイル名が一致しません: got=%s, want=%s", csvHistory.FileName, expectedFileName)
		return false
	}
	
	// 作成日時が設定されていることを確認
	if csvHistory.CreatedAt.IsZero() {
		t.Errorf("CSV履歴の作成日時が正しく設定されていません")
		return false
	}
	
	return true
}

// CSV履歴レスポンスの配列を検証するヘルパー関数
func validateCSVHistoryResponses(t *testing.T, responses []model.CSVHistoryResponse, expectedCount int) bool {
	if len(responses) != expectedCount {
		t.Errorf("CSV履歴の数が一致しません: got=%d, want=%d", len(responses), expectedCount)
		return false
	}
	
	for _, response := range responses {
		if response.ID == 0 || response.FileName == "" || response.CreatedAt.IsZero() {
			t.Errorf("無効なCSV履歴レスポンス: %+v", response)
			return false
		}
	}
	
	return true
}