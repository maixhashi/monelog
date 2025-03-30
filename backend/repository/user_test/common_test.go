package user_test

import (
	"monelog/model"
)

// テストヘルパー関数
func createTestUser() (*model.User, error) {
	user := testUserData
	err := userRepo.CreateUser(&user)
	return &user, err
}

// テストデータジェネレーター
func generateUniqueTestUser(email string) model.User {
	return model.User{
		Email:    email,
		Password: "testpassword123",
	}
}

// テストアサーション用のヘルパー
func validateUserFields(user *model.User) bool {
	return user.ID != 0 &&
		user.Email != "" &&
		user.Password != "" &&
		!user.CreatedAt.IsZero() &&
		!user.UpdatedAt.IsZero()
}
