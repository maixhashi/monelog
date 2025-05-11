package card_statement_test

import (
	"testing"
)

func TestCardStatementUsecase_GetCardStatementById(t *testing.T) {
	setupCardStatementUsecaseTest()
	
	// テストデータの作成
	cardStatement := createTestCardStatement(t, "テスト明細", testUser.ID, 2023, 1)
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("存在する明細を正しく取得する", func(t *testing.T) {
			t.Logf("明細ID %d を取得します", cardStatement.ID)
			
			response, err := cardStatementUsecase.GetCardStatementById(testUser.ID, cardStatement.ID)
			
			if err != nil {
				t.Errorf("GetCardStatementById() error = %v", err)
			}
			
			validateCardStatementResponse(t, response, cardStatement.ID, cardStatement.Description)
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("存在しないIDを指定した場合はエラーを返す", func(t *testing.T) {
			t.Logf("存在しないID %d を指定して明細を取得しようとします", nonExistentCardStatementID)
			
			_, err := cardStatementUsecase.GetCardStatementById(testUser.ID, nonExistentCardStatementID)
			
			if err == nil {
				t.Error("存在しないIDを指定したときにエラーが返されませんでした")
			} else {
				t.Logf("期待通りエラーが返されました: %v", err)
			}
		})
		
		t.Run("他のユーザーの明細は取得できない", func(t *testing.T) {
			// 他のユーザーの明細を作成
			otherUserCardStatement := createTestCardStatement(t, "他ユーザーの明細", otherUser.ID, 2023, 1)
			t.Logf("他ユーザーの明細(ID=%d)を別ユーザー(ID=%d)として取得しようとします", otherUserCardStatement.ID, testUser.ID)
			
			_, err := cardStatementUsecase.GetCardStatementById(testUser.ID, otherUserCardStatement.ID)
			
			if err == nil {
				t.Error("他のユーザーの明細を取得できてしまいました")
			} else {
				t.Logf("期待通りエラーが返されました: %v", err)
			}
		})
	})
}
