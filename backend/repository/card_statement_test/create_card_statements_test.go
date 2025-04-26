package card_statement_test

import (
	"monelog/model"
	"testing"
)

func TestCardStatementRepository_CreateCardStatements(t *testing.T) {
	setupCardStatementTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("複数のカード明細を一括で作成できる", func(t *testing.T) {
			cardStatements := []model.CardStatement{
				{
					Type:              "発生",
					StatementNo:       1,
					CardType:          "楽天カード",
					Description:       "一括テスト1",
					UseDate:           "2023/01/01",
					PaymentDate:       "2023/02/27",
					PaymentMonth:      "2023年02月",
					Amount:            1000,
					TotalChargeAmount: 1000,
					ChargeAmount:      0,
					RemainingBalance:  1000,
					PaymentCount:      0,
					InstallmentCount:  1,
					UserId:            csTestUser.ID,
				},
				{
					Type:              "発生",
					StatementNo:       2,
					CardType:          "楽天カード",
					Description:       "一括テスト2",
					UseDate:           "2023/01/15",
					PaymentDate:       "2023/02/27",
					PaymentMonth:      "2023年02月",
					Amount:            2000,
					TotalChargeAmount: 2000,
					ChargeAmount:      0,
					RemainingBalance:  2000,
					PaymentCount:      0,
					InstallmentCount:  1,
					UserId:            csTestUser.ID,
				},
			}
			
			err := csRepo.CreateCardStatements(cardStatements)
			
			if err != nil {
				t.Errorf("CreateCardStatements() error = %v", err)
			}
			
			// 各カード明細のIDが設定されていることを確認
			for _, cs := range cardStatements {
				if cs.ID == 0 {
					t.Error("CreateCardStatements() failed to set ID")
				}
			}
			
			// データベースから取得して確認
			var count int64
			csDB.Model(&model.CardStatement{}).Where("description LIKE ?", "一括テスト%").Count(&count)
			
			if count != 2 {
				t.Errorf("CreateCardStatements() created %d records, want 2", count)
			}
			
			// 内容の確認
			var savedStatements []model.CardStatement
			csDB.Where("description LIKE ?", "一括テスト%").Order("amount").Find(&savedStatements)
			
			if len(savedStatements) != 2 {
				t.Errorf("Failed to retrieve created card statements")
				return
			}
			
			if savedStatements[0].Amount != 1000 || savedStatements[1].Amount != 2000 {
				t.Errorf("CreateCardStatements() saved incorrect amounts: %v, %v", 
					savedStatements[0].Amount, savedStatements[1].Amount)
			}
		})
	})
}
