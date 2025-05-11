package card_statement_test

import (
	"testing"
)

func TestCardStatementRepository_GetAllCardStatements(t *testing.T) {
	setupCardStatementTest()
	
	// テストデータの作成
	createTestCardStatement("楽天カード", "テスト明細1", cardStatementTestUser.ID, 2023, 1)
	createTestCardStatement("楽天カード", "テスト明細2", cardStatementTestUser.ID, 2023, 1)
	createTestCardStatement("MUFGカード", "テスト明細3", cardStatementOtherUser.ID, 2023, 1)
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("正しいユーザーIDのカード明細のみを取得する", func(t *testing.T) {
			result, err := cardStatementRepo.GetAllCardStatements(cardStatementTestUser.ID)
			
			if err != nil {
				t.Errorf("GetAllCardStatements() error = %v", err)
			}
			
			if len(result) != 2 {
				t.Errorf("GetAllCardStatements() got %d statements, want 2", len(result))
			}
			
			descriptions := make(map[string]bool)
			for _, statement := range result {
				descriptions[statement.Description] = true
				
				if statement.UserId != cardStatementTestUser.ID {
					t.Errorf("取得したカード明細のユーザーIDが一致しません: got %d, want %d", 
						statement.UserId, cardStatementTestUser.ID)
				}
			}
			
			if !descriptions["テスト明細1"] || !descriptions["テスト明細2"] {
				t.Errorf("期待したカード明細が結果に含まれていません: %v", result)
			}
		})
	})
}
