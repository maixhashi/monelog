package card_statement_test

import (
	"monelog/model"
	"testing"
)

func createTestCardStatement(cardType string, description string, amount int, userId uint) *model.CardStatement {
	cardStatement := &model.CardStatement{
		Type:              "発生",
		StatementNo:       1,
		CardType:          cardType,
		Description:       description,
		UseDate:           "2023/01/01",
		PaymentDate:       "2023/02/27",
		PaymentMonth:      "2023年02月",
		Amount:            amount,
		TotalChargeAmount: amount,
		ChargeAmount:      0,
		RemainingBalance:  amount,
		PaymentCount:      0,
		InstallmentCount:  1,
		AnnualRate:        0.0,
		MonthlyRate:       0.0,
		UserId:            userId,
	}
	csDB.Create(cardStatement)
	return cardStatement
}

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
