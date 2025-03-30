package repository

import (
	"monelog/model"

	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserByEmail(email string) (*model.User, error)
	CreateUser(user *model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

func (ur *userRepository) GetUserByEmail(email string) (*model.User, error) {
	user := &model.User{}
	if err := ur.db.Where("email = ?", email).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *userRepository) CreateUser(user *model.User) error {
	return ur.db.Create(user).Error
}