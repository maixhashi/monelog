package card_statement_test

import (
	"monelog/model"
	"testing"
)

func TestCardStatementRepository_CreateCardStatement(t *testing.T) {
	setupCardStatementTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("カード明細を正常に作成できる", func(t *testing.T) {
			newStatement := &model.CardStatement{
				Type:              "発生",
				StatementNo:       1,
				CardType:          "楽天カード",
				Description:       "新規テスト明細",
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
				UserId:            cardStatementTestUser.ID,
				Year:              2023,
				Month:             1,
			}
			
			err := cardStatementRepo.CreateCardStatement(newStatement)
			
			if err != nil {
				t.Errorf("CreateCardStatement() error = %v", err)
			}
			
			validateCardStatement(t, newStatement)
			
			// データベースから取得して確認
			var savedStatement model.CardStatement
			result := cardStatementDB.First(&savedStatement, newStatement.ID)
			
			if result.Error != nil {
				t.Errorf("データベースからの取得に失敗: %v", result.Error)
			}
			
			if savedStatement.Description != "新規テスト明細" {
				t.Errorf("保存されたデータが一致しません: got %v, want %v", 
					savedStatement.Description, "新規テスト明細")
			}
		})
	})
}
