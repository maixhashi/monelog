package summary

// Summary CSVから解析した明細データ
type Summary struct {
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

// SaveRequest 一時データを保存するリクエスト
type SaveRequest struct {
	CardStatements []Summary `json:"card_statements" validate:"required"`
	CardType       string    `json:"card_type" validate:"required" example:"rakuten"`
	Year           int       `json:"year" validate:"required" example:"2023"`
	Month          int       `json:"month" validate:"required,min=1,max=12" example:"4"`
	UserId         uint      `json:"-"`
}