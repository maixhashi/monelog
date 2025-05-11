package validator

import (
	"monelog/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ICardStatementValidator interface {
	ValidateCardStatementRequest(request model.CardStatementRequest) error
	ValidateCardStatementPreviewRequest(request model.CardStatementPreviewRequest) error
	ValidateCardStatementSaveRequest(request model.CardStatementSaveRequest) error
	ValidateCardStatementByMonthRequest(request model.CardStatementByMonthRequest) error
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
		validation.Field(
			&request.Year,
			validation.Required.Error("year is required"),
		),
		validation.Field(
			&request.Month,
			validation.Required.Error("month is required"),
			validation.Min(1).Error("month must be between 1 and 12"),
			validation.Max(12).Error("month must be between 1 and 12"),
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
		// 年月のバリデーションを追加（任意項目）
		validation.Field(
			&request.Year,
			validation.Min(0).Error("year must be positive or zero"),
		),
		validation.Field(
			&request.Month,
			validation.Min(0).Error("month must be positive or zero"),
			validation.Max(12).Error("month must be between 0 and 12"),
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
		validation.Field(
			&request.Year,
			validation.Required.Error("year is required"),
		),
		validation.Field(
			&request.Month,
			validation.Required.Error("month is required"),
			validation.Min(1).Error("month must be between 1 and 12"),
			validation.Max(12).Error("month must be between 1 and 12"),
		),
	)
}
func (csv *cardStatementValidator) ValidateCardStatementByMonthRequest(request model.CardStatementByMonthRequest) error {
	return validation.ValidateStruct(&request,
		validation.Field(
			&request.Year,
			validation.Required.Error("year is required"),
		),
		validation.Field(
			&request.Month,
			validation.Required.Error("month is required"),
			validation.Min(1).Error("month must be between 1 and 12"),
			validation.Max(12).Error("month must be between 1 and 12"),
		),
	)
}