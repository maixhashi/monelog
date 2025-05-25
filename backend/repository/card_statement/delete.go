package card_statement

import (
	"monelog/model"
)

// DeleteCardStatement はカードステートメントを削除します
func (csr *cardStatementRepository) DeleteCardStatement(cardStatement *model.CardStatement) error {
	return csr.db.Delete(cardStatement).Error
}

// DeleteCardStatements はユーザーIDに基づいてすべてのカードステートメントを削除します
func (csr *cardStatementRepository) DeleteCardStatements(userId uint) error {
	return csr.db.Where("user_id = ?", userId).Delete(&model.CardStatement{}).Error
}

// DeleteCardStatementsByMonth はユーザーID、年、月に基づいてカードステートメントを削除します
func (csr *cardStatementRepository) DeleteCardStatementsByMonth(userId uint, year int, month int) error {
	return csr.db.Where("user_id = ? AND year = ? AND month = ?", userId, year, month).
		Delete(&model.CardStatement{}).Error
}