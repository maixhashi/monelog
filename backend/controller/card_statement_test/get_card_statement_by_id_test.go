package card_statement_test

import (
	"net/http"
	"testing"
)

func TestCardStatementController_GetCardStatementById(t *testing.T) {
	setupCardStatementControllerTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("指定IDのカード明細を取得する", func(t *testing.T) {
			// テスト用カード明細の作成
			cardStatement := createTestCardStatement(cardStatementTestUser.ID, "楽天カード", "Amazon.co.jp")
			
			// テスト実行
			_, c, rec := setupEchoWithCardStatementId(cardStatementTestUser.ID, cardStatement.ID, http.MethodGet, "/", "")
			err := cardStatementController.GetCardStatementById(c)
			
			// 検証
			if err != nil {
				t.Errorf("GetCardStatementById() error = %v", err)
			}
			
			if rec.Code != http.StatusOK {
				t.Errorf("GetCardStatementById() status code = %d, want %d", rec.Code, http.StatusOK)
			}
			
			// レスポンスボディをパース
			response := parseCardStatementResponse(t, rec.Body.Bytes())
			
			if response.ID != cardStatement.ID {
				t.Errorf("GetCardStatementById() returned ID = %d, want %d", response.ID, cardStatement.ID)
			}
			
			if response.Description != "Amazon.co.jp" {
				t.Errorf("GetCardStatementById() returned Description = %s, want %s", response.Description, "Amazon.co.jp")
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("存在しないIDを指定した場合", func(t *testing.T) {
			// テスト実行
			_, c, rec := setupEchoWithCardStatementId(cardStatementTestUser.ID, nonExistentCardStatementID, http.MethodGet, "/", "")
			err := cardStatementController.GetCardStatementById(c)
			
			// エラーは返さないが、内部でエラーハンドリングしてステータスコードを返す
			if err != nil {
				t.Errorf("GetCardStatementById() unexpected error = %v", err)
			}
			
			if rec.Code != http.StatusInternalServerError {
				t.Errorf("GetCardStatementById() status code = %d, want %d", rec.Code, http.StatusInternalServerError)
			}
		})
		
		t.Run("無効なIDフォーマットを指定した場合", func(t *testing.T) {
			// テスト実行
			_, c, rec := setupEchoWithJWTAndBody(cardStatementTestUser.ID, http.MethodGet, "/", "")
			c.SetParamNames("cardStatementId")
			c.SetParamValues("invalid-id")
			
			err := cardStatementController.GetCardStatementById(c)
			
			// エラーは返さないが、内部でエラーハンドリングしてステータスコードを返す
			if err != nil {
				t.Errorf("GetCardStatementById() unexpected error = %v", err)
			}
			
			if rec.Code != http.StatusBadRequest {
				t.Errorf("GetCardStatementById() status code = %d, want %d", rec.Code, http.StatusBadRequest)
			}
		})
		
		t.Run("他ユーザーのカード明細IDを指定した場合", func(t *testing.T) {
			// 他ユーザーのカード明細を作成
			otherUserCardStatement := createTestCardStatement(cardStatementOtherUser.ID, "エポスカード", "ヨドバシカメラ")
			
			// テスト実行
			_, c, rec := setupEchoWithCardStatementId(cardStatementTestUser.ID, otherUserCardStatement.ID, http.MethodGet, "/", "")
			err := cardStatementController.GetCardStatementById(c)
			
			// エラーは返さないが、内部でエラーハンドリングしてステータスコードを返す
			if err != nil {
				t.Errorf("GetCardStatementById() unexpected error = %v", err)
			}
			
			if rec.Code != http.StatusInternalServerError {
				t.Errorf("GetCardStatementById() status code = %d, want %d", rec.Code, http.StatusInternalServerError)
			}
		})
	})
}
