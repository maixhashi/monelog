package model

import "time"

// User はデータベースのユーザーモデル
type User struct {
	ID        uint      `json:"id" gorm:"primaryKey" example:"1"`
	Email     string    `json:"email" gorm:"unique" example:"user@example.com"`
	Password  string    `json:"password" example:"password123"`
	CreatedAt time.Time `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2023-01-01T00:00:00Z"`
}

// UserResponse はクライアントに返すユーザー情報
type UserResponse struct {
	ID    uint   `json:"id" example:"1"`
	Email string `json:"email" example:"user@example.com"`
}

// UserLoginRequest はログインリクエスト用の構造体
type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email" example:"user@example.com"`
	Password string `json:"password" validate:"required" example:"password123"`
}

// UserSignupRequest はサインアップリクエスト用の構造体
type UserSignupRequest struct {
	Email    string `json:"email" validate:"required,email" example:"user@example.com"`
	Password string `json:"password" validate:"required" example:"password123"`
}

// ToUser はUserSignupRequestからUserへの変換メソッド
func (r *UserSignupRequest) ToUser() User {
	return User{
		Email:    r.Email,
		Password: r.Password,
	}
}

// CsrfTokenResponse はCSRFトークンのレスポンス
type CsrfTokenResponse struct {
	CsrfToken string `json:"csrf_token" example:"token-string-here"`
}

// ToUserResponse はUserからUserResponseへの変換メソッド
func (u *User) ToUserResponse() UserResponse {
	return UserResponse{
		ID:    u.ID,
		Email: u.Email,
	}
}