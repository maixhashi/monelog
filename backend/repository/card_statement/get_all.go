package card_statement

import (
	"monelog/model"
)

// GetAllCardStatements はユーザーIDに基づいてすべてのカードステートメントを取得します
func (csr *cardStatementRepository) GetAllCardStatements(userId uint) ([]model.CardStatement, error) {
	var cardStatements []model.CardStatement
	if err := csr.db.Where("user_id=?", userId).Order("statement_no, payment_count").Find(&cardStatements).Error; err != nil {
		return nil, err
	}
	return cardStatements, nil
}