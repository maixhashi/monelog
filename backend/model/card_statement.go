package model

import "time"

// CardStatementTitleMaxLength カード明細のタイトル最大長
const CardStatementTitleMaxLength = 100

// CardStatement データベースのカード明細モデル
type CardStatement struct {
	ID                uint      `json:"id" gorm:"primaryKey" example:"1"`
	Type              string    `json:"type" gorm:"not null" example:"発生"` // 発生 or 分割
	StatementNo       int       `json:"statement_no" gorm:"not null" example:"1"`
	CardType          string    `json:"card_type" gorm:"not null" example:"楽天カード"`
	Description       string    `json:"description" gorm:"not null" example:"Amazon.co.jp"`
	UseDate           string    `json:"use_date" gorm:"not null" example:"2023/01/01"`
	PaymentDate       string    `json:"payment_date" gorm:"not null" example:"2023/02/27"`
	PaymentMonth      string    `json:"payment_month" gorm:"not null" example:"2023年02月"`
	Amount            int       `json:"amount" gorm:"not null" example:"10000"`
	TotalChargeAmount int       `json:"total_charge_amount" gorm:"not null" example:"10000"`
	ChargeAmount      int       `json:"charge_amount" gorm:"not null" example:"0"`
	RemainingBalance  int       `json:"remaining_balance" gorm:"not null" example:"10000"`
	PaymentCount      int       `json:"payment_count" gorm:"not null" example:"0"`
	InstallmentCount  int       `json:"installment_count" gorm:"not null" example:"1"`
	AnnualRate        float64   `json:"annual_rate" gorm:"not null" example:"0.0"`
	MonthlyRate       float64   `json:"monthly_rate" gorm:"not null" example:"0.0"`
	CreatedAt         time.Time `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt         time.Time `json:"updated_at" example:"2023-01-01T00:00:00Z"`
	User              User      `json:"-" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId            uint      `json:"user_id" gorm:"not null" example:"1"`
}

// CardStatementRequest カード明細のCSVアップロードリクエスト
type CardStatementRequest struct {
	CardType string `json:"card_type" validate:"required" example:"rakuten"` // rakuten, mufg, epos
	UserId   uint   `json:"-"`                                               // クライアントからは送信されず、JWTから取得
}

// CardStatementResponse カード明細のレスポンス
type CardStatementResponse struct {
	ID                uint      `json:"id" example:"1"`
	Type              string    `json:"type" example:"発生"`
	StatementNo       int       `json:"statement_no" example:"1"`
	CardType          string    `json:"card_type" example:"楽天カード"`
	Description       string    `json:"description" example:"Amazon.co.jp"`
	UseDate           string    `json:"use_date" example:"2023/01/01"`
	PaymentDate       string    `json:"payment_date" example:"2023/02/27"`
	PaymentMonth      string    `json:"payment_month" example:"2023年02月"`
	Amount            int       `json:"amount" example:"10000"`
	TotalChargeAmount int       `json:"total_charge_amount" example:"10000"`
	ChargeAmount      int       `json:"charge_amount" example:"0"`
	RemainingBalance  int       `json:"remaining_balance" example:"10000"`
	PaymentCount      int       `json:"payment_count" example:"0"`
	InstallmentCount  int       `json:"installment_count" example:"1"`
	AnnualRate        float64   `json:"annual_rate" example:"0.0"`
	MonthlyRate       float64   `json:"monthly_rate" example:"0.0"`
	CreatedAt         time.Time `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt         time.Time `json:"updated_at" example:"2023-01-01T00:00:00Z"`
}

// CardStatementSummary CSVから解析した明細データ
type CardStatementSummary struct {
	Type              string  `json:"type"`
	StatementNo       int     `json:"statement_no"`
	CardType          string  `json:"card_type"`
	Description       string  `json:"description"`
	UseDate           string  `json:"use_date"`
	PaymentDate       string  `json:"payment_date"`
	PaymentMonth      string  `json:"payment_month"`
	Amount            int     `json:"amount"`
	TotalChargeAmount int     `json:"total_charge_amount"`
	ChargeAmount      int     `json:"charge_amount"`
	RemainingBalance  int     `json:"remaining_balance"`
	PaymentCount      int     `json:"payment_count"`
	InstallmentCount  int     `json:"installment_count"`
	AnnualRate        float64 `json:"annual_rate"`
	MonthlyRate       float64 `json:"monthly_rate"`
}

// ToResponse CardStatementからCardStatementResponseへの変換メソッド
func (cs *CardStatement) ToResponse() CardStatementResponse {
	return CardStatementResponse{
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
	}
}

// ToModel CardStatementSummaryからCardStatementへの変換メソッド
func (css *CardStatementSummary) ToModel(userId uint) CardStatement {
	return CardStatement{
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
	}
}

// CardStatementPreviewRequest CSVプレビュー用リクエスト
type CardStatementPreviewRequest struct {
	CardType string `json:"card_type" validate:"required" example:"rakuten"`
	UserId   uint   `json:"-"`
}

// CardStatementSaveRequest 一時データを保存するリクエスト
type CardStatementSaveRequest struct {
	CardStatements []CardStatementSummary `json:"card_statements" validate:"required"`
	CardType       string                 `json:"card_type" validate:"required" example:"rakuten"`
	Year           int                    `json:"year" validate:"required" example:"2023"`
	Month          int                    `json:"month" validate:"required,min=1,max=12" example:"4"`
	UserId         uint                   `json:"-"`
}
// CardStatementByMonthRequest 支払月ごとのカード明細取得リクエスト
type CardStatementByMonthRequest struct {
	Year   int  `query:"year" validate:"required" example:"2023"`
	Month  int  `query:"month" validate:"required,min=1,max=12" example:"4"`
	UserId uint `json:"-"` // JWTから取得
}