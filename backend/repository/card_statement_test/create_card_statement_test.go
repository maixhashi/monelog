package card_statement_test

import (
	"monelog/model"
	"testing"
)

func TestCreateCardStatement(t *testing.T) {
	// テスト環境のセットアップ
	setupCardStatementTest(t)
	defer cleanupCardStatementTest(t)

	// テストデータの準備
	testCases := []struct {
		name          string
		cardStatement model.CardStatement
		wantErr       bool
	}{
		{
			name: "正常なカードステートメントの作成",
			cardStatement: model.CardStatement{
				UserId:       cardStatementTestUser,
				StatementNo:  1,
				PaymentCount: 1,
				Year:         2023,
				Month:        1,
				Amount:       10000,
			},
			wantErr: false,
		},
		// 他のテストケースも追加可能
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// テスト実行
			err := cardStatementRepo.CreateCardStatement(&tc.cardStatement)

			// 結果の検証
			if (err != nil) != tc.wantErr {
				t.Errorf("CreateCardStatement() error = %v, wantErr %v", err, tc.wantErr)
				return
			}

			// 作成されたデータの検証
			if !tc.wantErr {
				var created model.CardStatement
				if err := cardStatementDB.First(&created, tc.cardStatement.ID).Error; err != nil {
					t.Errorf("Failed to retrieve created card statement: %v", err)
				}
			}
		})
	}
}
