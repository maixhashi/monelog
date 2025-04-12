package card_statement_test

import (
	"testing"
)

func TestCardStatementUsecase_GetCardStatementById(t *testing.T) {
	setupCardStatementUsecaseTest()
	
	// テストデータの作成
	cardStatement := createTestCardStatement(t, "テスト明細", testUser.ID)
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("存在するカード明細を正しく取得する", func(t *testing.T) {
			t.Logf("カード明細ID %d を取得します", cardStatement.ID)
			
			response, err := cardStatementUsecase.GetCardStatementById(testUser.ID, cardStatement.ID)
			
			if err != nil {
				t.Errorf("GetCardStatementById() error = %v", err)
			}
			
			validateCardStatementResponse(t, response, cardStatement.ID, cardStatement.Description)
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("存在しないIDを指定した場合はエラーを返す", func(t *testing.T) {
			t.Logf("存在しないID %d を指定してカード明細を取得しようとします", nonExistentCardStatementID)
			
			_, err := cardStatementUsecase.GetCardStatementById(testUser.ID, nonExistentCardStatementID)
			
			if err == nil {
				t.Error("存在しないIDを指定したときにエラーが返されませんでした")
			} else {
				t.Logf("期待通りエラーが返されました: %v", err)
			}
		})
		
		t.Run("他のユーザーのカード明細は取得できない", func(t *testing.T) {
			// 他のユーザーのカード明細を作成
			otherUserCardStatement := createTestCardStatement(t, "他ユーザーの明細", otherUser.ID)
			t.Logf("他ユーザーのカード明細(ID=%d)を別ユーザー(ID=%d)として取得しようとします", otherUserCardStatement.ID, testUser.ID)
			
			_, err := cardStatementUsecase.GetCardStatementById(testUser.ID, otherUserCardStatement.ID)
			
			if err == nil {
				t.Error("他のユーザーのカード明細を取得できてしまいました")
			} else {
				t.Logf("期待通りエラーが返されました: %v", err)
			}
		})
	})
}
