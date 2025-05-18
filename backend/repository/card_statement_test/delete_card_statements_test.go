package card_statement_test

import (
	"monelog/model"
	"testing"
)

func TestDeleteCardStatements(t *testing.T) {
	// テスト環境のセットアップ
	setupCardStatementTest(t) // 引数を追加

	defer cleanupCardStatementTest(t)

	// テストデータの作成
	testStatements := []model.CardStatement{
		{UserId: cardStatementTestUser, StatementNo: 1, PaymentCount: 1}, // cardStatementTestUser.ID -> cardStatementTestUser
		{UserId: cardStatementOtherUser, StatementNo: 1, PaymentCount: 1}, // cardStatementOtherUser を使用
	}

	// テストデータをデータベースに挿入
	for _, stmt := range testStatements {
		if err := cardStatementDB.Create(&stmt).Error; err != nil {
			t.Fatalf("Failed to create test data: %v", err)
		}
	}

	// テスト実行
	err := cardStatementRepo.DeleteCardStatements(cardStatementTestUser) // cardStatementTestUser.ID -> cardStatementTestUser

	// 結果の検証
	if err != nil {
		t.Errorf("DeleteCardStatements() error = %v", err)
		return
	}

	// 削除されたことを確認
	var count int64
	cardStatementDB.Model(&model.CardStatement{}).Where("user_id = ?", cardStatementTestUser).Count(&count) // cardStatementTestUser.ID -> cardStatementTestUser
	if count != 0 {
		t.Errorf("Expected 0 card statements for test user, got %d", count)
	}

	// 他のユーザーのデータは削除されていないことを確認
	cardStatementDB.Model(&model.CardStatement{}).Where("user_id = ?", cardStatementOtherUser).Count(&count)
	if count != 1 {
		t.Errorf("Expected 1 card statement for other user, got %d", count)
	}
}
