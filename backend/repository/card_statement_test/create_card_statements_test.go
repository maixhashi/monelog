package card_statement_test

import (
	"monelog/model"
	"testing"
)

func TestCardStatementRepository_CreateCardStatements(t *testing.T) {
	setupCardStatementTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("複数のカード明細を一括作成できる", func(t *testing.T) {
			cardStatements := []model.CardStatement{
				{
					Type:              "発生",
					StatementNo:       1,
					CardType:          "楽天カード",
					Description:       "Amazon",
					UseDate:           "2023/01/01",
					PaymentDate:       "2023/02/27",
					PaymentMonth:      "2023年02月",
					Amount:            1000,
					TotalChargeAmount: 1000,
					ChargeAmount:      0,
					RemainingBalance:  1000,
					PaymentCount:      0,
					InstallmentCount:  1,
					AnnualRate:        0.0,
					MonthlyRate:       0.0,
					UserId:            csTestUser.ID,
				},
				{
					Type:              "発生",
					StatementNo:       2,
					CardType:          "楽天カード",
					Description:       "楽天市場",
					UseDate:           "2023/01/05",
					PaymentDate:       "2023/02/27",
					PaymentMonth:      "2023年02月",
					Amount:            2000,
					TotalChargeAmount: 2000,
					ChargeAmount:      0,
					RemainingBalance:  2000,
					PaymentCount:      0,
					InstallmentCount:  1,
					AnnualRate:        0.0,
					MonthlyRate:       0.0,
					UserId:            csTestUser.ID,
				},
			}
			
			err := csRepo.CreateCardStatements(cardStatements)
			
			if err != nil {
				t.Errorf("CreateCardStatements() error = %v", err)
			}
			
			// データベースから取得して確認
			var savedCardStatements []model.CardStatement
			csDB.Where("user_id = ?", csTestUser.ID).Find(&savedCardStatements)
			
			if len(savedCardStatements) != 2 {
				t.Errorf("CreateCardStatements() got %d card statements, want 2", len(savedCardStatements))
			}
			
			descriptions := make(map[string]bool)
			for _, cs := range savedCardStatements {
				descriptions[cs.Description] = true
				validateCardStatement(t, &cs)
			}
			
			if !descriptions["Amazon"] || !descriptions["楽天市場"] {
				t.Errorf("期待したカード明細が結果に含まれていません: %v", savedCardStatements)
			}
		})
	})
}
