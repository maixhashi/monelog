package card_statement_test

import (
	"monelog/model"
	"testing"
)

func TestCardStatementRepository_GetAllCardStatements(t *testing.T) {
	setupCardStatementTest()
	
	cardStatements := []model.CardStatement{
		{Type: "発生", CardType: "楽天カード", Description: "Amazon", Amount: 1000, UserId: csTestUser.ID},
		{Type: "発生", CardType: "楽天カード", Description: "楽天市場", Amount: 2000, UserId: csTestUser.ID},
		{Type: "発生", CardType: "MUFG", Description: "ヨドバシカメラ", Amount: 3000, UserId: csOtherUser.ID},
	}
	
	for _, cs := range cardStatements {
		csDB.Create(&cs)
	}
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("正しいユーザーIDのカード明細のみを取得する", func(t *testing.T) {
			result, err := csRepo.GetAllCardStatements(csTestUser.ID)
			
			if err != nil {
				t.Errorf("GetAllCardStatements() error = %v", err)
			}
			
			if len(result) != 2 {
				t.Errorf("GetAllCardStatements() got %d card statements, want 2", len(result))
			}
			
			descriptions := make(map[string]bool)
			for _, cs := range result {
				descriptions[cs.Description] = true
			}
			
			if !descriptions["Amazon"] || !descriptions["楽天市場"] {
				t.Errorf("期待したカード明細が結果に含まれていません: %v", result)
			}
		})
	})
}
