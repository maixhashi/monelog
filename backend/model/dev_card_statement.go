package model

// DevCardStatementRequest 開発環境限定のカード明細削除リクエスト
// dev_tokenを削除
type DevCardStatementRequest struct {
    // フィールドを空にする（または必要に応じて他のフィールドを追加）
}

// DevCardStatementResponse 開発環境限定のカード明細削除レスポンス
type DevCardStatementResponse struct {
	Message     string `json:"message" example:"All card statements deleted successfully"`
	DeletedRows int64  `json:"deleted_rows" example:"42"`
}
