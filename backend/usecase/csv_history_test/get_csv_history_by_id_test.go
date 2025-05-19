package csv_history_test

import (
	"testing"
)

func TestCSVHistoryUsecase_GetCSVHistoryById(t *testing.T) {
	setupCSVHistoryUsecaseTest()
	
	// テストデータの作成
	csvHistory := createTestCSVHistory(t, "テストCSV", testUser.ID, 2023, 1)
	otherCsvHistory := createTestCSVHistory(t, "他ユーザーCSV", otherUser.ID, 2023, 1)
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("存在するCSV履歴を取得できる", func(t *testing.T) {
			response, err := csvHistoryUsecase.GetCSVHistoryById(testUser.ID, csvHistory.ID)
			
			if err != nil {
				t.Errorf("GetCSVHistoryById() error = %v", err)
			}
			
			if response.ID != csvHistory.ID {
				t.Errorf("CSV履歴IDが一致しません: got=%d, want=%d", response.ID, csvHistory.ID)
			}
			
			if response.FileName != csvHistory.FileName {
				t.Errorf("CSV履歴のファイル名が一致しません: got=%s, want=%s", response.FileName, csvHistory.FileName)
			}
			
			// ファイルデータが含まれていることを確認
			if len(response.FileData) == 0 {
				t.Errorf("CSV履歴のファイルデータが空です")
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("存在しないCSV履歴IDを指定するとエラーになる", func(t *testing.T) {
			_, err := csvHistoryUsecase.GetCSVHistoryById(testUser.ID, nonExistentCSVHistoryID)
			
			if err == nil {
				t.Errorf("存在しないCSV履歴IDを指定したのにエラーが発生しませんでした")
			}
		})
		
		t.Run("他ユーザーのCSV履歴IDを指定するとエラーになる", func(t *testing.T) {
			_, err := csvHistoryUsecase.GetCSVHistoryById(testUser.ID, otherCsvHistory.ID)
			
			if err == nil {
				t.Errorf("他ユーザーのCSV履歴IDを指定したのにエラーが発生しませんでした")
			}
		})
	})
}