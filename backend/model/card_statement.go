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
	Year              int       `json:"year" gorm:"not null" example:"2023"`
	Month             int       `json:"month" gorm:"not null" example:"4"`
}