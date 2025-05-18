package card_statement

import (
	"monelog/model"
)

// GetCardStatementsByMonth はユーザーID、年、月に基づいてカードステートメントを取得します
func (csr *cardStatementRepository) GetCardStatementsByMonth(userId uint, year int, month int) ([]model.CardStatement, error) {
	var cardStatements []model.CardStatement
	if err := csr.db.Where("user_id = ? AND year = ? AND month = ?", userId, year, month).
		Order("statement_no, payment_count").
		Find(&cardStatements).Error; err != nil {
		return nil, err
	}
	return cardStatements, nil
}