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
