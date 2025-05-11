package card_statement_test

import (
	"monelog/model"
	"testing"
	"time"
)

// テスト用のカード明細を作成する関数
func createTestCardStatement(cardType string, description string, userId uint, year int, month int) *model.CardStatement {
	cardStatement := &model.CardStatement{
		Type:              "発生",
		StatementNo:       1,
		CardType:          cardType,
		Description:       description,
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
		UserId:            userId,
		Year:              year,
		Month:             month,
	}
	cardStatementDB.Create(cardStatement)
	return cardStatement
}

// カード明細の検証関数
func validateCardStatement(t *testing.T, cardStatement *model.CardStatement) {
	if cardStatement.ID == 0 {
		t.Error("CardStatement ID should not be zero")
	}
	if cardStatement.CreatedAt.IsZero() {
		t.Error("CreatedAt should not be zero")
	}
	if cardStatement.UpdatedAt.IsZero() {
		t.Error("UpdatedAt should not be zero")
	}
}

// 複数のテスト用カード明細を作成する関数
func createMultipleTestCardStatements(userId uint, count int, year int, month int) []model.CardStatement {
	statements := make([]model.CardStatement, count)
	
	for i := 0; i < count; i++ {
		statements[i] = model.CardStatement{
			Type:              "発生",
			StatementNo:       i + 1,
			CardType:          "楽天カード",
			Description:       "テスト明細 " + time.Now().String(),
			UseDate:           "2023/01/01",
			PaymentDate:       "2023/02/27",
			PaymentMonth:      "2023年02月",
			Amount:            1000 * (i + 1),
			TotalChargeAmount: 1000 * (i + 1),
			ChargeAmount:      0,
			RemainingBalance:  1000 * (i + 1),
			PaymentCount:      0,
			InstallmentCount:  1,
			AnnualRate:        0.0,
			MonthlyRate:       0.0,
			UserId:            userId,
			Year:              year,
			Month:             month,
		}
	}
	
	cardStatementDB.Create(&statements)
	return statements
}
