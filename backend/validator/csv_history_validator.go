package validator

import (
	"monelog/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ICSVHistoryValidator interface {
	ValidateCSVHistorySaveRequest(request model.CSVHistorySaveRequest) error
}

type csvHistoryValidator struct{}

func NewCSVHistoryValidator() ICSVHistoryValidator {
	return &csvHistoryValidator{}
}

func (chv *csvHistoryValidator) ValidateCSVHistorySaveRequest(request model.CSVHistorySaveRequest) error {
	return validation.ValidateStruct(&request,
		validation.Field(
			&request.FileName,
			validation.Required.Error("file_name is required"),
		),
		validation.Field(
			&request.CardType,
			validation.Required.Error("card_type is required"),
			validation.In("rakuten", "mufg", "epos").Error("card_type must be one of: rakuten, mufg, epos"),
		),
	)
}
