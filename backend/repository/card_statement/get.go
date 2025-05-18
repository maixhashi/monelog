package card_statement

import (
	"fmt"
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