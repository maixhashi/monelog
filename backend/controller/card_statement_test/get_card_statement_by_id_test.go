package card_statement_test

import (
	"fmt"
	"net/http"
	"testing"
)

func TestCardStatementController_GetCardStatementById(t *testing.T) {
	setupCardStatementControllerTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("指定されたIDのカード明細を取得する", func(t *testing.T) {
			// テスト用カード明細の作成
			cardStatement := createTestCardStatement("楽天カード", "Amazon.co.jp", cardStatementTestUser.ID, 2023, 4)
			
			// テスト実行
			_, c, rec := setupEchoWithCardStatementId(cardStatementTestUser.ID, cardStatement.ID, http.MethodGet, "/card-statements/"+fmt.Sprintf("%d", cardStatement.ID), "")
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
			
			// 取得したカード明細の確認
			if response.ID != cardStatement.ID {
				t.Errorf("GetCardStatementById() returned ID = %d, want %d", response.ID, cardStatement.ID)
			}
			
			if response.Description != "Amazon.co.jp" {
				t.Errorf("GetCardStatementById() returned Description = %s, want %s", response.Description, "Amazon.co.jp")
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("存在しないIDを指定した場合エラーになる", func(t *testing.T) {
			// テスト実行
			_, c, rec := setupEchoWithCardStatementId(cardStatementTestUser.ID, nonExistentCardStatementID, http.MethodGet, "/card-statements/"+fmt.Sprintf("%d", nonExistentCardStatementID), "")
			err := cardStatementController.GetCardStatementById(c)
			
			// 検証
			if err != nil {
				t.Errorf("GetCardStatementById() error = %v", err)
				return
			}
			
			// エラーレスポンスが返されることを確認
			if rec.Code != http.StatusInternalServerError {
				t.Errorf("GetCardStatementById() status code = %d, want %d", rec.Code, http.StatusInternalServerError)
			}
		})
		
		t.Run("他ユーザーのカード明細IDを指定した場合エラーになる", func(t *testing.T) {
			// 他ユーザーのカード明細を作成
			otherUserCardStatement := createTestCardStatement("楽天カード", "他ユーザーの明細", cardStatementOtherUser.ID, 2023, 4)
			
			// テスト実行
			_, c, rec := setupEchoWithCardStatementId(cardStatementTestUser.ID, otherUserCardStatement.ID, http.MethodGet, "/card-statements/"+fmt.Sprintf("%d", otherUserCardStatement.ID), "")
			err := cardStatementController.GetCardStatementById(c)
			
			// 検証
			if err != nil {
				t.Errorf("GetCardStatementById() error = %v", err)
				return
			}
			
			// エラーレスポンスが返されることを確認
			if rec.Code != http.StatusInternalServerError {
				t.Errorf("GetCardStatementById() status code = %d, want %d", rec.Code, http.StatusInternalServerError)
			}
		})
		
		t.Run("無効なIDフォーマットを指定した場合エラーになる", func(t *testing.T) {
			// 無効なIDフォーマット
			_, c, rec := setupEchoWithJWT(cardStatementTestUser.ID)
			c.SetParamNames("cardStatementId")
			c.SetParamValues("invalid-id")
			
			err := cardStatementController.GetCardStatementById(c)
			
			// 検証
			if err != nil {
				t.Errorf("GetCardStatementById() error = %v", err)
				return
			}
			
			// エラーレスポンスが返されることを確認
			if rec.Code != http.StatusBadRequest {
				t.Errorf("GetCardStatementById() status code = %d, want %d", rec.Code, http.StatusBadRequest)
			}
		})
	})
}
