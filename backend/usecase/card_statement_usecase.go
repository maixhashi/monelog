package usecase

import (
	"bytes"
	"io"
	"mime/multipart"
	"monelog/model"
	"monelog/parser"
	"monelog/repository"
	"monelog/validator"
)

type ICardStatementUsecase interface {
	GetAllCardStatements(userId uint) ([]model.CardStatementResponse, error)
	GetCardStatementById(userId uint, cardStatementId uint) (model.CardStatementResponse, error)
	ProcessCSV(file *multipart.FileHeader, request model.CardStatementRequest) ([]model.CardStatementResponse, error)
	PreviewCSV(file *multipart.FileHeader, request model.CardStatementPreviewRequest) ([]model.CardStatementResponse, error)
	SaveCardStatements(request model.CardStatementSaveRequest) ([]model.CardStatementResponse, error)
	GetCardStatementsByMonth(request model.CardStatementByMonthRequest) ([]model.CardStatementResponse, error)
}

type cardStatementUsecase struct {
	csr repository.ICardStatementRepository
	csv validator.ICardStatementValidator
}

func NewCardStatementUsecase(csr repository.ICardStatementRepository, csv validator.ICardStatementValidator) ICardStatementUsecase {
	return &cardStatementUsecase{csr, csv}
}

func (csu *cardStatementUsecase) GetAllCardStatements(userId uint) ([]model.CardStatementResponse, error) {
	cardStatements, err := csu.csr.GetAllCardStatements(userId)
	if err != nil {
		return nil, err
	}

	responses := make([]model.CardStatementResponse, len(cardStatements))
	for i, cardStatement := range cardStatements {
		responses[i] = cardStatement.ToResponse()
	}
	return responses, nil
}

func (csu *cardStatementUsecase) GetCardStatementById(userId uint, cardStatementId uint) (model.CardStatementResponse, error) {
	cardStatement, err := csu.csr.GetCardStatementById(userId, cardStatementId)
	if err != nil {
		return model.CardStatementResponse{}, err
	}
	return cardStatement.ToResponse(), nil
}

func (csu *cardStatementUsecase) ProcessCSV(file *multipart.FileHeader, request model.CardStatementRequest) ([]model.CardStatementResponse, error) {
	if err := csu.csv.ValidateCardStatementRequest(request); err != nil {
		return nil, err
	}

	// ファイルを開く
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	// ファイルの内容を読み込む
	buf := new(bytes.Buffer)
	if _, err = io.Copy(buf, src); err != nil {
		return nil, err
	}

	// カード種類に応じたパーサーを取得
	cardParser, err := parser.GetParser(request.CardType)
	if err != nil {
		return nil, err
	}

	// CSVを解析
	summaries, err := cardParser.Parse(buf.Bytes())
	if err != nil {
		return nil, err
	}

	// 既存のデータを削除
	if err := csu.csr.DeleteCardStatements(request.UserId); err != nil {
		return nil, err
	}

	// 解析結果をデータベースに保存
	cardStatements := make([]model.CardStatement, len(summaries))
	for i, summary := range summaries {
		cardStatements[i] = summary.ToModel(request.UserId)
	}

	if err := csu.csr.CreateCardStatements(cardStatements); err != nil {
		return nil, err
	}

	// レスポンスを作成
	responses := make([]model.CardStatementResponse, len(cardStatements))
	for i, cardStatement := range cardStatements {
		responses[i] = cardStatement.ToResponse()
	}

	return responses, nil
}

func (csu *cardStatementUsecase) PreviewCSV(file *multipart.FileHeader, request model.CardStatementPreviewRequest) ([]model.CardStatementResponse, error) {
	if err := csu.csv.ValidateCardStatementPreviewRequest(request); err != nil {
		return nil, err
	}

	// ファイルを開く
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	// ファイルの内容を読み込む
	buf := new(bytes.Buffer)
	if _, err = io.Copy(buf, src); err != nil {
		return nil, err
	}

	// カード種類に応じたパーサーを取得
	cardParser, err := parser.GetParser(request.CardType)
	if err != nil {
		return nil, err
	}

	// CSVを解析
	summaries, err := cardParser.Parse(buf.Bytes())
	if err != nil {
		return nil, err
	}

	// レスポンスを作成（DBには保存しない）
	responses := make([]model.CardStatementResponse, len(summaries))
	for i, summary := range summaries {
		cardStatement := summary.ToModel(request.UserId)
		responses[i] = cardStatement.ToResponse()
	}

	return responses, nil
}

func (csu *cardStatementUsecase) SaveCardStatements(request model.CardStatementSaveRequest) ([]model.CardStatementResponse, error) {
	if err := csu.csv.ValidateCardStatementSaveRequest(request); err != nil {
		return nil, err
	}

	// 既存のデータを削除
	if err := csu.csr.DeleteCardStatements(request.UserId); err != nil {
		return nil, err
	}

	// 新しいデータを保存
	cardStatements := make([]model.CardStatement, len(request.CardStatements))
	for i, summary := range request.CardStatements {
		cardStatements[i] = summary.ToModel(request.UserId)
	}

	if err := csu.csr.CreateCardStatements(cardStatements); err != nil {
		return nil, err
	}

	// レスポンスを作成
	responses := make([]model.CardStatementResponse, len(cardStatements))
	for i, cardStatement := range cardStatements {
		responses[i] = cardStatement.ToResponse()
	}

	return responses, nil
}

func (csu *cardStatementUsecase) GetCardStatementsByMonth(request model.CardStatementByMonthRequest) ([]model.CardStatementResponse, error) {
	// Validate the request
	if err := csu.csv.ValidateCardStatementByMonthRequest(request); err != nil {
		return nil, err
	}
	
	// Get card statements for the specified month
	cardStatements, err := csu.csr.GetCardStatementsByMonth(request.UserId, request.Year, request.Month)
	if err != nil {
		return nil, err
	}
	
	// Convert to response format
	responses := make([]model.CardStatementResponse, len(cardStatements))
	for i, cardStatement := range cardStatements {
		responses[i] = cardStatement.ToResponse()
	}
	
	return responses, nil
}
