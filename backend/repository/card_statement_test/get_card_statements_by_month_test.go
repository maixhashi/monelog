package card_statement_test

import (
	"monelog/model"
	"testing"
)

func TestCardStatementRepository_GetCardStatementsByMonth(t *testing.T) {
	setupCardStatementTest()
	
	// 異なる支払月のカード明細を作成
	januaryStatement := &model.CardStatement{
		Type:              "発生",
		StatementNo:       1,
		CardType:          "楽天カード",
		Description:       "1月支払い",
		UseDate:           "2022/12/01",
		PaymentDate:       "2023/01/27", // 1月支払い
		PaymentMonth:      "2023年01月",
		Amount:            1000,
		TotalChargeAmount: 1000,
		ChargeAmount:      0,
		RemainingBalance:  1000,
		PaymentCount:      0,
		InstallmentCount:  1,
		UserId:            csTestUser.ID,
	}
	csDB.Create(januaryStatement)
	
	februaryStatement := &model.CardStatement{
		Type:              "発生",
		StatementNo:       2,
		CardType:          "楽天カード",
		Description:       "2月支払い",
		UseDate:           "2023/01/01",
		PaymentDate:       "2023/02/27", // 2月支払い
		PaymentMonth:      "2023年02月",
		Amount:            2000,
		TotalChargeAmount: 2000,
		ChargeAmount:      0,
		RemainingBalance:  2000,
		PaymentCount:      0,
		InstallmentCount:  1,
		UserId:            csTestUser.ID,
	}
	csDB.Create(februaryStatement)
	
	// 他のユーザーの2月支払い
	otherUserFebruaryStatement := &model.CardStatement{
		Type:              "発生",
		StatementNo:       3,
		CardType:          "MUFG",
		Description:       "他ユーザーの2月支払い",
		UseDate:           "2023/01/15",
		PaymentDate:       "2023/02/15", // 2月支払い
		PaymentMonth:      "2023年02月",
		Amount:            3000,
		TotalChargeAmount: 3000,
		ChargeAmount:      0,
		RemainingBalance:  3000,
		PaymentCount:      0,
		InstallmentCount:  1,
		UserId:            csOtherUser.ID,
	}
	csDB.Create(otherUserFebruaryStatement)
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("指定した年月の支払いに関するカード明細のみを取得する", func(t *testing.T) {
			// 2023年2月の支払いを取得
			result, err := csRepo.GetCardStatementsByMonth(csTestUser.ID, 2023, 2)
			
			if err != nil {
				t.Errorf("GetCardStatementsByMonth() error = %v", err)
			}
			
			if len(result) != 1 {
				t.Errorf("GetCardStatementsByMonth() got %d card statements, want 1", len(result))
			}
			
			if result[0].Description != "2月支払い" {
				t.Errorf("GetCardStatementsByMonth() got Description = %v, want %v", result[0].Description, "2月支払い")
			}
		})
		
		t.Run("存在しない年月を指定した場合、空の配列を返す", func(t *testing.T) {
			// 2023年3月の支払いを取得（データなし）
			result, err := csRepo.GetCardStatementsByMonth(csTestUser.ID, 2023, 3)
			
			if err != nil {
				t.Errorf("GetCardStatementsByMonth() error = %v", err)
			}
			
			if len(result) != 0 {
				t.Errorf("GetCardStatementsByMonth() got %d card statements, want 0", len(result))
			}
		})
	})
}