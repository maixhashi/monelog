package csv_history_test

import (
	"fmt"
	"net/http"
	"testing"
)

func TestDeleteCSVHistory(t *testing.T) {
	setupCSVHistoryControllerTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("自分のCSV履歴を削除する場合", func(t *testing.T) {
			// テスト用のCSV履歴を作成
			csvHistory := createTestCSVHistory(t, "test_delete.csv", csvHistoryTestUser.ID, 2023, 1)
			
			// リクエストとレスポンスを作成
			rec, c := createRequestResponse(http.MethodDelete, fmt.Sprintf("/csv-histories/%d", csvHistory.ID))
			setUserContext(c, csvHistoryTestUser)
			c.SetParamNames("csvHistoryId")
			c.SetParamValues(fmt.Sprintf("%d", csvHistory.ID))
			
			// コントローラーを実行
			err := csvHistoryController.DeleteCSVHistory(c)
			
			// アサーション
			if err != nil {
				t.Fatalf("コントローラーの実行に失敗しました: %v", err)
			}
			
			assertStatusCode(t, rec, http.StatusNoContent)
			assertCSVHistoryNotExists(t, csvHistory.ID)
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("存在しないCSV履歴を削除しようとした場合", func(t *testing.T) {
			// リクエストとレスポンスを作成
			rec, c := createRequestResponse(http.MethodDelete, fmt.Sprintf("/csv-histories/%d", nonExistentCSVHistoryID))
			setUserContext(c, csvHistoryTestUser)
			c.SetParamNames("csvHistoryId")
			c.SetParamValues(fmt.Sprintf("%d", nonExistentCSVHistoryID))
			
			// コントローラーを実行
			err := csvHistoryController.DeleteCSVHistory(c)
			
			// アサーション
			if err != nil {
				t.Fatalf("コントローラーの実行に失敗しました: %v", err)
			}
			
			// 実際のステータスコードを確認
			if rec.Code != http.StatusNotFound && rec.Code != http.StatusInternalServerError {
				t.Errorf("予期しないステータスコード: got=%d, want=%d or %d", 
					rec.Code, http.StatusNotFound, http.StatusInternalServerError)
			}
		})
		
		t.Run("他のユーザーのCSV履歴を削除しようとした場合", func(t *testing.T) {
			// 別のユーザーのCSV履歴を作成
			otherCsvHistory := createTestCSVHistory(t, "other_delete.csv", csvHistoryOtherUser.ID, 2023, 1)
			
			// リクエストとレスポンスを作成
			rec, c := createRequestResponse(http.MethodDelete, fmt.Sprintf("/csv-histories/%d", otherCsvHistory.ID))
			setUserContext(c, csvHistoryTestUser)
			c.SetParamNames("csvHistoryId")
			c.SetParamValues(fmt.Sprintf("%d", otherCsvHistory.ID))
			
			// コントローラーを実行
			err := csvHistoryController.DeleteCSVHistory(c)
			
			// アサーション
			if err != nil {
				t.Fatalf("コントローラーの実行に失敗しました: %v", err)
			}
			
			// 実際のステータスコードを確認
			if rec.Code != http.StatusForbidden && 
				rec.Code != http.StatusInternalServerError && 
				rec.Code != http.StatusNotFound {
				t.Errorf("予期しないステータスコード: got=%d, want=%d, %d or %d", 
					rec.Code, http.StatusForbidden, http.StatusInternalServerError, http.StatusNotFound)
			}
		})
		
		t.Run("無効なIDパラメータの場合", func(t *testing.T) {
			// リクエストとレスポンスを作成
			rec, c := createRequestResponse(http.MethodDelete, "/csv-histories/invalid")
			setUserContext(c, csvHistoryTestUser)
			c.SetParamNames("csvHistoryId")
			c.SetParamValues("invalid")
			
			// コントローラーを実行
			err := csvHistoryController.DeleteCSVHistory(c)
			
			// アサーション
			if err != nil {
				t.Fatalf("コントローラーの実行に失敗しました: %v", err)
			}
			
			validateErrorResponse(t, rec, http.StatusBadRequest)
		})
	})
}