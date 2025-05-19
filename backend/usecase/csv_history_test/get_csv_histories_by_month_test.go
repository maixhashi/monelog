package csv_history_test

import (
	"testing"
)

func TestCSVHistoryUsecase_GetCSVHistoriesByMonth(t *testing.T) {
	setupCSVHistoryUsecaseTest()
	
	// テストデータの作成
	createTestCSVHistory(t, "1月CSV1", testUser.ID, 2023, 1)
	createTestCSVHistory(t, "1月CSV2", testUser.ID, 2023, 1)
	createTestCSVHistory(t, "2月CSV", testUser.ID, 2023, 2)
	createTestCSVHistory(t, "他ユーザー1月CSV", otherUser.ID, 2023, 1)
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("指定した年月のCSV履歴のみを取得する", func(t *testing.T) {
			t.Logf("ユーザーID %d の2023年1月のCSV履歴を取得します", testUser.ID)
			
			csvHistoryResponses, err := csvHistoryUsecase.GetCSVHistoriesByMonth(testUser.ID, 2023, 1)
			
			if err != nil {
				t.Errorf("GetCSVHistoriesByMonth() error = %v", err)
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
			
			if !fileNames["1月CSV1"] || !fileNames["1月CSV2"] {
				t.Errorf("期待したCSV履歴が結果に含まれていません: %v", csvHistoryResponses)
			}
		})
		
		t.Run("別の月のCSV履歴を取得する", func(t *testing.T) {
			t.Logf("ユーザーID %d の2023年2月のCSV履歴を取得します", testUser.ID)
			
			csvHistoryResponses, err := csvHistoryUsecase.GetCSVHistoriesByMonth(testUser.ID, 2023, 2)
			
			if err != nil {
				t.Errorf("GetCSVHistoriesByMonth() error = %v", err)
			}
			
			if !validateCSVHistoryResponses(t, csvHistoryResponses, 1) {
				t.Fatalf("CSV履歴レスポンスの検証に失敗しました")
			}
			
			// ファイル名の確認
			if csvHistoryResponses[0].FileName != "2月CSV" {
				t.Errorf("期待したCSV履歴が結果に含まれていません: got=%s, want=%s", 
					csvHistoryResponses[0].FileName, "2月CSV")
			}
		})
		
		t.Run("存在しない年月のCSV履歴を取得すると空の配列が返る", func(t *testing.T) {
			csvHistoryResponses, err := csvHistoryUsecase.GetCSVHistoriesByMonth(testUser.ID, 2022, 12)
			
			if err != nil {
				t.Errorf("GetCSVHistoriesByMonth() error = %v", err)
			}
			
			if len(csvHistoryResponses) != 0 {
				t.Errorf("存在しない年月のCSV履歴を取得したのに結果が空ではありません: %v", csvHistoryResponses)
			}
		})
	})
}