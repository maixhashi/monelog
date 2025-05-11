package card_statement_test

import (
	"testing"
)

func TestCardStatementRepository_GetCardStatementsByMonth(t *testing.T) {
	setupCardStatementTest()
	
	// テストデータの作成
	createTestCardStatement("楽天カード", "1月明細1", cardStatementTestUser.ID, 2023, 1)
	createTestCardStatement("楽天カード", "1月明細2", cardStatementTestUser.ID, 2023, 1)
	createTestCardStatement("楽天カード", "2月明細1", cardStatementTestUser.ID, 2023, 2)
	createTestCardStatement("楽天カード", "他ユーザー1月明細", cardStatementOtherUser.ID, 2023, 1)
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("指定した年月のカード明細のみを取得する", func(t *testing.T) {
			result, err := cardStatementRepo.GetCardStatementsByMonth(cardStatementTestUser.ID, 2023, 1)
			
			if err != nil {
				t.Errorf("GetCardStatementsByMonth() error = %v", err)
			}
			
			if len(result) != 2 {
				t.Errorf("GetCardStatementsByMonth() got %d statements, want 2", len(result))
			}
			
			descriptions := make(map[string]bool)
			for _, statement := range result {
				descriptions[statement.Description] = true
				
				if statement.Year != 2023 || statement.Month != 1 {
					t.Errorf("取得したカード明細の年月が一致しません: got %d年%d月, want 2023年1月", 
						statement.Year, statement.Month)
				}
			}
			
			if !descriptions["1月明細1"] || !descriptions["1月明細2"] {
				t.Errorf("期待したカード明細が結果に含まれていません: %v", result)
			}
		})
		
		t.Run("存在しない年月の場合は空の配列を返す", func(t *testing.T) {
			result, err := cardStatementRepo.GetCardStatementsByMonth(cardStatementTestUser.ID, 2022, 12)
			
			if err != nil {
				t.Errorf("GetCardStatementsByMonth() error = %v", err)
			}
			
			if len(result) != 0 {
				t.Errorf("GetCardStatementsByMonth() got %d statements, want 0", len(result))
			}
		})
	})
}