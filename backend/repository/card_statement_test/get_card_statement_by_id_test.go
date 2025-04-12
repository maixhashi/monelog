package card_statement_test

import (
	"testing"
)

func TestCardStatementRepository_GetCardStatementById(t *testing.T) {
	setupCardStatementTest()
	
	cardStatement := createTestCardStatement("楽天カード", "Amazon", 1000, csTestUser.ID)
	otherCardStatement := createTestCardStatement("MUFG", "ヨドバシカメラ", 3000, csOtherUser.ID)
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("正しいユーザーIDとカード明細IDで明細を取得できる", func(t *testing.T) {
			result, err := csRepo.GetCardStatementById(csTestUser.ID, cardStatement.ID)
			
			if err != nil {
				t.Errorf("GetCardStatementById() error = %v", err)
			}
			
			if result.ID != cardStatement.ID {
				t.Errorf("GetCardStatementById() got ID = %v, want %v", result.ID, cardStatement.ID)
			}
			
			if result.Description != "Amazon" {
				t.Errorf("GetCardStatementById() got Description = %v, want Amazon", result.Description)
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("存在しないIDの場合はエラーを返す", func(t *testing.T) {
			_, err := csRepo.GetCardStatementById(csTestUser.ID, nonExistentCardStatementID)
			
			if err == nil {
				t.Error("GetCardStatementById() should return error for non-existent ID")
			}
		})
		
		t.Run("他のユーザーのカード明細にアクセスできない", func(t *testing.T) {
			_, err := csRepo.GetCardStatementById(csTestUser.ID, otherCardStatement.ID)
			
			if err == nil {
				t.Error("GetCardStatementById() should return error when accessing other user's card statement")
			}
		})
	})
}
