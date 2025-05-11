package card_statement_test

import (
	"monelog/model"
	"testing"
)

// バリデーションエラーの有無を確認するヘルパー関数
func assertValidationResult(t *testing.T, err error, expectError bool, testName string) {
	if (err != nil) != expectError {
		if expectError {
			t.Errorf("%s: エラーが期待されましたが、エラーはありませんでした", testName)
		} else {
			t.Errorf("%s: エラーは期待されていませんでしたが、エラーが発生しました: %v", testName, err)
		}
	} else if err != nil {
		t.Logf("%s: 期待通りのエラーが発生しました: %v", testName, err)
	} else {
		t.Logf("%s: 期待通りエラーは発生しませんでした", testName)
	}
}

// テスト用のカード明細サマリーを作成するヘルパー関数
func createValidCardStatementSummaries() []model.CardStatementSummary {
	return []model.CardStatementSummary{
		{
			Type:              "発生",
			StatementNo:       1,
			CardType:          "楽天カード",
			Description:       "Amazon.co.jp",
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
		},
	}
}