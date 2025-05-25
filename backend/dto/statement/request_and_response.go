package statement

import "time"

// Request カード明細のCSVアップロードリクエスト
type Request struct {
	CardType string `json:"card_type" validate:"required" example:"rakuten"` // rakuten, mufg, epos
	Year     int    `json:"year" validate:"required" example:"2023"`
	Month    int    `json:"month" validate:"required,min=1,max=12" example:"4"`
	UserId   uint   `json:"-"` // クライアントからは送信されず、JWTから取得
}

// Response カード明細のレスポンス
type Response struct {
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
	Year              int       `json:"year" example:"2023"`
	Month             int       `json:"month" example:"4"`
}

// ByMonthRequest 支払月ごとのカード明細取得リクエスト
type ByMonthRequest struct {
	Year   int  `query:"year" validate:"required" example:"2023"`
	Month  int  `query:"month" validate:"required,min=1,max=12" example:"4"`
	UserId uint `json:"-"` // JWTから取得
}