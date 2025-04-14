package validator

import (
	"monelog/model"
	"os"
)

type IDevCardStatementValidator interface {
	ValidateDevCardStatementRequest(request model.DevCardStatementRequest) error
	IsDevEnvironment() bool
}

type devCardStatementValidator struct{}

func NewDevCardStatementValidator() IDevCardStatementValidator {
	return &devCardStatementValidator{}
}

func (dcsv *devCardStatementValidator) ValidateDevCardStatementRequest(request model.DevCardStatementRequest) error {
	// dev_tokenのバリデーションを削除
	// 必要に応じて他のバリデーションを追加
	return nil
}

func (dcsv *devCardStatementValidator) IsDevEnvironment() bool {
	env := os.Getenv("APP_ENV")
	return env == "development" || env == "dev" || env == ""
}
