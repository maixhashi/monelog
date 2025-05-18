package card_statement

import (
	"monelog/model"
)

// CreateCardStatement はカードステートメントを作成します
func (csr *cardStatementRepository) CreateCardStatement(cardStatement *model.CardStatement) error {
	return csr.db.Create(cardStatement).Error
}

// CreateCardStatements は複数のカードステートメントを一括で作成します
func (csr *cardStatementRepository) CreateCardStatements(cardStatements []model.CardStatement) error {
	// 空のスライスの場合は早期リターン
	if len(cardStatements) == 0 {
		return nil
	}
	return csr.db.Create(&cardStatements).Error
}