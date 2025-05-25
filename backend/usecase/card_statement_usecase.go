package usecase

import (
	"mime/multipart"
	"monelog/dto"
	"monelog/repository"
	"monelog/usecase/card_statement"
	"monelog/validator"
)

// ICardStatementUsecase はカード明細に関するユースケースのインターフェースを定義します
type ICardStatementUsecase interface {
	GetAllCardStatements(userId uint) ([]dto.CardStatementResponse, error)
	GetCardStatementById(userId uint, cardStatementId uint) (dto.CardStatementResponse, error)
	ProcessCSV(file *multipart.FileHeader, request dto.CardStatementRequest) ([]dto.CardStatementResponse, error)
	PreviewCSV(file *multipart.FileHeader, request dto.CardStatementPreviewRequest) ([]dto.CardStatementResponse, error)
	SaveCardStatements(request dto.CardStatementSaveRequest) ([]dto.CardStatementResponse, error)
	GetCardStatementsByMonth(request dto.CardStatementByMonthRequest) ([]dto.CardStatementResponse, error)
}

// NewCardStatementUsecase は新しいカード明細ユースケースのインスタンスを作成します
func NewCardStatementUsecase(csr repository.ICardStatementRepository, csv validator.ICardStatementValidator) ICardStatementUsecase {
	return card_statement.NewCardStatementUsecase(csr, csv)
}
