package card_statement

import (
	"monelog/repository"
	"monelog/validator"
)

// cardStatementUsecase はカード明細に関するユースケースの実装です
type cardStatementUsecase struct {
	csr repository.ICardStatementRepository
	csv validator.ICardStatementValidator
}

// NewCardStatementUsecase はカード明細ユースケースの新しいインスタンスを作成します
func NewCardStatementUsecase(csr repository.ICardStatementRepository, csv validator.ICardStatementValidator) *cardStatementUsecase {
	return &cardStatementUsecase{csr, csv}
}