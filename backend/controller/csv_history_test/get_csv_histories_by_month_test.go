package csv_history_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetCSVHistoriesByMonth(t *testing.T) {
	setupCSVHistoryControllerTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("指定した年月のCSV履歴が存在する場合", func(t *testing.T) {
			// テスト用のCSV履歴を作成
			_ = createTestCSVHistory(t, "test_2023_01_1.csv", csvHistoryTestUser.ID, 2023, 1)
			_ = createTestCSVHistory(t, "test_2023_01_2.csv", csvHistoryTestUser.ID, 2023, 1)
			// 別の月のCSV履歴も作成するが、これは結果に含まれないはず
			_ = createTestCSVHistory(t, "test_2023_02.csv", csvHistoryTestUser.ID, 2023, 2)
			
			// リクエストとレスポンスを作成
			req := httptest.NewRequest(http.MethodGet, "/csv-histories/by-month", nil)
			q := req.URL.Query()
			q.Add("year", "2023")
			q.Add("month", "1")
			req.URL.RawQuery = q.Encode()
			
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			setUserContext(c, csvHistoryTestUser)
			
			// コントローラーを実行
			err := csvHistoryController.GetCSVHistoriesByMonth(c)
			
			// アサーション
			if err != nil {
				t.Fatalf("コントローラーの実行に失敗しました: %v", err)
			}
			
			assertStatusCode(t, rec, http.StatusOK)
			
			// レスポンスを検証（配列形式）
			var responses []map[string]interface{}
			if err := json.Unmarshal(rec.Body.Bytes(), &responses); err != nil {
				t.Errorf("レスポンスのJSONパースに失敗しました: %v", err)
				return
			}
			
			if len(responses) != 2 {
				t.Errorf("CSV履歴の数が一致しません: got=%d, want=2", len(responses))
			}
			
			// 他のユーザーの同じ年月のCSV履歴も作成するが、これは結果に含まれないはず
			_ = createTestCSVHistory(t, "other_2023_01.csv", csvHistoryOtherUser.ID, 2023, 1)
		})
		
		t.Run("指定した年月のCSV履歴が存在しない場合", func(t *testing.T) {
			// リクエストとレスポンスを作成
			req := httptest.NewRequest(http.MethodGet, "/csv-histories/by-month", nil)
			q := req.URL.Query()
			q.Add("year", "2022")
			q.Add("month", "12")
			req.URL.RawQuery = q.Encode()
			
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			setUserContext(c, csvHistoryTestUser)
			
			// コントローラーを実行
			err := csvHistoryController.GetCSVHistoriesByMonth(c)
			
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
			
			if len(responses) != 0 {
				t.Errorf("CSV履歴の数が一致しません: got=%d, want=0", len(responses))
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("年パラメータが不足している場合", func(t *testing.T) {
			// リクエストとレスポンスを作成
			req := httptest.NewRequest(http.MethodGet, "/csv-histories/by-month", nil)
			q := req.URL.Query()
			q.Add("month", "1")
			req.URL.RawQuery = q.Encode()
			
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			setUserContext(c, csvHistoryTestUser)
			
			// コントローラーを実行
			err := csvHistoryController.GetCSVHistoriesByMonth(c)
			
			// アサーション
			if err != nil {
				t.Fatalf("コントローラーの実行に失敗しました: %v", err)
			}
			
			assertStatusCode(t, rec, http.StatusBadRequest)
		})
		
		t.Run("月パラメータが不足している場合", func(t *testing.T) {
			// リクエストとレスポンスを作成
			req := httptest.NewRequest(http.MethodGet, "/csv-histories/by-month", nil)
			q := req.URL.Query()
			q.Add("year", "2023")
			req.URL.RawQuery = q.Encode()
			
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			setUserContext(c, csvHistoryTestUser)
			
			// コントローラーを実行
			err := csvHistoryController.GetCSVHistoriesByMonth(c)
			
			// アサーション
			if err != nil {
				t.Fatalf("コントローラーの実行に失敗しました: %v", err)
			}
			
			assertStatusCode(t, rec, http.StatusBadRequest)
		})
		
		t.Run("無効な年パラメータの場合", func(t *testing.T) {
			// リクエストとレスポンスを作成
			req := httptest.NewRequest(http.MethodGet, "/csv-histories/by-month", nil)
			q := req.URL.Query()
			q.Add("year", "invalid")
			q.Add("month", "1")
			req.URL.RawQuery = q.Encode()
			
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			setUserContext(c, csvHistoryTestUser)
			
			// コントローラーを実行
			err := csvHistoryController.GetCSVHistoriesByMonth(c)
			
			// アサーション
			if err != nil {
				t.Fatalf("コントローラーの実行に失敗しました: %v", err)
			}
			
			assertStatusCode(t, rec, http.StatusBadRequest)
		})
		
		t.Run("無効な月パラメータの場合", func(t *testing.T) {
			// リクエストとレスポンスを作成
			req := httptest.NewRequest(http.MethodGet, "/csv-histories/by-month", nil)
			q := req.URL.Query()
			q.Add("year", "2023")
			q.Add("month", "invalid")
			req.URL.RawQuery = q.Encode()
			
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			setUserContext(c, csvHistoryTestUser)
			
			// コントローラーを実行
			err := csvHistoryController.GetCSVHistoriesByMonth(c)
			
			// アサーション
			if err != nil {
				t.Fatalf("コントローラーの実行に失敗しました: %v", err)
			}
			
			assertStatusCode(t, rec, http.StatusBadRequest)
		})
		
		t.Run("月の範囲が無効な場合", func(t *testing.T) {
			// リクエストとレスポンスを作成（月が13の場合）
			req := httptest.NewRequest(http.MethodGet, "/csv-histories/by-month", nil)
			q := req.URL.Query()
			q.Add("year", "2023")
			q.Add("month", "13") // 無効な月
			req.URL.RawQuery = q.Encode()
			
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			setUserContext(c, csvHistoryTestUser)
			
			// コントローラーを実行
			err := csvHistoryController.GetCSVHistoriesByMonth(c)
			
			// アサーション
			if err != nil {
				t.Fatalf("コントローラーの実行に失敗しました: %v", err)
			}
			
			assertStatusCode(t, rec, http.StatusBadRequest)
		})
	})
}