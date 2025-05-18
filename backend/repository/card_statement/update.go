package card_statement

import (
	"monelog/model"
)

// UpdateCardStatement はカードステートメントを更新します
func (csr *cardStatementRepository) UpdateCardStatement(cardStatement *model.CardStatement) error {
	return csr.db.Save(cardStatement).Error
}