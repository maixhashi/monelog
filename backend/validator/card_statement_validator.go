package validator

import (
	"monelog/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ICardStatementValidator interface {
	ValidateCardStatementRequest(request model.CardStatementRequest) error
	ValidateCardStatementPreviewRequest(request model.CardStatementPreviewRequest) error
	ValidateCardStatementSaveRequest(request model.CardStatementSaveRequest) error
}

type cardStatementValidator struct{}

func NewCardStatementValidator() ICardStatementValidator {
	return &cardStatementValidator{}
}

func (csv *cardStatementValidator) ValidateCardStatementRequest(request model.CardStatementRequest) error {
	return validation.ValidateStruct(&request,
		validation.Field(
			&request.CardType,
			validation.Required.Error("card_type is required"),
			validation.In("rakuten", "mufg", "epos").Error("card_type must be one of: rakuten, mufg, epos"),
		),
	)
}

func (csv *cardStatementValidator) ValidateCardStatementPreviewRequest(request model.CardStatementPreviewRequest) error {
	return validation.ValidateStruct(&request,
		validation.Field(
			&request.CardType,
			validation.Required.Error("card_type is required"),
			validation.In("rakuten", "mufg", "epos").Error("card_type must be one of: rakuten, mufg, epos"),
		),
	)
}

func (csv *cardStatementValidator) ValidateCardStatementSaveRequest(request model.CardStatementSaveRequest) error {
	return validation.ValidateStruct(&request,
		validation.Field(
			&request.CardType,
			validation.Required.Error("card_type is required"),
			validation.In("rakuten", "mufg", "epos").Error("card_type must be one of: rakuten, mufg, epos"),
		),
		validation.Field(
			&request.CardStatements,
			validation.Required.Error("card_statements is required"),
		),
	)
}
