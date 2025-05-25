package card_statement_test

import (
	"monelog/dto"
	"testing"
)

func TestCardStatementUsecase_SaveCardStatements(t *testing.T) {
	setupCardStatementUsecaseTest()
	
	// 既存のデータを作成（削除されることを確認するため）
	existingStatement := createTestCardStatement(t, "既存の明細", testUser.ID, 2023, 1)
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("カード明細を正しく保存する", func(t *testing.T) {
			// テスト用のカード明細サマリーを作成
			summaries := []dto.CardStatementSummary{
				createTestCardStatementSummary("新規明細1"),
				createTestCardStatementSummary("新規明細2"),
			}
			
			request := dto.CardStatementSaveRequest{
				CardStatements: summaries,
				CardType:       "rakuten",
				Year:           2023,
				Month:          2,
				UserId:         testUser.ID,
			}
			
			t.Logf("ユーザーID %d の明細を保存します", testUser.ID)
			
			responses, err := cardStatementUsecase.SaveCardStatements(request)
			
			if err != nil {
				t.Errorf("SaveCardStatements() error = %v", err)
			}
			
			if !validateCardStatementResponses(t, responses, 2) {
				return
			}
			
			// 既存のデータが削除されていることを確認
			assertCardStatementNotExists(t, existingStatement.ID)
			
			// 新しいデータが保存されていることを確認
			descriptions := make(map[string]bool)
			for _, response := range responses {
				descriptions[response.Description] = true
				
				// 年月が正しく設定されていることを確認
				if response.Year != 2023 || response.Month != 2 {
					t.Errorf("年月が正しく設定されていません: got=%d年%d月, want=2023年2月", response.Year, response.Month)
				}
				
				// データベースに存在することを確認
				assertCardStatementExists(t, response.ID, response.Description, testUser.ID)
			}
			
			if !descriptions["新規明細1"] || !descriptions["新規明細2"] {
				t.Errorf("期待した明細が結果に含まれていません: %v", responses)
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("バリデーションエラー - 空の明細リスト", func(t *testing.T) {
			request := dto.CardStatementSaveRequest{
				CardStatements: []dto.CardStatementSummary{}, // 空のリスト
				CardType:       "rakuten",
				Year:           2023,
				Month:          2,
				UserId:         testUser.ID,
			}
			
			t.Logf("空の明細リストで保存を試みます")
			
			_, err := cardStatementUsecase.SaveCardStatements(request)
			
			if err == nil {
				t.Error("空の明細リストでエラーが返されませんでした")
			} else {
				t.Logf("期待通りエラーが返されました: %v", err)
			}
		})
		
		t.Run("バリデーションエラー - 無効なカード種類", func(t *testing.T) {
			summaries := []dto.CardStatementSummary{
				createTestCardStatementSummary("テスト明細"),
			}
			
			request := dto.CardStatementSaveRequest{
				CardStatements: summaries,
				CardType:       "invalid_card", // 無効なカード種類
				Year:           2023,
				Month:          2,
				UserId:         testUser.ID,
			}
			
			t.Logf("無効なカード種類で保存を試みます")
			
			_, err := cardStatementUsecase.SaveCardStatements(request)
			
			if err == nil {
				t.Error("無効なカード種類でエラーが返されませんでした")
			} else {
				t.Logf("期待通りエラーが返されました: %v", err)
			}
		})
	})
}
