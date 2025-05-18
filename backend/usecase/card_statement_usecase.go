package usecase

import (
	"mime/multipart"
	"monelog/model"
	"monelog/repository"
	"monelog/usecase/card_statement"
	"monelog/validator"
)

// ICardStatementUsecase はカード明細に関するユースケースのインターフェースを定義します
type ICardStatementUsecase interface {
	GetAllCardStatements(userId uint) ([]model.CardStatementResponse, error)
	GetCardStatementById(userId uint, cardStatementId uint) (model.CardStatementResponse, error)
	ProcessCSV(file *multipart.FileHeader, request model.CardStatementRequest) ([]model.CardStatementResponse, error)
	PreviewCSV(file *multipart.FileHeader, request model.CardStatementPreviewRequest) ([]model.CardStatementResponse, error)
	SaveCardStatements(request model.CardStatementSaveRequest) ([]model.CardStatementResponse, error)
	GetCardStatementsByMonth(request model.CardStatementByMonthRequest) ([]model.CardStatementResponse, error)
}

// NewCardStatementUsecase は新しいカード明細ユースケースのインスタンスを作成します
func NewCardStatementUsecase(csr repository.ICardStatementRepository, csv validator.ICardStatementValidator) ICardStatementUsecase {
	return card_statement.NewCardStatementUsecase(csr, csv)
}
