package csv_history_test

import (
	"encoding/json"
	"monelog/model"
	"net/http"
	"testing"
)

func TestGetAllCSVHistories(t *testing.T) {
	setupCSVHistoryControllerTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("CSV履歴が存在する場合", func(t *testing.T) {
			// テスト用のCSV履歴を作成
			_ = createTestCSVHistory(t, "test_all_1.csv", csvHistoryTestUser.ID, 2023, 1)
			_ = createTestCSVHistory(t, "test_all_2.csv", csvHistoryTestUser.ID, 2023, 2)
			
			// リクエストとレスポンスを作成
			rec, c := createRequestResponse(http.MethodGet, "/csv-histories")
			setUserContext(c, csvHistoryTestUser)
			
			// コントローラーを実行
			err := csvHistoryController.GetAllCSVHistories(c)
			
			// アサーション
			if err != nil {
				t.Fatalf("コントローラーの実行に失敗しました: %v", err)
			}
			
			assertStatusCode(t, rec, http.StatusOK)
			
			// レスポンスを検証
			var responses []map[string]interface{}
			if err := json.Unmarshal(rec.Body.Bytes(), &responses); err != nil {
				t.Errorf("レスポンスのJSONパースに失敗しました: %v", err)
				return
			}
			
			// 少なくとも2つのCSV履歴があることを確認
			if len(responses) < 2 {
				t.Errorf("CSV履歴の数が不足しています: got=%d, want>=2", len(responses))
			}
			
			// レスポンスの内容を検証
			validateCSVHistoryResponses(t, responses, len(responses))
		})
		
		t.Run("CSV履歴が存在しない場合", func(t *testing.T) {
			// 一時的にテストユーザーを作成
			tempUser := model.User{
				Email:    "temp@example.com",
				Password: "password",
			}
			result := csvHistoryDB.Create(&tempUser)
			if result.Error != nil {
				t.Fatalf("一時ユーザーの作成に失敗しました: %v", result.Error)
			}
			
			// リクエストとレスポンスを作成
			rec, c := createRequestResponse(http.MethodGet, "/csv-histories")
			setUserContext(c, tempUser)
			
			// コントローラーを実行
			err := csvHistoryController.GetAllCSVHistories(c)
			
			// アサーション
			if err != nil {
				t.Fatalf("コントローラーの実行に失敗しました: %v", err)
			}
			
			assertStatusCode(t, rec, http.StatusOK)
			
			// レスポンスを検証（空の配列）
			var responses []map[string]interface{}
			if err := json.Unmarshal(rec.Body.Bytes(), &responses); err != nil {
				t.Errorf("レスポンスのJSONパースに失敗しました: %v", err)
				return
			}
			
			// CSV履歴が0件であることを確認
			validateCSVHistoryResponses(t, responses, 0)
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("認証されていない場合", func(t *testing.T) {
			// リクエストとレスポンスを作成（ユーザーコンテキストを設定しない）
			_, c := createRequestResponse(http.MethodGet, "/csv-histories")
			
			// コントローラーを実行
			err := csvHistoryController.GetAllCSVHistories(c)
			
			// エラーが返されることを期待
			if err == nil {
				t.Fatalf("認証エラーが発生しませんでした")
			}
		})
	})
}