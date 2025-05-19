package csv_history_test

import (
	"testing"
)

func TestCSVHistoryUsecase_DeleteCSVHistory(t *testing.T) {
	setupCSVHistoryUsecaseTest()
	
	// テストデータの作成
	csvHistory := createTestCSVHistory(t, "削除テストCSV", testUser.ID, 2023, 1)
	otherCsvHistory := createTestCSVHistory(t, "他ユーザーCSV", otherUser.ID, 2023, 1)
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("自分のCSV履歴を削除できる", func(t *testing.T) {
			err := csvHistoryUsecase.DeleteCSVHistory(testUser.ID, csvHistory.ID)
			
			if err != nil {
				t.Errorf("DeleteCSVHistory() error = %v", err)
			}
			
			// 削除されたことを確認
			if !assertCSVHistoryNotExists(t, csvHistory.ID) {
				t.Fatalf("CSV履歴が削除されていません")
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("存在しないCSV履歴IDを指定するとエラーになる", func(t *testing.T) {
			err := csvHistoryUsecase.DeleteCSVHistory(testUser.ID, nonExistentCSVHistoryID)
			
			if err == nil {
				t.Errorf("存在しないCSV履歴IDを指定したのにエラーが発生しませんでした")
			}
		})
		
		t.Run("他ユーザーのCSV履歴IDを指定するとエラーになる", func(t *testing.T) {
			err := csvHistoryUsecase.DeleteCSVHistory(testUser.ID, otherCsvHistory.ID)
			
			if err == nil {
				t.Errorf("他ユーザーのCSV履歴IDを指定したのにエラーが発生しませんでした")
			}
			
			// 削除されていないことを確認
			if !assertCSVHistoryExists(t, otherCsvHistory.ID, "他ユーザーCSV", otherUser.ID) {
				t.Fatalf("他ユーザーのCSV履歴が誤って削除されました")
			}
		})
	})
}