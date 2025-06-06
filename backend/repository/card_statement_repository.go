package repository

import (
	"monelog/dto"
	"monelog/model"
	"monelog/repository/card_statement"

	"gorm.io/gorm"
)

// ICardStatementRepository はカードステートメントリポジトリのインターフェースを定義します
type ICardStatementRepository interface {
	GetAllCardStatements(userId uint) ([]model.CardStatement, error)
	GetCardStatementById(userId uint, cardStatementId uint) (model.CardStatement, error)
	GetCardStatementsByMonth(userId uint, year int, month int) ([]model.CardStatement, error)
	CreateCardStatement(cardStatement *model.CardStatement) error
	CreateCardStatements(cardStatements []model.CardStatement) error
	SaveCardStatements(request *dto.CardStatementSaveRequest) error
	UpdateCardStatement(cardStatement *model.CardStatement) error
	DeleteCardStatement(cardStatement *model.CardStatement) error
	DeleteCardStatements(userId uint) error
	DeleteCardStatementsByMonth(userId uint, year int, month int) error
}

// NewCardStatementRepository はカードステートメントリポジトリの新しいインスタンスを作成します
func NewCardStatementRepository(db *gorm.DB) ICardStatementRepository {
	return card_statement.NewCardStatementRepository(db)
}