package card_statement_test

import (
	"monelog/model"
	"testing"
)

func TestCardStatementUsecase_GetAllCardStatements(t *testing.T) {
	setupCardStatementUsecaseTest()
	
	// テストデータの作成
	cardStatements := []model.CardStatement{
		{
			Type:              "発生",
			StatementNo:       1,
			CardType:          "テストカード",
			Description:       "テスト明細1",
			UseDate:           "2023/01/01",
			PaymentDate:       "2023/02/27",
			PaymentMonth:      "2023年02月",
			Amount:            10000,
			TotalChargeAmount: 10000,
			ChargeAmount:      0,
			RemainingBalance:  10000,
			PaymentCount:      0,
			InstallmentCount:  1,
			AnnualRate:        0.0,
			MonthlyRate:       0.0,
			UserId:            testUser.ID,
		},
		{
			Type:              "発生",
			StatementNo:       2,
			CardType:          "テストカード",
			Description:       "テスト明細2",
			UseDate:           "2023/01/02",
			PaymentDate:       "2023/02/27",
			PaymentMonth:      "2023年02月",
			Amount:            20000,
			TotalChargeAmount: 20000,
			ChargeAmount:      0,
			RemainingBalance:  20000,
			PaymentCount:      0,
			InstallmentCount:  1,
			AnnualRate:        0.0,
			MonthlyRate:       0.0,
			UserId:            testUser.ID,
		},
		{
			Type:              "発生",
			StatementNo:       1,
			CardType:          "テストカード",
			Description:       "他ユーザー明細",
			UseDate:           "2023/01/03",
			PaymentDate:       "2023/02/27",
			PaymentMonth:      "2023年02月",
			Amount:            30000,
			TotalChargeAmount: 30000,
			ChargeAmount:      0,
			RemainingBalance:  30000,
			PaymentCount:      0,
			InstallmentCount:  1,
			AnnualRate:        0.0,
			MonthlyRate:       0.0,
			UserId:            otherUser.ID,
		},
	}
	
	for _, cardStatement := range cardStatements {
		db.Create(&cardStatement)
	}
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("正しいユーザーIDのカード明細のみを取得する", func(t *testing.T) {
			t.Logf("ユーザーID %d のカード明細を取得します", testUser.ID)
			
			cardStatementResponses, err := cardStatementUsecase.GetAllCardStatements(testUser.ID)
			
			if err != nil {
				t.Errorf("GetAllCardStatements() error = %v", err)
			}
			
			if len(cardStatementResponses) != 2 {
				t.Errorf("GetAllCardStatements() got %d card statements, want 2", len(cardStatementResponses))
			}
			
			// 明細説明の確認
			descriptions := make(map[string]bool)
			for _, cardStatement := range cardStatementResponses {
				descriptions[cardStatement.Description] = true
				t.Logf("取得したカード明細: ID=%d, Description=%s", cardStatement.ID, cardStatement.Description)
			}
			
			if !descriptions["テスト明細1"] || !descriptions["テスト明細2"] {
				t.Errorf("期待したカード明細が結果に含まれていません: %v", cardStatementResponses)
			}
			
			// レスポンス形式の検証
			for _, cardStatement := range cardStatementResponses {
				if cardStatement.ID == 0 || cardStatement.Description == "" || cardStatement.CreatedAt.IsZero() || cardStatement.UpdatedAt.IsZero() {
					t.Errorf("GetAllCardStatements() returned invalid card statement: %+v", cardStatement)
				}
			}
		})
	})
}
