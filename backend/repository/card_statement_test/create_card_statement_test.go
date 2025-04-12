package card_statement_test

import (
	"monelog/model"
	"testing"
)

func TestCardStatementRepository_CreateCardStatement(t *testing.T) {
	setupCardStatementTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("カード明細を作成できる", func(t *testing.T) {
			cardStatement := &model.CardStatement{
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
			}
			
			err := csRepo.CreateCardStatement(cardStatement)
			
			if err != nil {
				t.Errorf("CreateCardStatement() error = %v", err)
			}
			
			validateCardStatement(t, cardStatement)
			
			// データベースから取得して確認
			var savedCardStatement model.CardStatement
			csDB.First(&savedCardStatement, cardStatement.ID)
			
			if savedCardStatement.ID != cardStatement.ID {
				t.Errorf("CreateCardStatement() failed to save card statement")
			}
			
			if savedCardStatement.Description != "Amazon" {
				t.Errorf("CreateCardStatement() got Description = %v, want Amazon", savedCardStatement.Description)
			}
		})
	})
}
