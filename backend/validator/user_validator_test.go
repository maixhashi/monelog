package validator

import (
	"monelog/model"
	"testing"
)

func TestUserValidate(t *testing.T) {
	validator := NewUserValidator()

	testCases := []struct {
		name     string
		user     model.User
		hasError bool
	}{
		{
			name: "Valid user with valid email and password",
			user: model.User{
				Email:    generateEmailWithLocalPartLength(model.UserEmailMaxLength - len("@example.com")), // ドメイン部分の長さを正確に計算して最大長ぎりぎりのメールアドレスを生成
				Password: "password123",
			},
			hasError: false,
		},
		{
			name: "Empty email",
			user: model.User{
				Email:    "",
				Password: "password123",
			},
			hasError: true,
		},
		{
			name: "Invalid email format",
			user: model.User{
				Email:    "not-an-email",
				Password: "password123",
			},
			hasError: true,
		},
		{
			name: "Email with exact min length",
			user: model.User{
				Email:    generateEmailWithLocalPartLength(model.UserEmailMinLength),
				Password: "password123",
			},
			hasError: false,
		},
		{
			name: "Email with exact max length",
			user: model.User{
				Email:    generateEmailWithLocalPartLength(model.UserEmailMaxLength - len("@example.com")), // 正確にドメイン部分の長さを計算
				Password: "password123",
			},
			hasError: false,
		},
		{
			name: "Email too long",
			user: model.User{
				Email:    generateEmailWithLocalPartLength(model.UserEmailMaxLength - 10), // +1 で最大長を超える
				Password: "password123",
			},
			hasError: true,
		},
		{
			name: "Empty password",
			user: model.User{
				Email:    "test@example.com",
				Password: "",
			},
			hasError: true,
		},
		{
			name: "Password too short",
			user: model.User{
				Email:    "test@example.com",
				Password: generateStringWithLength(model.UserPasswordMinLength - 1),
			},
			hasError: true,
		},
		{
			name: "Password too long",
			user: model.User{
				Email:    "test@example.com",
				Password: generateStringWithLength(model.UserPasswordMaxLength + 1),
			},
			hasError: true,
		},
		{
			name: "Password with exact min length",
			user: model.User{
				Email:    "test@example.com",
				Password: generateStringWithLength(model.UserPasswordMinLength),
			},
			hasError: false,
		},
		{
			name: "Password with exact max length",
			user: model.User{
				Email:    "test@example.com",
				Password: generateStringWithLength(model.UserPasswordMaxLength),
			},
			hasError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validator.UserValidate(tc.user)
			if (err != nil) != tc.hasError {
				t.Errorf("UserValidate() error = %v, want error: %v", err, tc.hasError)
			}
		})
	}
}

// 指定された長さの文字列を生成するヘルパー関数
func generateStringWithLength(length int) string {
    str := ""
    for i := 0; i < length; i++ {
        str += "a"
    }
    return str
}

// 指定された長さのローカルパート（@より前の部分）を持つメールアドレスを生成する
func generateEmailWithLocalPartLength(localPartLength int) string {
    localPart := ""
    for i := 0; i < localPartLength; i++ {
        localPart += "a"
    }
    return localPart + "@example.com"
}