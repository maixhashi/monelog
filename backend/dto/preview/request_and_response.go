package preview

// Request CSVプレビュー用リクエスト
type Request struct {
	CardType string `json:"card_type" validate:"required" example:"rakuten"`
	Year     int    `json:"year" example:"2023"`
	Month    int    `json:"month" example:"4"`
	UserId   uint   `json:"-"`
}