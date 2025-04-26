package validator

import (
	"monelog/model"
	"monelog/testutils"
	"testing"
)

func TestCardStatementValidate(t *testing.T) {
	// テスト用DBの設定
	db := testutils.SetupTestDB()
	defer testutils.CleanupTestDB(db)
	
	// テストユーザーの作成
	user := testutils.CreateTestUser(db)
	
	validator := NewCardStatementValidator()

	testCases := []struct {
		name     string
		request  model.CardStatementRequest
		hasError bool
	}{
		{
			name: "Valid card statement with rakuten card type",
			request: model.CardStatementRequest{
				CardType: "rakuten",
				UserId:   user.ID,
			},
			hasError: false,
		},
		{
			name: "Valid card statement with mufg card type",
			request: model.CardStatementRequest{
				CardType: "mufg",
				UserId:   user.ID,
			},
			hasError: false,
		},
		{
			name: "Valid card statement with epos card type",
			request: model.CardStatementRequest{
				CardType: "epos",
				UserId:   user.ID,
			},
			hasError: false,
		},
		{
			name: "Empty card type",
			request: model.CardStatementRequest{
				CardType: "",
				UserId:   user.ID,
			},
			hasError: true,
		},
		{
			name: "Invalid card type",
			request: model.CardStatementRequest{
				CardType: "invalid_card",
				UserId:   user.ID,
			},
			hasError: true,
		},
		{
			name: "Zero user ID with valid card type",
			request: model.CardStatementRequest{
				CardType: "rakuten",
				UserId:   0,
			},
			hasError: false, // UserIDはバリデーションしていないので、エラーにならないはず
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validator.ValidateCardStatementRequest(tc.request)
			if (err != nil) != tc.hasError {
				t.Errorf("ValidateCardStatementRequest() error = %v, want error: %v", err, tc.hasError)
			}
		})
	}
}

// 追加の別テストケース - テスト用DBを使ったバリデーション
func TestCardStatementValidateWithDB(t *testing.T) {
	// テスト用DBの設定
	db := testutils.SetupTestDB()
	defer testutils.CleanupTestDB(db)
	
	// 異なるテストユーザーを作成
	user1 := testutils.CreateTestUser(db)
	user2 := testutils.CreateOtherUser(db)
	
	validator := NewCardStatementValidator()

	// ユーザー1のカード明細リクエストを作成
	request1 := model.CardStatementRequest{
		CardType: "rakuten",
		UserId:   user1.ID,
	}
	
	// ユーザー2のカード明細リクエストを作成
	request2 := model.CardStatementRequest{
		CardType: "mufg",
		UserId:   user2.ID,
	}
	
	// バリデーションのテスト
	err1 := validator.ValidateCardStatementRequest(request1)
	if err1 != nil {
		t.Errorf("ValidateCardStatementRequest() for user1 should not return error, got: %v", err1)
	}
	
	err2 := validator.ValidateCardStatementRequest(request2)
	if err2 != nil {
		t.Errorf("ValidateCardStatementRequest() for user2 should not return error, got: %v", err2)
	}
	
	// 無効なカード種類のテスト
	invalidRequest := model.CardStatementRequest{
		CardType: "visa", // サポートされていないカード種類
		UserId:   user1.ID,
	}
	
	err3 := validator.ValidateCardStatementRequest(invalidRequest)
	if err3 == nil {
		t.Errorf("ValidateCardStatementRequest() with invalid card type should return error")
	}
}

// プレビューリクエストのバリデーションテスト
func TestValidateCardStatementPreviewRequest(t *testing.T) {
	// テスト用DBの設定
	db := testutils.SetupTestDB()
	defer testutils.CleanupTestDB(db)
	
	// テストユーザーの作成
	user := testutils.CreateTestUser(db)
	
	validator := NewCardStatementValidator()

	testCases := []struct {
		name     string
		request  model.CardStatementPreviewRequest
		hasError bool
	}{
		{
			name: "Valid preview request with rakuten card type",
			request: model.CardStatementPreviewRequest{
				CardType: "rakuten",
				UserId:   user.ID,
			},
			hasError: false,
		},
		{
			name: "Valid preview request with mufg card type",
			request: model.CardStatementPreviewRequest{
				CardType: "mufg",
				UserId:   user.ID,
			},
			hasError: false,
		},
		{
			name: "Valid preview request with epos card type",
			request: model.CardStatementPreviewRequest{
				CardType: "epos",
				UserId:   user.ID,
			},
			hasError: false,
		},
		{
			name: "Empty card type",
			request: model.CardStatementPreviewRequest{
				CardType: "",
				UserId:   user.ID,
			},
			hasError: true,
		},
		{
			name: "Invalid card type",
			request: model.CardStatementPreviewRequest{
				CardType: "invalid_card",
				UserId:   user.ID,
			},
			hasError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validator.ValidateCardStatementPreviewRequest(tc.request)
			if (err != nil) != tc.hasError {
				t.Errorf("ValidateCardStatementPreviewRequest() error = %v, want error: %v", err, tc.hasError)
			}
		})
	}
}

