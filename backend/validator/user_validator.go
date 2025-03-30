package validator

import (
	"fmt"
	"monelog/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type IUserValidator interface {
	UserValidate(user model.User) error
}

type userValidator struct{}

func NewUserValidator() IUserValidator {
	return &userValidator{}
}

func (uv *userValidator) UserValidate(user model.User) error {
	return validation.ValidateStruct(&user,
		validation.Field(
			&user.Email,
			validation.Required.Error("メールアドレスは必須です"),
			validation.RuneLength(model.UserEmailMinLength, model.UserEmailMaxLength).Error(
				fmt.Sprintf("メールアドレスは%d文字から%d文字の間である必要があります", 
				model.UserEmailMinLength, model.UserEmailMaxLength)),
			is.Email.Error("有効なメールアドレス形式ではありません"),
		),
		validation.Field(
			&user.Password,
			validation.Required.Error("パスワードは必須です"),
			validation.RuneLength(model.UserPasswordMinLength, model.UserPasswordMaxLength).Error(
				fmt.Sprintf("パスワードは%d文字から%d文字の間である必要があります", 
				model.UserPasswordMinLength, model.UserPasswordMaxLength)),
		),
	)
}