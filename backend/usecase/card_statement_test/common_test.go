package card_statement_test

import (
	"monelog/model"
	"testing"
)

// カード明細の存在を確認するヘルパー関数
func assertCardStatementExists(t *testing.T, cardStatementId uint, expectedDescription string, expectedUserId uint) bool {
	var cardStatement model.CardStatement
	result := cardStatementDb.First(&cardStatement, cardStatementId)
	
	if result.Error != nil {
		t.Errorf("カード明細(ID=%d)がデータベースに存在しません: %v", cardStatementId, result.Error)
		return false
	}
	
	if cardStatement.Description != expectedDescription {
		t.Errorf("カード明細の説明が一致しません: got=%s, want=%s", cardStatement.Description, expectedDescription)
		return false
	}
	
	if cardStatement.UserId != expectedUserId {
		t.Errorf("カード明細のユーザーIDが一致しません: got=%d, want=%d", cardStatement.UserId, expectedUserId)
		return false
	}
	
	return true
}

// カード明細が存在しないことを確認するヘルパー関数
func assertCardStatementNotExists(t *testing.T, cardStatementId uint) bool {
	var count int64
	cardStatementDb.Model(&model.CardStatement{}).Where("id = ?", cardStatementId).Count(&count)
	
	if count != 0 {
		t.Errorf("カード明細(ID=%d)がデータベースに存在します", cardStatementId)
		return false
	}
	
	return true
}

// カード明細レスポンスの検証ヘルパー関数
func validateCardStatementResponse(t *testing.T, cardStatement model.CardStatementResponse, expectedId uint, expectedDescription string) bool {
	if cardStatement.ID != expectedId {
		t.Errorf("カード明細IDが一致しません: got=%d, want=%d", cardStatement.ID, expectedId)
		return false
	}
	
	if cardStatement.Description != expectedDescription {
		t.Errorf("カード明細の説明が一致しません: got=%s, want=%s", cardStatement.Description, expectedDescription)
		return false
	}
	
	if cardStatement.CreatedAt.IsZero() || cardStatement.UpdatedAt.IsZero() {
		t.Errorf("カード明細のタイムスタンプが正しく設定されていません: %+v", cardStatement)
		return false
	}
	
	return true
}
