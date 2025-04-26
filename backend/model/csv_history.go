package model

import "time"

// CSVHistory データベースのCSV履歴モデル
type CSVHistory struct {
	ID          uint      `json:"id" gorm:"primaryKey" example:"1"`
	FileName    string    `json:"file_name" gorm:"not null" example:"rakuten_202301.csv"`
	CardType    string    `json:"card_type" gorm:"not null" example:"rakuten"`
	FileData    []byte    `json:"file_data" gorm:"type:bytea;not null"`
	Year        int       `json:"year" gorm:"not null" example:"2023"` // 追加: 年
	Month       int       `json:"month" gorm:"not null" example:"1"`   // 追加: 月
	CreatedAt   time.Time `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt   time.Time `json:"updated_at" example:"2023-01-01T00:00:00Z"`
	User        User      `json:"-" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId      uint      `json:"user_id" gorm:"not null" example:"1"`
}

// CSVHistoryResponse CSV履歴のレスポンス
type CSVHistoryResponse struct {
	ID        uint      `json:"id" example:"1"`
	FileName  string    `json:"file_name" example:"rakuten_202301.csv"`
	CardType  string    `json:"card_type" example:"rakuten"`
	Year      int       `json:"year" example:"2023"` // 追加: 年
	Month     int       `json:"month" example:"1"`   // 追加: 月
	CreatedAt time.Time `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2023-01-01T00:00:00Z"`
}

// CSVHistoryDetailResponse CSV履歴の詳細レスポンス（ファイルデータを含む）
type CSVHistoryDetailResponse struct {
	ID        uint      `json:"id" example:"1"`
	FileName  string    `json:"file_name" example:"rakuten_202301.csv"`
	CardType  string    `json:"card_type" example:"rakuten"`
	FileData  []byte    `json:"file_data"`
	Year      int       `json:"year" example:"2023"` // 追加: 年
	Month     int       `json:"month" example:"1"`   // 追加: 月
	CreatedAt time.Time `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2023-01-01T00:00:00Z"`
}

// CSVHistorySaveRequest CSV履歴保存リクエスト
type CSVHistorySaveRequest struct {
	FileName string `json:"file_name" validate:"required" example:"rakuten_202301.csv"`
	CardType string `json:"card_type" validate:"required" example:"rakuten"`
	Year     int    `json:"year" validate:"required" example:"2023"` // 追加: 年
	Month    int    `json:"month" validate:"required,min=1,max=12" example:"1"` // 追加: 月
	UserId   uint   `json:"-"`
}

// ToResponse CSVHistoryからCSVHistoryResponseへの変換メソッド
func (ch *CSVHistory) ToResponse() CSVHistoryResponse {
	return CSVHistoryResponse{
		ID:        ch.ID,
		FileName:  ch.FileName,
		CardType:  ch.CardType,
		Year:      ch.Year,      // 追加
		Month:     ch.Month,     // 追加
		CreatedAt: ch.CreatedAt,
		UpdatedAt: ch.UpdatedAt,
	}
}

// ToDetailResponse CSVHistoryからCSVHistoryDetailResponseへの変換メソッド
func (ch *CSVHistory) ToDetailResponse() CSVHistoryDetailResponse {
	return CSVHistoryDetailResponse{
		ID:        ch.ID,
		FileName:  ch.FileName,
		CardType:  ch.CardType,
		FileData:  ch.FileData,
		Year:      ch.Year,      // 追加
		Month:     ch.Month,     // 追加
		CreatedAt: ch.CreatedAt,
		UpdatedAt: ch.UpdatedAt,
	}
}
