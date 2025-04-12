package card_statement_test

import (
	"encoding/json"
	"monelog/model"
	"testing"
	"time"
)

// テスト用カード明細を作成するヘルパー関数
func createTestCardStatement(userId uint, cardType string, description string) *model.CardStatement {
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
	}
	cardStatementDB.Create(cardStatement)
	return cardStatement
}

// 単一カード明細のレスポンスボディをパースするヘルパー関数
func parseCardStatementResponse(t *testing.T, responseBody []byte) model.CardStatementResponse {
	var response model.CardStatementResponse
	err := json.Unmarshal(responseBody, &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}
	return response
}

// 複数カード明細のレスポンスボディをパースするヘルパー関数
func parseCardStatementsResponse(t *testing.T, responseBody []byte) []model.CardStatementResponse {
	var response []model.CardStatementResponse
	err := json.Unmarshal(responseBody, &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}
	return response
}

// 現在時刻に基づいたユニークな説明文を生成するヘルパー関数
func generateUniqueDescription() string {
	return "テスト明細 " + time.Now().Format("20060102150405")
}
