package usecase

import (
	"bytes"
	"io"
	"mime/multipart"
	"monelog/model"
	"monelog/repository"
	"monelog/validator"
)

type ICSVHistoryUsecase interface {
	GetAllCSVHistories(userId uint) ([]model.CSVHistoryResponse, error)
	GetCSVHistoryById(userId uint, csvHistoryId uint) (model.CSVHistoryDetailResponse, error)
	SaveCSVHistory(file *multipart.FileHeader, request model.CSVHistorySaveRequest) (model.CSVHistoryResponse, error)
	DeleteCSVHistory(userId uint, csvHistoryId uint) error
	// 新しいメソッドを追加
	GetCSVHistoriesByMonth(userId uint, year int, month int) ([]model.CSVHistoryResponse, error)
}

type csvHistoryUsecase struct {
	chr repository.ICSVHistoryRepository
	chv validator.ICSVHistoryValidator
}

func NewCSVHistoryUsecase(chr repository.ICSVHistoryRepository, chv validator.ICSVHistoryValidator) ICSVHistoryUsecase {
	return &csvHistoryUsecase{chr, chv}
}

func (chu *csvHistoryUsecase) GetAllCSVHistories(userId uint) ([]model.CSVHistoryResponse, error) {
	csvHistories, err := chu.chr.GetAllCSVHistories(userId)
	if err != nil {
		return nil, err
	}

	responses := make([]model.CSVHistoryResponse, len(csvHistories))
	for i, csvHistory := range csvHistories {
		responses[i] = csvHistory.ToResponse()
	}
	return responses, nil
}

func (chu *csvHistoryUsecase) GetCSVHistoryById(userId uint, csvHistoryId uint) (model.CSVHistoryDetailResponse, error) {
	csvHistory, err := chu.chr.GetCSVHistoryById(userId, csvHistoryId)
	if err != nil {
		return model.CSVHistoryDetailResponse{}, err
	}
	return csvHistory.ToDetailResponse(), nil
}

func (chu *csvHistoryUsecase) SaveCSVHistory(file *multipart.FileHeader, request model.CSVHistorySaveRequest) (model.CSVHistoryResponse, error) {
	if err := chu.chv.ValidateCSVHistorySaveRequest(request); err != nil {
		return model.CSVHistoryResponse{}, err
	}

	// ファイルを開く
	src, err := file.Open()
	if err != nil {
		return model.CSVHistoryResponse{}, err
	}
	defer src.Close()

	// ファイルの内容を読み込む
	buf := new(bytes.Buffer)
	if _, err = io.Copy(buf, src); err != nil {
		return model.CSVHistoryResponse{}, err
	}

	// CSVHistoryモデルを作成
	csvHistory := model.CSVHistory{
		FileName: request.FileName,
		CardType: request.CardType,
		FileData: buf.Bytes(),
		UserId:   request.UserId,
		Year:     request.Year,   // 年を保存
		Month:    request.Month,  // 月を保存
	}

	// データベースに保存
	if err := chu.chr.CreateCSVHistory(&csvHistory); err != nil {
		return model.CSVHistoryResponse{}, err
	}

	return csvHistory.ToResponse(), nil
}

func (chu *csvHistoryUsecase) DeleteCSVHistory(userId uint, csvHistoryId uint) error {
	return chu.chr.DeleteCSVHistory(userId, csvHistoryId)
}

// 新しく追加するメソッド
func (chu *csvHistoryUsecase) GetCSVHistoriesByMonth(userId uint, year int, month int) ([]model.CSVHistoryResponse, error) {
	// リポジトリから指定された年月のCSV履歴を取得
	csvHistories, err := chu.chr.GetCSVHistoriesByMonth(userId, year, month)
	if err != nil {
		return nil, err
	}

	// レスポンス形式に変換
	responses := make([]model.CSVHistoryResponse, len(csvHistories))
	for i, csvHistory := range csvHistories {
		responses[i] = csvHistory.ToResponse()
	}
	
	return responses, nil
}
