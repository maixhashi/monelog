package card_statement_test

import (
	"monelog/model"
	"testing"
)

func TestGetCardStatementById(t *testing.T) {
	// テスト環境のセットアップ
	setupCardStatementTest(t) // 引数を追加

	defer cleanupCardStatementTest(t)

	// テストデータの作成
	testStatement := model.CardStatement{UserId: cardStatementTestUser, StatementNo: 1, PaymentCount: 1} // cardStatementTestUser.ID -> cardStatementTestUser
	otherUserStatement := model.CardStatement{UserId: cardStatementOtherUser, StatementNo: 1, PaymentCount: 1} // cardStatementOtherUser.ID -> cardStatementOtherUser

	// テストデータをデータベースに挿入
	if err := cardStatementDB.Create(&testStatement).Error; err != nil {
		t.Fatalf("Failed to create test data: %v", err)
	}
	if err := cardStatementDB.Create(&otherUserStatement).Error; err != nil {
		t.Fatalf("Failed to create test data: %v", err)
	}

	// テストケース
	testCases := []struct {
		name           string
		userId         uint
		cardStatementId uint
		wantErr        bool
	}{
		{
			name:           "存在するカードステートメントの取得",
			userId:         cardStatementTestUser, // cardStatementTestUser.ID -> cardStatementTestUser
			cardStatementId: testStatement.ID,
			wantErr:        false,
		},
		{
			name:           "存在しないカードステートメントの取得",
			userId:         cardStatementTestUser, // cardStatementTestUser.ID -> cardStatementTestUser
			cardStatementId: 999, // 存在しないID
			wantErr:        true,
		},
		{
			name:           "他のユーザーのカードステートメントの取得",
			userId:         cardStatementTestUser, // cardStatementTestUser.ID -> cardStatementTestUser
			cardStatementId: otherUserStatement.ID,
			wantErr:        true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// テスト実行
			result, err := cardStatementRepo.GetCardStatementById(tc.userId, tc.cardStatementId)

			// 結果の検証
			if (err != nil) != tc.wantErr {
				t.Errorf("GetCardStatementById() error = %v, wantErr %v", err, tc.wantErr)
				return
			}

			// 正常系の場合、結果の内容を確認
			if !tc.wantErr {
				if result.ID != tc.cardStatementId {
					t.Errorf("Expected card statement ID %d, got %d", tc.cardStatementId, result.ID)
				}
				if result.UserId != tc.userId {
					t.Errorf("Expected user ID %d, got %d", tc.userId, result.UserId)
				}
			}
		})
	}
}
