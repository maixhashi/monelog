package card_statement_test

import (
	"net/http"
	"testing"
)

func TestCardStatementController_GetCardStatementsByMonth(t *testing.T) {
	setupCardStatementControllerTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("指定された年月のカード明細を取得する", func(t *testing.T) {
			// テスト用カード明細の作成
			createTestCardStatement("楽天カード", "Amazon.co.jp", cardStatementTestUser.ID, 2023, 4)
			createTestCardStatement("楽天カード", "楽天市場", cardStatementTestUser.ID, 2023, 4)
			createTestCardStatement("楽天カード", "別の月の明細", cardStatementTestUser.ID, 2023, 5) // 別の月の明細
			createTestCardStatement("楽天カード", "他ユーザーの明細", cardStatementOtherUser.ID, 2023, 4) // 別ユーザーの明細
			
			// クエリパラメータの設定
			queryParams := map[string]string{
				"year":  "2023",
				"month": "4",
			}
			
			// テスト実行
			_, c, rec := setupEchoWithJWTAndQuery(cardStatementTestUser.ID, http.MethodGet, "/card-statements/by-month", queryParams)
			err := cardStatementController.GetCardStatementsByMonth(c)
			
			// 検証
			if err != nil {
				t.Errorf("GetCardStatementsByMonth() error = %v", err)
			}
			
			if rec.Code != http.StatusOK {
				t.Errorf("GetCardStatementsByMonth() status code = %d, want %d", rec.Code, http.StatusOK)
			}
			
			// レスポンスボディをパース
			response := parseCardStatementsResponse(t, rec.Body.Bytes())
			
			if len(response) != 2 {
				t.Errorf("GetCardStatementsByMonth() returned %d card statements, want 2", len(response))
			}
			
			// カード明細の説明の確認
			descriptions := make(map[string]bool)
			for _, cs := range response {
				descriptions[cs.Description] = true
				
				// 年月が正しいことを確認
				if cs.Year != 2023 || cs.Month != 4 {
					t.Errorf("GetCardStatementsByMonth() returned statement with year=%d, month=%d, want year=2023, month=4", cs.Year, cs.Month)
				}
			}
			
			if !descriptions["Amazon.co.jp"] || !descriptions["楽天市場"] {
				t.Errorf("期待したカード明細が結果に含まれていません: %v", response)
			}
			
			// 別の月の明細が含まれていないことを確認
			if descriptions["別の月の明細"] {
				t.Errorf("別の月の明細が結果に含まれています: %v", response)
			}
			
			// 他ユーザーの明細が含まれていないことを確認
			if descriptions["他ユーザーの明細"] {
				t.Errorf("他ユーザーの明細が結果に含まれています: %v", response)
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("年が指定されていない場合エラーになる", func(t *testing.T) {
			// 年なしのクエリパラメータ
			queryParams := map[string]string{
				"month": "4",
			}
			
			// テスト実行
			_, c, rec := setupEchoWithJWTAndQuery(cardStatementTestUser.ID, http.MethodGet, "/card-statements/by-month", queryParams)
			err := cardStatementController.GetCardStatementsByMonth(c)
			
			// 検証
			if err != nil {
				t.Errorf("GetCardStatementsByMonth() error = %v", err)
				return
			}
			
			// エラーレスポンスが返されることを確認
			if rec.Code != http.StatusBadRequest {
				t.Errorf("GetCardStatementsByMonth() status code = %d, want %d", rec.Code, http.StatusBadRequest)
			}
		})
		
		t.Run("月が指定されていない場合エラーになる", func(t *testing.T) {
			// 月なしのクエリパラメータ
			queryParams := map[string]string{
				"year": "2023",
			}
			
			// テスト実行
			_, c, rec := setupEchoWithJWTAndQuery(cardStatementTestUser.ID, http.MethodGet, "/card-statements/by-month", queryParams)
			err := cardStatementController.GetCardStatementsByMonth(c)
			
			// 検証
			if err != nil {
				t.Errorf("GetCardStatementsByMonth() error = %v", err)
				return
			}
			
			// エラーレスポンスが返されることを確認
			if rec.Code != http.StatusBadRequest {
				t.Errorf("GetCardStatementsByMonth() status code = %d, want %d", rec.Code, http.StatusBadRequest)
			}
		})
		
		t.Run("無効な年フォーマットの場合エラーになる", func(t *testing.T) {
			// 無効な年フォーマットのクエリパラメータ
			queryParams := map[string]string{
				"year":  "invalid",
				"month": "4",
			}
			
			// テスト実行
			_, c, rec := setupEchoWithJWTAndQuery(cardStatementTestUser.ID, http.MethodGet, "/card-statements/by-month", queryParams)
			err := cardStatementController.GetCardStatementsByMonth(c)
			
			// 検証
			if err != nil {
				t.Errorf("GetCardStatementsByMonth() error = %v", err)
				return
			}
			
			// エラーレスポンスが返されることを確認
			if rec.Code != http.StatusBadRequest {
				t.Errorf("GetCardStatementsByMonth() status code = %d, want %d", rec.Code, http.StatusBadRequest)
			}
		})
		
		t.Run("無効な月フォーマットの場合エラーになる", func(t *testing.T) {
			// 無効な月フォーマットのクエリパラメータ
			queryParams := map[string]string{
				"year":  "2023",
				"month": "invalid",
			}
			
			// テスト実行
			_, c, rec := setupEchoWithJWTAndQuery(cardStatementTestUser.ID, http.MethodGet, "/card-statements/by-month", queryParams)
			err := cardStatementController.GetCardStatementsByMonth(c)
			
			// 検証
			if err != nil {
				t.Errorf("GetCardStatementsByMonth() error = %v", err)
				return
			}
			
			// エラーレスポンスが返されることを確認
			if rec.Code != http.StatusBadRequest {
				t.Errorf("GetCardStatementsByMonth() status code = %d, want %d", rec.Code, http.StatusBadRequest)
			}
		})
		
		t.Run("範囲外の月（13以上）の場合エラーになる", func(t *testing.T) {
			// 範囲外の月のクエリパラメータ
			queryParams := map[string]string{
				"year":  "2023",
				"month": "13",
			}
			
			// テスト実行
			_, c, rec := setupEchoWithJWTAndQuery(cardStatementTestUser.ID, http.MethodGet, "/card-statements/by-month", queryParams)
			err := cardStatementController.GetCardStatementsByMonth(c)
			
			// 検証
			if err != nil {
				t.Errorf("GetCardStatementsByMonth() error = %v", err)
				return
			}
			
			// エラーレスポンスが返されることを確認
			if rec.Code != http.StatusInternalServerError {
				t.Errorf("GetCardStatementsByMonth() status code = %d, want %d", rec.Code, http.StatusInternalServerError)
			}
		})
	})
}