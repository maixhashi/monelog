package model

import "time"

// Task データベースのタスクモデル
type Task struct {
	ID        uint      `json:"id" gorm:"primaryKey" example:"1"`
	Title     string    `json:"title" gorm:"not null" example:"買い物に行く"`
	CreatedAt time.Time `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2023-01-01T00:00:00Z"`
	User      User      `json:"-" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId    uint      `json:"user_id" gorm:"not null" example:"1"`
}

// TaskRequest タスク作成・更新リクエスト
type TaskRequest struct {
	Title  string `json:"title" validate:"required,max=100" example:"買い物に行く"`
	UserId uint   `json:"-"` // クライアントからは送信されず、JWTから取得
}

// TaskResponse タスクのレスポンス
type TaskResponse struct {
	ID        uint      `json:"id" example:"1"`
	Title     string    `json:"title" example:"買い物に行く"`
	CreatedAt time.Time `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2023-01-01T00:00:00Z"`
}

// ToResponse TaskからTaskResponseへの変換メソッド
func (t *Task) ToResponse() TaskResponse {
	return TaskResponse{
		ID:        t.ID,
		Title:     t.Title,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}

// ToModel TaskRequestからTaskへの変換メソッド
func (tr *TaskRequest) ToModel() Task {
	return Task{
		Title:  tr.Title,
		UserId: tr.UserId,
	}
}