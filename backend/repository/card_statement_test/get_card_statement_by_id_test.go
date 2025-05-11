package card_statement_test

import (
	"testing"
)

func TestCardStatementRepository_GetCardStatementById(t *testing.T) {
	setupCardStatementTest()
	
	// テストデータの作成
	statement1 := createTestCardStatement("楽天カード", "テスト明細1", cardStatementTestUser.ID, 2023, 1)
	createTestCardStatement("MUFGカード", "テスト明細2", cardStatementOtherUser.ID, 2023, 1)
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("正しいIDとユーザーIDでカード明細を取得できる", func(t *testing.T) {
			result, err := cardStatementRepo.GetCardStatementById(cardStatementTestUser.ID, statement1.ID)
			
			if err != nil {
				t.Errorf("GetCardStatementById() error = %v", err)
			}
			
			if result.ID != statement1.ID {
				t.Errorf("GetCardStatementById() got ID = %v, want %v", result.ID, statement1.ID)
			}
			
			if result.Description != "テスト明細1" {
				t.Errorf("GetCardStatementById() got Description = %v, want %v", 
					result.Description, "テスト明細1")
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("存在しないIDの場合はエラーを返す", func(t *testing.T) {
			_, err := cardStatementRepo.GetCardStatementById(cardStatementTestUser.ID, nonExistentCardStatementID)
			
			if err == nil {
				t.Error("GetCardStatementById() expected error for non-existent ID, got nil")
			}
		})
		
		t.Run("他のユーザーのカード明細にアクセスできない", func(t *testing.T) {
			_, err := cardStatementRepo.GetCardStatementById(cardStatementTestUser.ID, statement1.ID+1)
			
			if err == nil {
				t.Error("GetCardStatementById() expected error for other user's statement, got nil")
			}
		})
	})
}
