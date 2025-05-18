package card_statement_test

import (
	"monelog/model"
	"testing"
)

func TestCreateCardStatements(t *testing.T) {
	// テスト環境のセットアップ
	setupCardStatementTest(t)
	defer cleanupCardStatementTest(t)

	// テストデータの準備
	testCases := []struct {
		name           string
		cardStatements []model.CardStatement
		wantErr        bool
	}{
		{
			name: "複数のカードステートメントの作成",
			cardStatements: []model.CardStatement{
				{
					UserId:       cardStatementTestUser,
					StatementNo:  1,
					PaymentCount: 1,
					Year:         2023,
					Month:        1,
					Amount:       10000,
				},
				{
					UserId:       cardStatementTestUser,
					StatementNo:  1,
					PaymentCount: 2,
					Year:         2023,
					Month:        2,
					Amount:       20000,
				},
			},
			wantErr: false,
		},
		{
			name:           "空の配列の場合",
			cardStatements: []model.CardStatement{},
			wantErr:        false,
		},
		// 他のテストケースも追加可能
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// テスト実行
			err := cardStatementRepo.CreateCardStatements(tc.cardStatements)

			// 結果の検証
			if (err != nil) != tc.wantErr {
				t.Errorf("CreateCardStatements() error = %v, wantErr %v", err, tc.wantErr)
				return
			}

			// 作成されたデータの検証
			if !tc.wantErr && len(tc.cardStatements) > 0 {
				var count int64
				cardStatementDB.Model(&model.CardStatement{}).Where("user_id = ?", cardStatementTestUser).Count(&count)
				if int(count) != len(tc.cardStatements) {
					t.Errorf("Expected %d card statements, got %d", len(tc.cardStatements), count)
				}
			}
		})
	}
}
