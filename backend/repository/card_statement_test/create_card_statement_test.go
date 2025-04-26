package card_statement_test

import (
	"monelog/model"
	"testing"
)

func TestCardStatementRepository_CreateCardStatement(t *testing.T) {
	setupCardStatementTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("新しいカード明細を作成できる", func(t *testing.T) {
			newCardStatement := &model.CardStatement{
				Type:              "発生",
				StatementNo:       1,
				CardType:          "楽天カード",
				Description:       "テスト明細",
				UseDate:           "2023/01/01",
				PaymentDate:       "2023/02/27",
				PaymentMonth:      "2023年02月",
				Amount:            5000,
				TotalChargeAmount: 5000,
				ChargeAmount:      0,
				RemainingBalance:  5000,
				PaymentCount:      0,
				InstallmentCount:  1,
				AnnualRate:        0.0,
				MonthlyRate:       0.0,
				UserId:            csTestUser.ID,
			}
			
			err := csRepo.CreateCardStatement(newCardStatement)
			
			if err != nil {
				t.Errorf("CreateCardStatement() error = %v", err)
			}
			
			validateCardStatement(t, newCardStatement)
			
			// データベースから取得して確認
			var savedCardStatement model.CardStatement
			result := csDB.First(&savedCardStatement, newCardStatement.ID)
			
			if result.Error != nil {
				t.Errorf("Failed to retrieve created card statement: %v", result.Error)
			}
			
			if savedCardStatement.Description != "テスト明細" {
				t.Errorf("CreateCardStatement() saved Description = %v, want %v", savedCardStatement.Description, "テスト明細")
			}
			
			if savedCardStatement.Amount != 5000 {
				t.Errorf("CreateCardStatement() saved Amount = %v, want %v", savedCardStatement.Amount, 5000)
			}
		})
	})
}
