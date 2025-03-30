package user_test

import (
	"monelog/model"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_CreateUser(t *testing.T) {
	setupUserRepositoryTest()

	t.Run("ユーザーを正常に作成できる", func(t *testing.T) {
		user := testUserData
		err := userRepo.CreateUser(&user)
		
		assert.NoError(t, err)
		assert.NotZero(t, user.ID)
		assert.NotZero(t, user.CreatedAt)
		assert.NotZero(t, user.UpdatedAt)
	})

	t.Run("同じメールアドレスのユーザーは作成できない", func(t *testing.T) {
		firstUser := testUserData
		_ = userRepo.CreateUser(&firstUser)
		
		duplicateUser := model.User{
			Email:    testUserData.Email,
			Password: "different_password",
		}
		
		err := userRepo.CreateUser(&duplicateUser)
		assert.Error(t, err)
	})
}
