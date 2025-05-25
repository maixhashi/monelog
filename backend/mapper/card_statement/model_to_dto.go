package card_statement

import (
	"monelog/dto"
	"monelog/model"
)

// ModelToDTO はモデルからDTOへの変換を担当するモジュール
type ModelToDTO struct{}

// ToResponse モデルからレスポンスDTOへの変換
func (m *ModelToDTO) ToResponse(cs *model.CardStatement) dto.CardStatementResponse {
	return dto.CardStatementResponse{
		ID:                cs.ID,
		Type:              cs.Type,
		StatementNo:       cs.StatementNo,
		CardType:          cs.CardType,
		Description:       cs.Description,
		UseDate:           cs.UseDate,
		PaymentDate:       cs.PaymentDate,
		PaymentMonth:      cs.PaymentMonth,
		Amount:            cs.Amount,
		TotalChargeAmount: cs.TotalChargeAmount,
		ChargeAmount:      cs.ChargeAmount,
		RemainingBalance:  cs.RemainingBalance,
		PaymentCount:      cs.PaymentCount,
		InstallmentCount:  cs.InstallmentCount,
		AnnualRate:        cs.AnnualRate,
		MonthlyRate:       cs.MonthlyRate,
		CreatedAt:         cs.CreatedAt,
		UpdatedAt:         cs.UpdatedAt,
		Year:              cs.Year,
		Month:             cs.Month,
	}
}

// ToResponseList モデルのスライスからレスポンスDTOのスライスへの変換
func (m *ModelToDTO) ToResponseList(statements []model.CardStatement) []dto.CardStatementResponse {
	responses := make([]dto.CardStatementResponse, len(statements))
	for i, statement := range statements {
		responses[i] = m.ToResponse(&statement)
	}
	return responses
}

// NewModelToDTO は新しいModelToDTOインスタンスを作成する
func NewModelToDTO() *ModelToDTO {
	return &ModelToDTO{}
}