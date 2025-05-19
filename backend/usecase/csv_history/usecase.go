package csv_history

import (
	"monelog/model"
	"monelog/repository"
	"monelog/validator"
	"mime/multipart"
)

type ICSVHistoryUsecase interface {
	GetAllCSVHistories(userId uint) ([]model.CSVHistoryResponse, error)
	GetCSVHistoryById(userId uint, csvHistoryId uint) (model.CSVHistoryDetailResponse, error)
	SaveCSVHistory(file *multipart.FileHeader, request model.CSVHistorySaveRequest) (model.CSVHistoryResponse, error)
	DeleteCSVHistory(userId uint, csvHistoryId uint) error
	GetCSVHistoriesByMonth(userId uint, year int, month int) ([]model.CSVHistoryResponse, error)
}

// csvHistoryUsecase はCSV履歴に関するユースケースの実装です
type csvHistoryUsecase struct {
	chr repository.ICSVHistoryRepository
	chv validator.ICSVHistoryValidator
}

// NewCSVHistoryUsecase はCSV履歴ユースケースの新しいインスタンスを作成します
func NewCSVHistoryUsecase(chr repository.ICSVHistoryRepository, chv validator.ICSVHistoryValidator) ICSVHistoryUsecase {
	return &csvHistoryUsecase{chr, chv}
}