// 保存リクエストのバリデーションテスト
func TestValidateCardStatementSaveRequest(t *testing.T) {
	// テスト用DBの設定
	db := testutils.SetupTestDB()
	defer testutils.CleanupTestDB(db)
	
	// テストユーザーの作成
	user := testutils.CreateTestUser(db)
	
	validator := NewCardStatementValidator()

	// テスト用のカード明細サマリーを作成
	validSummaries := []model.CardStatementSummary{
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

	testCases := []struct {
		name     string
		request  model.CardStatementSaveRequest
		hasError bool
	}{
		{
			name: "Valid save request with rakuten card type",
			request: model.CardStatementSaveRequest{
				CardType:       "rakuten",
				CardStatements: validSummaries,
				UserId:         user.ID,
			},
			hasError: false,
		},
		{
			name: "Valid save request with mufg card type",
			request: model.CardStatementSaveRequest{
				CardType:       "mufg",
				CardStatements: validSummaries,
				UserId:         user.ID,
			},
			hasError: false,
		},
		{
			name: "Valid save request with epos card type",
			request: model.CardStatementSaveRequest{
				CardType:       "epos",
				CardStatements: validSummaries,
				UserId:         user.ID,
			},
			hasError: false,
		},
		{
			name: "Empty card type",
			request: model.CardStatementSaveRequest{
				CardType:       "",
				CardStatements: validSummaries,
				UserId:         user.ID,
			},
			hasError: true,
		},
		{
			name: "Invalid card type",
			request: model.CardStatementSaveRequest{
				CardType:       "invalid_card",
				CardStatements: validSummaries,
				UserId:         user.ID,
			},
			hasError: true,
		},
		{
			name: "Empty card statements",
			request: model.CardStatementSaveRequest{
				CardType:       "rakuten",
				CardStatements: []model.CardStatementSummary{},
				UserId:         user.ID,
			},
			hasError: true,
		},
		{
			name: "Nil card statements",
			request: model.CardStatementSaveRequest{
				CardType:       "rakuten",
				CardStatements: nil,
				UserId:         user.ID,
			},
			hasError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validator.ValidateCardStatementSaveRequest(tc.request)
			if (err != nil) != tc.hasError {
				t.Errorf("ValidateCardStatementSaveRequest() error = %v, want error: %v", err, tc.hasError)
			}
		})
	}
}

// 月別リクエストのバリデーションテスト
func TestValidateCardStatementByMonthRequest(t *testing.T) {
	// テスト用DBの設定
	db := testutils.SetupTestDB()
	defer testutils.CleanupTestDB(db)
	
	// テストユーザーの作成
	user := testutils.CreateTestUser(db)
	
	validator := NewCardStatementValidator()

	testCases := []struct {
		name     string
		request  model.CardStatementByMonthRequest
		hasError bool
	}{
		{
			name: "Valid month request - January",
			request: model.CardStatementByMonthRequest{
				Year:   2023,
				Month:  1,
				UserId: user.ID,
			},
			hasError: false,
		},
		{
			name: "Valid month request - December",
			request: model.CardStatementByMonthRequest{
				Year:   2023,
				Month:  12,
				UserId: user.ID,
			},
			hasError: false,
		},
		{
			name: "Invalid month - 0",
			request: model.CardStatementByMonthRequest{
				Year:   2023,
				Month:  0,
				UserId: user.ID,
			},
			hasError: true,
		},
		{
			name: "Invalid month - 13",
			request: model.CardStatementByMonthRequest{
				Year:   2023,
				Month:  13,
				UserId: user.ID,
			},
			hasError: true,
		},
		{
			name: "Missing year",
			request: model.CardStatementByMonthRequest{
				Year:   0, // 0は無効な年として扱われる
				Month:  6,
				UserId: user.ID,
			},
			hasError: true,
		},
		{
			name: "Missing month",
			request: model.CardStatementByMonthRequest{
				Year:   2023,
				Month:  0, // 0は無効な月として扱われる
				UserId: user.ID,
			},
			hasError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validator.ValidateCardStatementByMonthRequest(tc.request)
			if (err != nil) != tc.hasError {
				t.Errorf("ValidateCardStatementByMonthRequest() error = %v, want error: %v", err, tc.hasError)
			}
		})
	}
}
