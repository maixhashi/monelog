package card_statement

import (
	"fmt"
	"monelog/model"
)

// GetCardStatementById はユーザーIDとカードステートメントIDに基づいてカードステートメントを取得します
func (csr *cardStatementRepository) GetCardStatementById(userId uint, cardStatementId uint) (model.CardStatement, error) {
	var cardStatement model.CardStatement
	result := csr.db.Where("user_id=?", userId).First(&cardStatement, cardStatementId)
	if result.Error != nil {
		return model.CardStatement{}, result.Error
	}
	if result.RowsAffected == 0 {
		return model.CardStatement{}, fmt.Errorf("card statement not found")
	}
	return cardStatement, nil
}