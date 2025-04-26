package card_statement_test

import (
	"encoding/json"
	"monelog/model"
	"net/http"
	"testing"
)

func TestCardStatementController_SaveCardStatements(t *testing.T) {
	setupCardStatementControllerTest()

	t.Run("正常系", func(t *testing.T) {
		t.Run("カード明細を正常に保存する", func(t *testing.T) {
			// リクエストボディの作成
			cardStatements := []model.CardStatementSummary{
				{
					Type:              "発生",
					StatementNo:       1,
					CardType:          "楽天カード",
					Description:       "テスト明細1",
					UseDate:           "2023/01/15",
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
				{
					Type:              "発生",
					StatementNo:       2,
					CardType:          "楽天カード",
					Description:       "テスト明細2",
					UseDate:           "2023/01/20",
					PaymentDate:       "2023/02/27",
					PaymentMonth:      "2023年02月",
					Amount:            5000,
					TotalChargeAmount: 5000,
					ChargeAmount:      0,
					RemainingBalance:  5000,
					PaymentCount:      0,
					InstallmentCount:  1,
					AnnualRate:        0.0,
					MonthlyRate:       0.0,
				},
			}

			saveRequest := model.CardStatementSaveRequest{
				CardStatements: cardStatements,
				CardType:       "rakuten",
			}

			// JSONに変換
			jsonBody, err := json.Marshal(saveRequest)
			if err != nil {
				t.Fatalf("Failed to marshal request body: %v", err)
			}

			// テスト実行
			_, c, rec := setupEchoWithJWTAndBody(cardStatementTestUser.ID, http.MethodPost, "/card-statements/save", string(jsonBody))
			err = cardStatementController.SaveCardStatements(c)

			// 検証
			if err != nil {
				t.Errorf("SaveCardStatements() error = %v", err)
			}

			if rec.Code != http.StatusCreated {
				t.Errorf("SaveCardStatements() status code = %d, want %d", rec.Code, http.StatusCreated)
			}

			// レスポンスボディをパース
			response := parseCardStatementsResponse(t, rec.Body.Bytes())

			// 保存されたカード明細の数を確認
			if len(response) != 2 {
				t.Errorf("SaveCardStatements() returned %d card statements, want 2", len(response))
			}

			// 保存されたカード明細の内容を確認
			descriptions := make(map[string]bool)
			for _, cs := range response {
				descriptions[cs.Description] = true
			}

			if !descriptions["テスト明細1"] || !descriptions["テスト明細2"] {
				t.Errorf("SaveCardStatements() did not return expected descriptions: %v", descriptions)
			}
		})
	})

	t.Run("異常系", func(t *testing.T) {
		t.Run("不正なリクエスト形式の場合はエラーを返す", func(t *testing.T) {
			// 不正なJSONを作成
			invalidJSON := `{"card_type": "rakuten", "card_statements": "not_an_array"}`

			// テスト実行
			_, c, rec := setupEchoWithJWTAndBody(cardStatementTestUser.ID, http.MethodPost, "/card-statements/save", invalidJSON)
			err := cardStatementController.SaveCardStatements(c)

			// エラーは返さないが、ステータスコードは400になる想定
			if err != nil {
				t.Errorf("SaveCardStatements() error = %v", err)
			}

			if rec.Code != http.StatusBadRequest {
				t.Errorf("SaveCardStatements() status code = %d, want %d", rec.Code, http.StatusBadRequest)
			}
		})

		t.Run("card_typeが指定されていない場合はエラーを返す", func(t *testing.T) {
			// card_typeなしのリクエストを作成
			saveRequest := model.CardStatementSaveRequest{
				CardStatements: []model.CardStatementSummary{
					{
						Type:              "発生",
						StatementNo:       1,
						CardType:          "楽天カード",
						Description:       "テスト明細",
						UseDate:           "2023/01/15",
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
				},
				// CardTypeを省略
			}

			// JSONに変換
			jsonBody, err := json.Marshal(saveRequest)
			if err != nil {
				t.Fatalf("Failed to marshal request body: %v", err)
			}

			// テスト実行
			_, c, rec := setupEchoWithJWTAndBody(cardStatementTestUser.ID, http.MethodPost, "/card-statements/save", string(jsonBody))
			err = cardStatementController.SaveCardStatements(c)

			// エラーは返さないが、ステータスコードは500になる想定（バリデーションエラー）
			if err != nil {
				t.Errorf("SaveCardStatements() error = %v", err)
			}

			if rec.Code != http.StatusInternalServerError {
				t.Errorf("SaveCardStatements() status code = %d, want %d", rec.Code, http.StatusInternalServerError)
			}
		})

		t.Run("無効なcard_typeが指定された場合はエラーを返す", func(t *testing.T) {
			// 無効なcard_typeのリクエストを作成
			saveRequest := model.CardStatementSaveRequest{
				CardStatements: []model.CardStatementSummary{
					{
						Type:              "発生",
						StatementNo:       1,
						CardType:          "楽天カード",
						Description:       "テスト明細",
						UseDate:           "2023/01/15",
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
				},
				CardType: "invalid_card_type", // 無効なカード種類
			}

			// JSONに変換
			jsonBody, err := json.Marshal(saveRequest)
			if err != nil {
				t.Fatalf("Failed to marshal request body: %v", err)
			}

			// テスト実行
			_, c, rec := setupEchoWithJWTAndBody(cardStatementTestUser.ID, http.MethodPost, "/card-statements/save", string(jsonBody))
			err = cardStatementController.SaveCardStatements(c)

			// エラーは返さないが、ステータスコードは500になる想定（バリデーションエラー）
			if err != nil {
				t.Errorf("SaveCardStatements() error = %v", err)
			}

			if rec.Code != http.StatusInternalServerError {
				t.Errorf("SaveCardStatements() status code = %d, want %d", rec.Code, http.StatusInternalServerError)
			}
		})

		t.Run("card_statementsが空の場合はエラーを返す", func(t *testing.T) {
			// 空のcard_statementsのリクエストを作成
			saveRequest := model.CardStatementSaveRequest{
				CardStatements: []model.CardStatementSummary{}, // 空の配列
				CardType:       "rakuten",
			}

			// JSONに変換
			jsonBody, err := json.Marshal(saveRequest)
			if err != nil {
				t.Fatalf("Failed to marshal request body: %v", err)
			}

			// テスト実行
			_, c, rec := setupEchoWithJWTAndBody(cardStatementTestUser.ID, http.MethodPost, "/card-statements/save", string(jsonBody))
			err = cardStatementController.SaveCardStatements(c)

			// エラーは返さないが、ステータスコードは500になる想定（バリデーションエラー）
			if err != nil {
				t.Errorf("SaveCardStatements() error = %v", err)
			}

			if rec.Code != http.StatusInternalServerError {
				t.Errorf("SaveCardStatements() status code = %d, want %d", rec.Code, http.StatusInternalServerError)
			}
		})
	})
}