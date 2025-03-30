package user_test

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"os"
	"testing"
)

// 一意のメールアドレスを生成
func generateUniqueEmail() string {
	// UUIDを生成（ほぼ確実に一意）
	uuid := uuid.New().String()
	// UUIDの最初の8文字だけを使用
	shortUUID := uuid[:8]
	
	return fmt.Sprintf("test-%s@example.com", shortUUID)
}

// JWTトークンを検証するヘルパー関数
func validateJWTToken(t *testing.T, tokenString string) bool {
	if tokenString == "" {
		t.Error("空のトークンが返されました")
		return false
	}

	// トークンが有効なJWTかどうか検証
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil || !token.Valid {
		t.Errorf("生成されたトークンが無効です: %v", err)
		return false
	}

	// クレームの検証
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		t.Error("トークンクレームの取得に失敗しました")
		return false
	}

	userId, exists := claims["user_id"]
	if !exists || userId == 0 {
		t.Errorf("トークンにuser_idが含まれていないか、無効な値です: %v", userId)
		return false
	}

	return true
}
