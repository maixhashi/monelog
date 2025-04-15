package model

import "time"

// CSVHistory データベースのCSV履歴モデル
type CSVHistory struct {
	ID        uint      `json:"id" gorm:"primaryKey" example:"1"`
	FileName  string    `json:"file_name" gorm:"not null" example:"rakuten_202301.csv"`
	CardType  string    `json:"card_type" gorm:"not null" example:"rakuten"`
	FileData  []byte    `json:"file_data" gorm:"type:bytea;not null"`
	CreatedAt time.Time `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2023-01-01T00:00:00Z"`
	User      User      `json:"-" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId    uint      `json:"user_id" gorm:"not null" example:"1"`
}

// CSVHistoryResponse CSV履歴のレスポンス
type CSVHistoryResponse struct {
	ID        uint      `json:"id" example:"1"`
	FileName  string    `json:"file_name" example:"rakuten_202301.csv"`
	CardType  string    `json:"card_type" example:"rakuten"`
	CreatedAt time.Time `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2023-01-01T00:00:00Z"`
}

// CSVHistoryDetailResponse CSV履歴の詳細レスポンス（ファイルデータを含む）
type CSVHistoryDetailResponse struct {
	ID        uint      `json:"id" example:"1"`
	FileName  string    `json:"file_name" example:"rakuten_202301.csv"`
	CardType  string    `json:"card_type" example:"rakuten"`
	FileData  []byte    `json:"file_data"`
	CreatedAt time.Time `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2023-01-01T00:00:00Z"`
}

// CSVHistorySaveRequest CSV履歴保存リクエスト
type CSVHistorySaveRequest struct {
	FileName string `json:"file_name" validate:"required" example:"rakuten_202301.csv"`
	CardType string `json:"card_type" validate:"required" example:"rakuten"`
	UserId   uint   `json:"-"`
}

// ToResponse CSVHistoryからCSVHistoryResponseへの変換メソッド
func (ch *CSVHistory) ToResponse() CSVHistoryResponse {
	return CSVHistoryResponse{
		ID:        ch.ID,
		FileName:  ch.FileName,
		CardType:  ch.CardType,
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
		CreatedAt: ch.CreatedAt,
		UpdatedAt: ch.UpdatedAt,
	}
}
