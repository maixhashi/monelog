package card_statement

import (
	"monelog/model"
	"gorm.io/gorm"
)

type CardStatementRepository interface {
	GetAllCardStatements(userId uint) ([]model.CardStatement, error)
	GetCardStatementById(userId uint, cardStatementId uint) (model.CardStatement, error)
	GetCardStatementsByMonth(userId uint, year int, month int) ([]model.CardStatement, error)
	CreateCardStatement(cardStatement *model.CardStatement) error
	CreateCardStatements(cardStatements []model.CardStatement) error
	UpdateCardStatement(cardStatement *model.CardStatement) error
	DeleteCardStatement(cardStatement *model.CardStatement) error
	DeleteCardStatements(userId uint) error
}

// cardStatementRepository はカードステートメントリポジトリの実装です
type cardStatementRepository struct {
	db *gorm.DB
}

// NewCardStatementRepository はカードステートメントリポジトリの新しいインスタンスを作成します
func NewCardStatementRepository(db *gorm.DB) CardStatementRepository {
	return &cardStatementRepository{db}
}