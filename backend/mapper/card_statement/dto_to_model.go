package card_statement

import (
	"monelog/dto"
	"monelog/model"
)

// DTOToModel はDTOからモデルへの変換を担当するモジュール
type DTOToModel struct{}

// FromSummary サマリーDTOからモデルへの変換
func (d *DTOToModel) FromSummary(css *dto.CardStatementSummary, userId uint, year int, month int) model.CardStatement {
	return model.CardStatement{
		Type:              css.Type,
		StatementNo:       css.StatementNo,
		CardType:          css.CardType,
		Description:       css.Description,
		UseDate:           css.UseDate,
		PaymentDate:       css.PaymentDate,
		PaymentMonth:      css.PaymentMonth,
		Amount:            css.Amount,
		TotalChargeAmount: css.TotalChargeAmount,
		ChargeAmount:      css.ChargeAmount,
		RemainingBalance:  css.RemainingBalance,
		PaymentCount:      css.PaymentCount,
		InstallmentCount:  css.InstallmentCount,
		AnnualRate:        css.AnnualRate,
		MonthlyRate:       css.MonthlyRate,
		UserId:            userId,
		Year:              year,
		Month:             month,
	}
}

// NewDTOToModel は新しいDTOToModelインスタンスを作成する
func NewDTOToModel() *DTOToModel {
	return &DTOToModel{}
}