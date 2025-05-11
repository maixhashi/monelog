package card_statement_test

import (
	"monelog/model"
	"testing"
)

func TestCardStatementRepository_DeleteCardStatements(t *testing.T) {
	setupCardStatementTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("指定したユーザーのカード明細のみを削除する", func(t *testing.T) {
			// テストデータの作成
			createMultipleTestCardStatements(cardStatementTestUser.ID, 3, 2023, 1)
			createMultipleTestCardStatements(cardStatementOtherUser.ID, 2, 2023, 1)
			
			// 削除前の確認
			var beforeCountUser1 int64
			cardStatementDB.Model(&model.CardStatement{}).Where("user_id = ?", cardStatementTestUser.ID).Count(&beforeCountUser1)
			
			var beforeCountUser2 int64
			cardStatementDB.Model(&model.CardStatement{}).Where("user_id = ?", cardStatementOtherUser.ID).Count(&beforeCountUser2)
			
			if beforeCountUser1 < 3 || beforeCountUser2 < 2 {
				t.Errorf("テストデータの作成に失敗: user1=%d, user2=%d", beforeCountUser1, beforeCountUser2)
			}
			
			// 削除実行
			err := cardStatementRepo.DeleteCardStatements(cardStatementTestUser.ID)
			
			if err != nil {
				t.Errorf("DeleteCardStatements() error = %v", err)
			}
			
			// 削除後の確認
			var afterCountUser1 int64
			cardStatementDB.Model(&model.CardStatement{}).Where("user_id = ?", cardStatementTestUser.ID).Count(&afterCountUser1)
			
			var afterCountUser2 int64
			cardStatementDB.Model(&model.CardStatement{}).Where("user_id = ?", cardStatementOtherUser.ID).Count(&afterCountUser2)
			
			if afterCountUser1 != 0 {
				t.Errorf("ユーザー1の明細が削除されていません: got %d, want 0", afterCountUser1)
			}
			
			if afterCountUser2 != beforeCountUser2 {
				t.Errorf("ユーザー2の明細が誤って削除されています: before=%d, after=%d", 
					beforeCountUser2, afterCountUser2)
			}
		})
		
		t.Run("存在しないユーザーIDの場合もエラーなく処理される", func(t *testing.T) {
			nonExistentUserId := uint(9999)
			
			err := cardStatementRepo.DeleteCardStatements(nonExistentUserId)
			
			if err != nil {
				t.Errorf("DeleteCardStatements() with non-existent user ID error = %v", err)
			}
		})
	})
}
