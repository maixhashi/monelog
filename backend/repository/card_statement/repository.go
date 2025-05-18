package card_statement

import (
	"gorm.io/gorm"
)

// cardStatementRepository はカードステートメントリポジトリの実装です
type cardStatementRepository struct {
	db *gorm.DB
}

// NewCardStatementRepository はカードステートメントリポジトリの新しいインスタンスを作成します
func NewCardStatementRepository(db *gorm.DB) *cardStatementRepository {
	return &cardStatementRepository{db: db}
}