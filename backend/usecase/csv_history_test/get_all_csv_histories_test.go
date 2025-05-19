package csv_history_test

import (
	"testing"
)

func TestCSVHistoryUsecase_GetAllCSVHistories(t *testing.T) {
	setupCSVHistoryUsecaseTest()
	
	// テストデータの作成
	createTestCSVHistory(t, "テストCSV1", testUser.ID, 2023, 1)
	createTestCSVHistory(t, "テストCSV2", testUser.ID, 2023, 2)
	createTestCSVHistory(t, "他ユーザーCSV", otherUser.ID, 2023, 1)
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("正しいユーザーIDのCSV履歴のみを取得する", func(t *testing.T) {
			t.Logf("ユーザーID %d のCSV履歴を取得します", testUser.ID)
			
			csvHistoryResponses, err := csvHistoryUsecase.GetAllCSVHistories(testUser.ID)
			
			if err != nil {
				t.Errorf("GetAllCSVHistories() error = %v", err)
			}
			
			if !validateCSVHistoryResponses(t, csvHistoryResponses, 2) {
				t.Fatalf("CSV履歴レスポンスの検証に失敗しました")
			}
			
			// ファイル名の確認
			fileNames := make(map[string]bool)
			for _, csvHistory := range csvHistoryResponses {
				fileNames[csvHistory.FileName] = true
				t.Logf("取得したCSV履歴: ID=%d, FileName=%s", csvHistory.ID, csvHistory.FileName)
			}
			
			if !fileNames["テストCSV1"] || !fileNames["テストCSV2"] {
				t.Errorf("期待したCSV履歴が結果に含まれていません: %v", csvHistoryResponses)
			}
		})
	})
}