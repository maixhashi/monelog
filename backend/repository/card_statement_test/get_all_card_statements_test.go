package card_statement_test

import (
	"monelog/model"
	"testing"
)

func TestGetAllCardStatements(t *testing.T) {
	// テスト環境のセットアップ
	setupCardStatementTest(t) // 引数を追加

	defer cleanupCardStatementTest(t)

	// テストデータの作成
	testStatements := []model.CardStatement{
		{UserId: cardStatementTestUser, StatementNo: 1, PaymentCount: 1}, // cardStatementTestUser.ID -> cardStatementTestUser
		{UserId: cardStatementTestUser, StatementNo: 1, PaymentCount: 2}, // cardStatementTestUser.ID -> cardStatementTestUser
		{UserId: cardStatementOtherUser, StatementNo: 1, PaymentCount: 1}, // cardStatementOtherUser.ID -> cardStatementOtherUser
	}

	// テストデータをデータベースに挿入
	for _, stmt := range testStatements {
		if err := cardStatementDB.Create(&stmt).Error; err != nil {
			t.Fatalf("Failed to create test data: %v", err)
		}
	}

	// テスト実行
	result, err := cardStatementRepo.GetAllCardStatements(cardStatementTestUser) // cardStatementTestUser.ID -> cardStatementTestUser

	// 結果の検証
	if err != nil {
		t.Errorf("GetAllCardStatements() error = %v", err)
		return
	}

	// 結果の数を確認
	expectedCount := 2 // テストユーザーのステートメント数
	if len(result) != expectedCount {
		t.Errorf("Expected %d card statements, got %d", expectedCount, len(result))
	}

	// 結果の順序を確認
	if len(result) >= 2 {
		if result[0].StatementNo != 1 || result[0].PaymentCount != 1 {
			t.Errorf("Expected first result to be statement 1, payment 1, got statement %d, payment %d", 
				result[0].StatementNo, result[0].PaymentCount)
		}
		if result[1].StatementNo != 1 || result[1].PaymentCount != 2 {
			t.Errorf("Expected second result to be statement 1, payment 2, got statement %d, payment %d", 
				result[1].StatementNo, result[1].PaymentCount)
		}
	}
}
