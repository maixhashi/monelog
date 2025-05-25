package parser

import (
	"testing"
)

func TestRakutenParser_Parse(t *testing.T) {
	// Setup test CSV data
	csvData := []byte("利用日,利用店名・商品名,利用者,支払方法,利用金額,支払手数料,支払総額,支払月4月,支払金額,5月繰越残高,5月以降支払金額\n2025/02/22,返済方法変更ＷＥＢ　102649ｱﾏｿﾞﾝ ﾏ-ｹﾂﾄ ﾌﾟ,本人,分割変更12回払い(1回目),12581,0,21360,4月,1171,2,43")

	// Create parser instance
	parser := NewRakutenParser()

	// Parse the CSV data
	summaries, err := parser.Parse(csvData)
	if err != nil {
		t.Fatalf("Failed to parse CSV data: %v", err)
	}

	// Check if we have the expected number of summaries
	// For a 12-month installment, we should have 1 "発生" record + 12 "分割" records = 13 records
	expectedCount := 13
	if len(summaries) != expectedCount {
		t.Errorf("Expected %d summaries, got %d", expectedCount, len(summaries))
	}

	// Verify the first record (発生 record)
	if len(summaries) > 0 {
		firstRecord := summaries[0]
		
		// Check record type
		if firstRecord.Type != "発生" {
			t.Errorf("Expected first record type to be '発生', got '%s'", firstRecord.Type)
		}
		
		// Check description
		expectedDescription := "返済方法変更ＷＥＢ　102649ｱﾏｿﾞﾝ ﾏ-ｹﾂﾄ ﾌﾟ"
		if firstRecord.Description != expectedDescription {
			t.Errorf("Expected description '%s', got '%s'", expectedDescription, firstRecord.Description)
		}
		
		// Check use date
		expectedUseDate := "2025/02/22"
		if firstRecord.UseDate != expectedUseDate {
			t.Errorf("Expected use date '%s', got '%s'", expectedUseDate, firstRecord.UseDate)
		}
		
		// Check amount
		expectedAmount := 12581
		if firstRecord.Amount != expectedAmount {
			t.Errorf("Expected amount %d, got %d", expectedAmount, firstRecord.Amount)
		}
		
		// Check total charge amount
		expectedTotalChargeAmount := 21360
		if firstRecord.TotalChargeAmount != expectedTotalChargeAmount {
			t.Errorf("Expected total charge amount %d, got %d", expectedTotalChargeAmount, firstRecord.TotalChargeAmount)
		}
		
		// Check installment count
		expectedInstallmentCount := 12
		if firstRecord.InstallmentCount != expectedInstallmentCount {
			t.Errorf("Expected installment count %d, got %d", expectedInstallmentCount, firstRecord.InstallmentCount)
		}
	}

	// Verify the second record (first 分割 record)
	if len(summaries) > 1 {
		secondRecord := summaries[1]
		
		// Check record type
		if secondRecord.Type != "分割" {
			t.Errorf("Expected second record type to be '分割', got '%s'", secondRecord.Type)
		}
		
		// Check payment count
		expectedPaymentCount := 1
		if secondRecord.PaymentCount != expectedPaymentCount {
			t.Errorf("Expected payment count %d, got %d", expectedPaymentCount, secondRecord.PaymentCount)
		}
		
		// Check charge amount for first payment
		expectedChargeAmount := 1171
		if secondRecord.ChargeAmount != expectedChargeAmount {
			t.Errorf("Expected charge amount %d, got %d", expectedChargeAmount, secondRecord.ChargeAmount)
		}
	}

	// Verify the last record (last 分割 record)
	if len(summaries) >= expectedCount {
		lastRecord := summaries[expectedCount-1]
		
		// Check record type
		if lastRecord.Type != "分割" {
			t.Errorf("Expected last record type to be '分割', got '%s'", lastRecord.Type)
		}
		
		// Check payment count
		expectedPaymentCount := 12
		if lastRecord.PaymentCount != expectedPaymentCount {
			t.Errorf("Expected payment count %d, got %d", expectedPaymentCount, lastRecord.PaymentCount)
		}
		
		// Check remaining balance for last payment
		expectedRemainingBalance := 0
		if lastRecord.RemainingBalance != expectedRemainingBalance {
			t.Errorf("Expected remaining balance %d, got %d", expectedRemainingBalance, lastRecord.RemainingBalance)
		}
	}

	// Test that all payment records have the correct statement number
	for i, summary := range summaries {
		expectedStatementNo := 1
		if summary.StatementNo != expectedStatementNo {
			t.Errorf("Summary at index %d: expected statement number %d, got %d", i, expectedStatementNo, summary.StatementNo)
		}
	}

	// Verify that the sum of all charge amounts equals the total charge amount
	var totalChargeSum int
	for _, summary := range summaries {
		if summary.Type == "分割" {
			totalChargeSum += summary.ChargeAmount
		}
	}
	
	if totalChargeSum != summaries[0].TotalChargeAmount {
		t.Errorf("Sum of charge amounts (%d) does not equal total charge amount (%d)", totalChargeSum, summaries[0].TotalChargeAmount)
	}
}

