package card_statement

import (
	"monelog/dto"
	"monelog/mapper"
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

// SaveCardStatements はDTOからモデルに変換して保存します
func (csr *cardStatementRepository) SaveCardStatements(request *dto.CardStatementSaveRequest) error {
	// 既存のデータを削除
	if err := csr.DeleteCardStatementsByMonth(request.UserId, request.Year, request.Month); err != nil {
		return err
	}

	// 新しいデータを保存
	cardStatements := make([]model.CardStatement, len(request.CardStatements))
	for i, summary := range request.CardStatements {
		cardStatements[i] = mapper.ToCardStatementModel(&summary, request.UserId, request.Year, request.Month)
	}

	return csr.CreateCardStatements(cardStatements)
}