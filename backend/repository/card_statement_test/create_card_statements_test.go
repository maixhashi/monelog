package card_statement_test

import (
	"monelog/model"
	"testing"
)

func TestCardStatementRepository_CreateCardStatements(t *testing.T) {
	setupCardStatementTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("複数のカード明細を一度に作成できる", func(t *testing.T) {
			statements := []model.CardStatement{
				{
					Type:              "発生",
					StatementNo:       1,
					CardType:          "楽天カード",
					Description:       "一括テスト明細1",
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
					UserId:            cardStatementTestUser.ID,
					Year:              2023,
					Month:             1,
				},
				{
					Type:              "発生",
					StatementNo:       2,
					CardType:          "楽天カード",
					Description:       "一括テスト明細2",
					UseDate:           "2023/01/02",
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
					UserId:            cardStatementTestUser.ID,
					Year:              2023,
					Month:             1,
				},
			}
			
			err := cardStatementRepo.CreateCardStatements(statements)
			
			if err != nil {
				t.Errorf("CreateCardStatements() error = %v", err)
			}
			
			// データベースから取得して確認
			var savedStatements []model.CardStatement
			result := cardStatementDB.Where("description LIKE ?", "一括テスト明細%").Find(&savedStatements)
			
			if result.Error != nil {
				t.Errorf("データベースからの取得に失敗: %v", result.Error)
			}
			
			if len(savedStatements) != 2 {
				t.Errorf("保存された明細数が一致しません: got %d, want 2", len(savedStatements))
			}
			
			descriptions := make(map[string]bool)
			for _, statement := range savedStatements {
				descriptions[statement.Description] = true
				
				if statement.ID == 0 {
					t.Error("保存された明細のIDが設定されていません")
				}
			}
			
			if !descriptions["一括テスト明細1"] || !descriptions["一括テスト明細2"] {
				t.Errorf("期待した明細が保存されていません: %v", savedStatements)
			}
		})
		
		t.Run("空の配列の場合はエラーなく処理される", func(t *testing.T) {
			emptyStatements := []model.CardStatement{}
			
			err := cardStatementRepo.CreateCardStatements(emptyStatements)
			
			if err != nil {
				t.Errorf("CreateCardStatements() with empty array error = %v", err)
			}
		})
	})
}
