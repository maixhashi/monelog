package card_statement

import (
	"monelog/model"
)

// DeleteCardStatements はユーザーIDに基づいてカードステートメントを削除します
func (csr *cardStatementRepository) DeleteCardStatements(userId uint) error {
	result := csr.db.Where("user_id=?", userId).Delete(&model.CardStatement{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}