package card_statement_test

import (
	"net/http"
	"testing"
)

func TestCardStatementController_GetAllCardStatements(t *testing.T) {
	setupCardStatementControllerTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("ユーザーのカード明細を全て取得する", func(t *testing.T) {
			// テスト用カード明細の作成
			createTestCardStatement("楽天カード", "Amazon.co.jp", cardStatementTestUser.ID, 2023, 4)
			createTestCardStatement("楽天カード", "楽天市場", cardStatementTestUser.ID, 2023, 4)
			createTestCardStatement("楽天カード", "他ユーザーの明細", cardStatementOtherUser.ID, 2023, 4) // 別ユーザーの明細
			
			// テスト実行
			_, c, rec := setupEchoWithJWT(cardStatementTestUser.ID)
			err := cardStatementController.GetAllCardStatements(c)
			
			// 検証
			if err != nil {
				t.Errorf("GetAllCardStatements() error = %v", err)
			}
			
			if rec.Code != http.StatusOK {
				t.Errorf("GetAllCardStatements() status code = %d, want %d", rec.Code, http.StatusOK)
			}
			
			// レスポンスボディをパース
			response := parseCardStatementsResponse(t, rec.Body.Bytes())
			
			if len(response) != 2 {
				t.Errorf("GetAllCardStatements() returned %d card statements, want 2", len(response))
			}
			
			// カード明細の説明の確認
			descriptions := make(map[string]bool)
			for _, cs := range response {
				descriptions[cs.Description] = true
			}
			
			if !descriptions["Amazon.co.jp"] || !descriptions["楽天市場"] {
				t.Errorf("期待したカード明細が結果に含まれていません: %v", response)
			}
			
			// 他ユーザーの明細が含まれていないことを確認
			if descriptions["他ユーザーの明細"] {
				t.Errorf("他ユーザーの明細が結果に含まれています: %v", response)
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		// データベース接続エラーなどのケースをモックして追加可能
		// 現在の実装では直接テストできないため省略
	})
}
