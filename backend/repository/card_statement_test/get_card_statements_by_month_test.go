package card_statement_test

import (
	"monelog/model"
	"testing"
)

func TestGetCardStatementsByMonth(t *testing.T) {
	// テスト環境のセットアップ
	setupCardStatementTest(t) // 引数を追加

	defer cleanupCardStatementTest(t)

	// テストデータの作成
	testStatements := []model.CardStatement{
		{UserId: cardStatementTestUser, StatementNo: 1, PaymentCount: 1, Year: 2023, Month: 1}, // cardStatementTestUser.ID -> cardStatementTestUser
		{UserId: cardStatementTestUser, StatementNo: 1, PaymentCount: 2, Year: 2023, Month: 1}, // cardStatementTestUser.ID -> cardStatementTestUser
		{UserId: cardStatementTestUser, StatementNo: 2, PaymentCount: 1, Year: 2023, Month: 2}, // cardStatementTestUser.ID -> cardStatementTestUser
		{UserId: cardStatementOtherUser, StatementNo: 1, PaymentCount: 1, Year: 2023, Month: 1}, // cardStatementOtherUser.ID -> cardStatementOtherUser
	}

	// テストデータをデータベースに挿入
	for _, stmt := range testStatements {
		if err := cardStatementDB.Create(&stmt).Error; err != nil {
			t.Fatalf("Failed to create test data: %v", err)
		}
	}

	// テストケース
	testCases := []struct {
		name          string
		userId        uint
		year          int
		month         int
		expectedCount int
	}{
		{
			name:          "1月のカードステートメントの取得",
			userId:        cardStatementTestUser, // cardStatementTestUser.ID -> cardStatementTestUser
			year:          2023,
			month:         1,
			expectedCount: 2,
		},
		{
			name:          "2月のカードステートメントの取得",
			userId:        cardStatementTestUser, // cardStatementTestUser.ID -> cardStatementTestUser
			year:          2023,
			month:         2,
			expectedCount: 1,
		},
		{
			name:          "存在しない月のカードステートメントの取得",
			userId:        cardStatementTestUser, // cardStatementTestUser.ID -> cardStatementTestUser
			year:          2023,
			month:         3,
			expectedCount: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// テスト実行
			result, err := cardStatementRepo.GetCardStatementsByMonth(tc.userId, tc.year, tc.month)

			// 結果の検証
			if err != nil {
				t.Errorf("GetCardStatementsByMonth() error = %v", err)
				return
			}

			// 結果の数を確認
			if len(result) != tc.expectedCount {
				t.Errorf("Expected %d card statements, got %d", tc.expectedCount, len(result))
			}

			// 結果の内容を確認
			for _, stmt := range result {
				if stmt.UserId != tc.userId {
					t.Errorf("Expected user ID %d, got %d", tc.userId, stmt.UserId)
				}
				if stmt.Year != tc.year {
					t.Errorf("Expected year %d, got %d", tc.year, stmt.Year)
				}
				if stmt.Month != tc.month {
					t.Errorf("Expected month %d, got %d", tc.month, stmt.Month)
				}
			}
		})
	}
}