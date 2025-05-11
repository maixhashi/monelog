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
		t.Run("カード明細を保存する", func(t *testing.T) {
			// テスト用リクエストデータの作成
			cardStatements := []model.CardStatementSummary{
				{
					Type:              "発生",
					StatementNo:       1,
					CardType:          "楽天カード",
					Description:       "Amazon.co.jp",
					UseDate:           "2023/04/01",
					PaymentDate:       "2023/05/27",
					PaymentMonth:      "2023年05月",
					Amount:            5000,
					TotalChargeAmount: 5000,
					ChargeAmount:      0,
					RemainingBalance:  5000,
					PaymentCount:      0,
					InstallmentCount:  1,
					AnnualRate:        0.0,
					MonthlyRate:       0.0,
				},
				{
					Type:              "発生",
					StatementNo:       2,
					CardType:          "楽天カード",
					Description:       "楽天市場",
					UseDate:           "2023/04/15",
					PaymentDate:       "2023/05/27",
					PaymentMonth:      "2023年05月",
					Amount:            3000,
					TotalChargeAmount: 3000,
					ChargeAmount:      0,
					RemainingBalance:  3000,
					PaymentCount:      0,
					InstallmentCount:  1,
					AnnualRate:        0.0,
					MonthlyRate:       0.0,
				},
			}
			
			saveRequest := model.CardStatementSaveRequest{
				CardStatements: cardStatements,
				CardType:       "rakuten",
				Year:           2023,
				Month:          4,
			}
			
			// リクエストボディをJSON化
			requestBody, err := json.Marshal(saveRequest)
			if err != nil {
				t.Fatalf("Failed to marshal request: %v", err)
			}
			
			// テスト実行
			_, c, rec := setupEchoWithJWTAndBody(cardStatementTestUser.ID, http.MethodPost, "/card-statements/save", string(requestBody))
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
			
			if len(response) != 2 {
				t.Errorf("SaveCardStatements() returned %d card statements, want 2", len(response))
			}
			
			// カード明細の説明の確認
			descriptions := make(map[string]bool)
			for _, cs := range response {
				descriptions[cs.Description] = true
				
				// 年月が正しいことを確認
				if cs.Year != 2023 || cs.Month != 4 {
					t.Errorf("SaveCardStatements() returned statement with year=%d, month=%d, want year=2023, month=4", cs.Year, cs.Month)
				}
			}
			
			if !descriptions["Amazon.co.jp"] || !descriptions["楽天市場"] {
				t.Errorf("期待したカード明細が結果に含まれていません: %v", response)
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("空のカード明細リストでエラーになる", func(t *testing.T) {
			// 空のカード明細リストを持つリクエスト
			saveRequest := model.CardStatementSaveRequest{
				CardStatements: []model.CardStatementSummary{},
				CardType:       "rakuten",
				Year:           2023,
				Month:          4,
			}
			
			// リクエストボディをJSON化
			requestBody, err := json.Marshal(saveRequest)
			if err != nil {
				t.Fatalf("Failed to marshal request: %v", err)
			}
			
			// テスト実行
			_, c, rec := setupEchoWithJWTAndBody(cardStatementTestUser.ID, http.MethodPost, "/card-statements/save", string(requestBody))
			err = cardStatementController.SaveCardStatements(c)
			
			// 検証
			if err != nil {
				t.Errorf("SaveCardStatements() error = %v", err)
				return
			}
			
			// エラーレスポンスが返されることを確認
			if rec.Code != http.StatusBadRequest {
				t.Errorf("SaveCardStatements() status code = %d, want %d", rec.Code, http.StatusBadRequest)
			}
		})
		
		t.Run("無効な年でエラーになる", func(t *testing.T) {
			// 無効な年を持つリクエスト
			cardStatements := []model.CardStatementSummary{
				{
					Type:              "発生",
					StatementNo:       1,
					CardType:          "楽天カード",
					Description:       "Amazon.co.jp",
					UseDate:           "2023/04/01",
					PaymentDate:       "2023/05/27",
					PaymentMonth:      "2023年05月",
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
				Year:           0, // 無効な年
				Month:          4,
			}
			
			// リクエストボディをJSON化
			requestBody, err := json.Marshal(saveRequest)
			if err != nil {
				t.Fatalf("Failed to marshal request: %v", err)
			}
			
			// テスト実行
			_, c, rec := setupEchoWithJWTAndBody(cardStatementTestUser.ID, http.MethodPost, "/card-statements/save", string(requestBody))
			err = cardStatementController.SaveCardStatements(c)
			
			// 検証
			if err != nil {
				t.Errorf("SaveCardStatements() error = %v", err)
				return
			}
			
			// エラーレスポンスが返されることを確認
			if rec.Code != http.StatusBadRequest {
				t.Errorf("SaveCardStatements() status code = %d, want %d", rec.Code, http.StatusBadRequest)
			}
		})
		
		t.Run("無効な月でエラーになる", func(t *testing.T) {
			// 無効な月を持つリクエスト
			cardStatements := []model.CardStatementSummary{
				{
					Type:              "発生",
					StatementNo:       1,
					CardType:          "楽天カード",
					Description:       "Amazon.co.jp",
					UseDate:           "2023/04/01",
					PaymentDate:       "2023/05/27",
					PaymentMonth:      "2023年05月",
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
				Year:           2023,
				Month:          13, // 無効な月
			}
			
			// リクエストボディをJSON化
			requestBody, err := json.Marshal(saveRequest)
			if err != nil {
				t.Fatalf("Failed to marshal request: %v", err)
			}
			
			// テスト実行
			_, c, rec := setupEchoWithJWTAndBody(cardStatementTestUser.ID, http.MethodPost, "/card-statements/save", string(requestBody))
			err = cardStatementController.SaveCardStatements(c)
			
			// 検証
			if err != nil {
				t.Errorf("SaveCardStatements() error = %v", err)
				return
			}
			
			// エラーレスポンスが返されることを確認
			if rec.Code != http.StatusBadRequest {
				t.Errorf("SaveCardStatements() status code = %d, want %d", rec.Code, http.StatusBadRequest)
			}
		})
		
		t.Run("無効なカード種類でエラーになる", func(t *testing.T) {
			// 無効なカード種類を持つリクエスト
			cardStatements := []model.CardStatementSummary{
				{
					Type:              "発生",
					StatementNo:       1,
					CardType:          "楽天カード",
					Description:       "Amazon.co.jp",
					UseDate:           "2023/04/01",
					PaymentDate:       "2023/05/27",
					PaymentMonth:      "2023年05月",
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
				CardType:       "invalid_card_type", // 無効なカード種類
				Year:           2023,
				Month:          4,
			}
			
			// リクエストボディをJSON化
			requestBody, err := json.Marshal(saveRequest)
			if err != nil {
				t.Fatalf("Failed to marshal request: %v", err)
			}
			
			// テスト実行
			_, c, rec := setupEchoWithJWTAndBody(cardStatementTestUser.ID, http.MethodPost, "/card-statements/save", string(requestBody))
			err = cardStatementController.SaveCardStatements(c)
			
			// 検証
			if err != nil {
				t.Errorf("SaveCardStatements() error = %v", err)
				return
			}
			
			// エラーレスポンスが返されることを確認
			if rec.Code != http.StatusInternalServerError {
				t.Errorf("SaveCardStatements() status code = %d, want %d", rec.Code, http.StatusInternalServerError)
			}
		})
		
		t.Run("不正なJSONでエラーになる", func(t *testing.T) {
			// 不正なJSON
			invalidJSON := `{"card_statements": [{"invalid_json": true]}`
			
			// テスト実行
			_, c, rec := setupEchoWithJWTAndBody(cardStatementTestUser.ID, http.MethodPost, "/card-statements/save", invalidJSON)
			err := cardStatementController.SaveCardStatements(c)
			
			// 検証
			if err != nil {
				t.Errorf("SaveCardStatements() error = %v", err)
				return
			}
			
			// エラーレスポンスが返されることを確認
			if rec.Code != http.StatusBadRequest {
				t.Errorf("SaveCardStatements() status code = %d, want %d", rec.Code, http.StatusBadRequest)
			}
		})
	})
}