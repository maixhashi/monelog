package card_statement

import (
	"monelog/model"
)

// CreateCardStatement は新しいカードステートメントを作成します
func (csr *cardStatementRepository) CreateCardStatement(cardStatement *model.CardStatement) error {
	return csr.db.Create(cardStatement).Error
}

// CreateCardStatements は複数のカードステートメントを一括で作成します
func (csr *cardStatementRepository) CreateCardStatements(cardStatements []model.CardStatement) error {
	if len(cardStatements) == 0 {
		return nil // 空の配列の場合は早期リターン
	}
	return csr.db.Create(&cardStatements).Error
}