func TestRakutenParser_ParseWithEmptyData(t *testing.T) {
	// Test with empty data
	csvData := []byte("")
	parser := NewRakutenParser()
	
	summaries, err := parser.Parse(csvData)
	if err == nil {
		t.Error("Expected error when parsing empty data, got nil")
	}
	
	if len(summaries) != 0 {
		t.Errorf("Expected 0 summaries for empty data, got %d", len(summaries))
	}
}

func TestRakutenParser_ParseWithHeaderOnly(t *testing.T) {
	// Test with header only
	csvData := []byte("利用日,利用店名・商品名,利用者,支払方法,利用金額,支払手数料,支払総額,支払月4月,支払金額,5月繰越残高,5月以降支払金額")
	parser := NewRakutenParser()
	
	summaries, err := parser.Parse(csvData)
	if err != nil {
		t.Fatalf("Failed to parse CSV data with header only: %v", err)
	}
	
	if len(summaries) != 0 {
		t.Errorf("Expected 0 summaries for header-only data, got %d", len(summaries))
	}
}

func TestRakutenParser_ParseWithInvalidDate(t *testing.T) {
	// Test with invalid date format
	csvData := []byte("利用日,利用店名・商品名,利用者,支払方法,利用金額,支払手数料,支払総額,支払月4月,支払金額,5月繰越残高,5月以降支払金額\ninvalid-date,返済方法変更ＷＥＢ　102649ｱﾏｿﾞﾝ ﾏ-ｹﾂﾄ ﾌﾟ,本人,分割変更12回払い(1回目),12581,0,21360,4月,1171,2,43")
	parser := NewRakutenParser()
	
	summaries, err := parser.Parse(csvData)
	if err != nil {
		t.Fatalf("Failed to parse CSV data with invalid date: %v", err)
	}
	
	if len(summaries) != 0 {
		t.Errorf("Expected 0 summaries for data with invalid date, got %d", len(summaries))
	}
}

func TestRakutenParser_ParseWithOneTimePayment(t *testing.T) {
	// Test with one-time payment
	csvData := []byte("利用日,利用店名・商品名,利用者,支払方法,利用金額,支払手数料,支払総額,支払月4月,支払金額,5月繰越残高,5月以降支払金額\n2025/02/22,Amazon.co.jp,本人,一括払い,5000,0,5000,4月,5000,0,0")
	parser := NewRakutenParser()
	
	summaries, err := parser.Parse(csvData)
	if err != nil {
		t.Fatalf("Failed to parse CSV data with one-time payment: %v", err)
	}
	
	// For one-time payment, we should have 1 "発生" record + 1 "分割" record = 2 records
	expectedCount := 2
	if len(summaries) != expectedCount {
		t.Errorf("Expected %d summaries for one-time payment, got %d", expectedCount, len(summaries))
	}
	
	if len(summaries) >= 2 {
		// Check first record (発生)
		if summaries[0].Type != "発生" || summaries[0].InstallmentCount != 1 {
			t.Errorf("First record should be '発生' with installment count 1")
		}
		
		// Check second record (分割)
		if summaries[1].Type != "分割" || summaries[1].PaymentCount != 1 || summaries[1].ChargeAmount != 5000 {
			t.Errorf("Second record should be '分割' with payment count 1 and charge amount 5000")
		}
	}
}

func TestGetAnnualRate(t *testing.T) {
	tests := []struct {
		cardType         string
		installmentCount int
		expectedRate     float64
	}{
		{"楽天カード", 3, 0.1225},  // 12.25%
		{"楽天カード", 6, 0.1375},  // 13.75%
		{"楽天カード", 10, 0.145},  // 14.5%
		{"楽天カード", 12, 0.1475}, // 14.75%
		{"楽天カード", 24, 0.15},   // 15%
		{"楽天カード", 1, 0.15},    // 一括払いの場合
		{"不明なカード", 12, 0.15},  // 不明なカード種別の場合
	}
	
	for _, test := range tests {
		result := GetAnnualRate(test.cardType, test.installmentCount)
		if result != test.expectedRate {
			t.Errorf("GetAnnualRate(%s, %d) = %f, expected %f", 
				test.cardType, test.installmentCount, result, test.expectedRate)
		}
	}
}

func TestCalculateMonthlyRate(t *testing.T) {
	tests := []struct {
		annualRate    float64
		expectedRate  float64
	}{
		{12.0, 1.0},
		{0.0, 0.0},
		{15.0, 1.25},
	}
	
	for _, test := range tests {
		result := CalculateMonthlyRate(test.annualRate)
		if !almostEqual(result, test.expectedRate, 0.001) {
			t.Errorf("CalculateMonthlyRate(%f) = %f, expected %f", 
				test.annualRate, result, test.expectedRate)
		}
	}
}

// Helper function to compare float values with tolerance
func almostEqual(a, b, tolerance float64) bool {
	return (a-b) < tolerance && (b-a) < tolerance
}