package card_statement_test

import (
	"monelog/model"
	"testing"
)

func TestCardStatementRepository_DeleteCardStatements(t *testing.T) {
	setupCardStatementTest()
	
	// テスト用のカード明細を作成
	cardStatements := []model.CardStatement{
		{Type: "発生", CardType: "楽天カード", Description: "Amazon", Amount: 1000, UserId: csTestUser.ID},
		{Type: "発生", CardType: "楽天カード", Description: "楽天市場", Amount: 2000, UserId: csTestUser.ID},
		{Type: "発生", CardType: "MUFG", Description: "ヨドバシカメラ", Amount: 3000, UserId: csOtherUser.ID},
	}
	
	for _, cs := range cardStatements {
		csDB.Create(&cs)
	}
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("指定したユーザーIDのカード明細のみを削除する", func(t *testing.T) {
			// 削除前の確認
			var beforeCount int64
			csDB.Model(&model.CardStatement{}).Where("user_id = ?", csTestUser.ID).Count(&beforeCount)
			if beforeCount != 2 {
				t.Errorf("テスト前のカード明細数が想定と異なります: got %d, want 2", beforeCount)
			}
			
			err := csRepo.DeleteCardStatements(csTestUser.ID)
			
			if err != nil {
				t.Errorf("DeleteCardStatements() error = %v", err)
			}
			
			// 削除後の確認
			var afterCount int64
			csDB.Model(&model.CardStatement{}).Where("user_id = ?", csTestUser.ID).Count(&afterCount)
			if afterCount != 0 {
				t.Errorf("DeleteCardStatements() 削除後のカード明細数が想定と異なります: got %d, want 0", afterCount)
			}
			
			// 他のユーザーのカード明細は削除されていないことを確認
			var otherUserCount int64
			csDB.Model(&model.CardStatement{}).Where("user_id = ?", csOtherUser.ID).Count(&otherUserCount)
			if otherUserCount != 1 {
				t.Errorf("DeleteCardStatements() 他のユーザーのカード明細が削除されています: got %d, want 1", otherUserCount)
			}
		})
	})
}
