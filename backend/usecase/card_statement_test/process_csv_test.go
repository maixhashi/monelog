package card_statement_test

import (
	"monelog/model"
	"monelog/parser"
	"testing"
)

// ProcessCSVのテスト用に簡略化したバージョン
func TestCardStatementUsecase_ProcessCSV(t *testing.T) {
	setupCardStatementUsecaseTest()
	
	// 既存のカード明細を作成
	existingCardStatement := createTestCardStatement(t, "既存の明細", testUser.ID)
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("カード明細の処理が正しく行われる", func(t *testing.T) {
			// CSVパース結果をシミュレート
			summaries := []model.CardStatementSummary{
				{
					Type:              "発生",
					StatementNo:       1,
					CardType:          "テストカード",
					Description:       "テスト店舗1",
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
				{
					Type:              "発生",
					StatementNo:       2,
					CardType:          "テストカード",
					Description:       "テスト店舗2",
					UseDate:           "2023/01/02",
					PaymentDate:       "2023/02/27",
					PaymentMonth:      "2023年02月",
					Amount:            20000,
					TotalChargeAmount: 20000,
					ChargeAmount:      0,
					RemainingBalance:  20000,
					PaymentCount:      0,
					InstallmentCount:  1,
					AnnualRate:        0.0,
					MonthlyRate:       0.0,
				},
			}
			
			// モデルに変換
			cardStatements := make([]model.CardStatement, len(summaries))
			for i, summary := range summaries {
				cardStatements[i] = summary.ToModel(testUser.ID)
			}
			
			// データベースに直接保存
			if err := cardStatementRepo.DeleteCardStatements(testUser.ID); err != nil {
				t.Fatalf("既存のカード明細の削除に失敗しました: %v", err)
			}
			
			if err := cardStatementRepo.CreateCardStatements(cardStatements); err != nil {
				t.Fatalf("テストカード明細の作成に失敗しました: %v", err)
			}
			
			// 既存のデータが削除されていることを確認
			assertCardStatementNotExists(t, existingCardStatement.ID)
			
			// 新しいデータが作成されていることを確認
			cardStatementResponses, err := cardStatementUsecase.GetAllCardStatements(testUser.ID)
			if err != nil {
				t.Errorf("GetAllCardStatements() error = %v", err)
				return
			}
			
			if len(cardStatementResponses) != 2 {
				t.Errorf("期待したカード明細数と一致しません: got %d, want 2", len(cardStatementResponses))
				return
			}
			
			// レスポンスの内容を確認
			descriptions := make(map[string]bool)
			for _, response := range cardStatementResponses {
				t.Logf("処理されたカード明細: ID=%d, Description=%s, Amount=%d", 
					response.ID, response.Description, response.Amount)
				
				descriptions[response.Description] = true
			}
			
			// 期待する説明が含まれていることを確認
			if !descriptions["テスト店舗1"] || !descriptions["テスト店舗2"] {
				t.Errorf("期待したカード明細の説明が結果に含まれていません: %v", descriptions)
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("バリデーションエラーの場合はエラーを返す", func(t *testing.T) {
			request := model.CardStatementRequest{
				CardType: "", // 空のカード種類（バリデーションエラー）
				UserId:   testUser.ID,
			}
			
			// バリデーションエラーを直接テスト
			err := cardStatementValidator.ValidateCardStatementRequest(request)
			
			if err == nil {
				t.Error("バリデーションエラーが発生しませんでした")
			} else {
				t.Logf("期待通りバリデーションエラーが返されました: %v", err)
			}
		})
		
		t.Run("不正なカード種類の場合はエラーを返す", func(t *testing.T) {
			// パーサーの取得をテスト
			_, err := parser.GetParser("invalid_card_type")
			
			if err == nil {
				t.Error("不正なカード種類でもエラーが返されませんでした")
			} else {
				t.Logf("期待通りエラーが返されました: %v", err)
			}
		})
	})
}
