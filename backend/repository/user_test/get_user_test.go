package user_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_GetUserByEmail(t *testing.T) {
	setupUserRepositoryTest()

	t.Run("存在するユーザーを取得できる", func(t *testing.T) {
		user := testUserData
		err := userRepo.CreateUser(&user)
		assert.NoError(t, err)
		
		foundUser, err := userRepo.GetUserByEmail(user.Email)
		
		assert.NoError(t, err)
		assert.Equal(t, user.ID, foundUser.ID)
		assert.Equal(t, user.Email, foundUser.Email)
		assert.Equal(t, user.Password, foundUser.Password)
	})

	t.Run("存在しないユーザーを取得するとエラーになる", func(t *testing.T) {
		nonExistentEmail := "nonexistent@example.com"
		
		_, err := userRepo.GetUserByEmail(nonExistentEmail)
		
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "record not found")
	})
}
