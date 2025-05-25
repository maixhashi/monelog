package csv_history

import "time"

// SaveRequest CSV履歴保存リクエスト
type SaveRequest struct {
	FileName string `json:"file_name" validate:"required" example:"rakuten_202301.csv"`
	CardType string `json:"card_type" validate:"required" example:"rakuten"`
	Year     int    `json:"year" validate:"required" example:"2023"`
	Month    int    `json:"month" validate:"required,min=1,max=12" example:"1"`
	UserId   uint   `json:"-"`
}

// Response CSV履歴のレスポンス
type Response struct {
	ID        uint      `json:"id" example:"1"`
	FileName  string    `json:"file_name" example:"rakuten_202301.csv"`
	CardType  string    `json:"card_type" example:"rakuten"`
	Year      int       `json:"year" example:"2023"`
	Month     int       `json:"month" example:"1"`
	CreatedAt time.Time `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2023-01-01T00:00:00Z"`
}

// DetailResponse CSV履歴の詳細レスポンス（ファイルデータを含む）
type DetailResponse struct {
	ID        uint      `json:"id" example:"1"`
	FileName  string    `json:"file_name" example:"rakuten_202301.csv"`
	CardType  string    `json:"card_type" example:"rakuten"`
	FileData  []byte    `json:"file_data"`
	Year      int       `json:"year" example:"2023"`
	Month     int       `json:"month" example:"1"`
	CreatedAt time.Time `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2023-01-01T00:00:00Z"`
}