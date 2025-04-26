package card_statement_test

import (
	"testing"
)

func TestCardStatementRepository_GetCardStatementById(t *testing.T) {
	setupCardStatementTest()
	
	// テスト用のカード明細を作成
	testCardStatement := createTestCardStatement("楽天カード", "Amazon", 1000, csTestUser.ID)
	otherUserCardStatement := createTestCardStatement("MUFG", "ヨドバシカメラ", 3000, csOtherUser.ID)
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("存在するカード明細IDを指定した場合、正しいカード明細を取得する", func(t *testing.T) {
			result, err := csRepo.GetCardStatementById(csTestUser.ID, testCardStatement.ID)
			
			if err != nil {
				t.Errorf("GetCardStatementById() error = %v", err)
			}
			
			if result.ID != testCardStatement.ID {
				t.Errorf("GetCardStatementById() got ID = %v, want %v", result.ID, testCardStatement.ID)
			}
			
			if result.Description != "Amazon" {
				t.Errorf("GetCardStatementById() got Description = %v, want %v", result.Description, "Amazon")
			}
			
			if result.Amount != 1000 {
				t.Errorf("GetCardStatementById() got Amount = %v, want %v", result.Amount, 1000)
			}
			
			validateCardStatement(t, &result)
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("存在しないカード明細IDを指定した場合、エラーを返す", func(t *testing.T) {
			_, err := csRepo.GetCardStatementById(csTestUser.ID, nonExistentCardStatementID)
			
			if err == nil {
				t.Error("GetCardStatementById() expected error, got nil")
			}
		})
		
		t.Run("他のユーザーのカード明細IDを指定した場合、エラーを返す", func(t *testing.T) {
			_, err := csRepo.GetCardStatementById(csTestUser.ID, otherUserCardStatement.ID)
			
			if err == nil {
				t.Error("GetCardStatementById() expected error, got nil")
			}
		})
	})
}
