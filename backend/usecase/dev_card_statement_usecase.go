package usecase

import (
	"errors"
	"monelog/model"
	"monelog/repository"
	"monelog/validator"
)

type IDevCardStatementUsecase interface {
	DeleteAllCardStatements(request model.DevCardStatementRequest) (model.DevCardStatementResponse, error)
}

type devCardStatementUsecase struct {
	dcsr repository.IDevCardStatementRepository
	dcsv validator.IDevCardStatementValidator
}

func NewDevCardStatementUsecase(
	dcsr repository.IDevCardStatementRepository,
	dcsv validator.IDevCardStatementValidator,
) IDevCardStatementUsecase {
	return &devCardStatementUsecase{dcsr, dcsv}
}

func (dcsu *devCardStatementUsecase) DeleteAllCardStatements(request model.DevCardStatementRequest) (model.DevCardStatementResponse, error) {
	// 開発環境かどうかチェック
	if !dcsu.dcsv.IsDevEnvironment() {
		return model.DevCardStatementResponse{}, errors.New("this operation is only allowed in development environment")
	}

	// リクエストのバリデーション（空になりました）
	if err := dcsu.dcsv.ValidateDevCardStatementRequest(request); err != nil {
		return model.DevCardStatementResponse{}, err
	}

	// 全レコード削除
	deletedRows, err := dcsu.dcsr.DeleteAllCardStatements()
	if err != nil {
		return model.DevCardStatementResponse{}, err
	}

	// レスポンス作成
	response := model.DevCardStatementResponse{
		Message:     "All card statements deleted successfully",
		DeletedRows: deletedRows,
	}

	return response, nil
}
