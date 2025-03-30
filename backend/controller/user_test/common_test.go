package user_test

import (
	"encoding/json"
	"fmt"
	"monelog/model"
	"net/http/httptest"
	"testing"
	"time"
)

// テスト用の一意なメールアドレスを生成
func generateTestEmail() string {
	timestamp := time.Now().UnixNano() % 10000 // 短い数値に制限
	return fmt.Sprintf("test%d@example.com", timestamp)
}

// テスト用ユーザーを作成
func createTestUser(t *testing.T) (string, string) {
	email := generateTestEmail()
	password := "password123"
	
	// ユーザー登録
	signupReq := model.UserSignupRequest{
		Email:    email,
		Password: password,
	}
	
	_, err := userUsecase.SignUp(signupReq)
	if err != nil {
		t.Fatalf("テストユーザーの登録に失敗しました: %v", err)
	}
	
	return email, password
}

// レスポンスボディをパースするヘルパー関数
func parseUserResponse(t *testing.T, responseBody []byte) model.UserResponse {
	var response model.UserResponse
	err := json.Unmarshal(responseBody, &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}
	return response
}

// Cookieの存在と値を確認するヘルパー関数
func checkCookie(t *testing.T, rec *httptest.ResponseRecorder, name string, shouldExist bool) {
	cookies := rec.Result().Cookies()
	var found bool
	var cookieValue string
	
	for _, cookie := range cookies {
		if cookie.Name == name {
			found = true
			cookieValue = cookie.Value
			break
		}
	}
	
	if shouldExist && !found {
		t.Errorf("Cookie '%s' not found in response", name)
	} else if !shouldExist && found {
		t.Errorf("Cookie '%s' should not exist in response", name)
	}
	
	if shouldExist && found && cookieValue == "" {
		t.Errorf("Cookie '%s' exists but has empty value", name)
	}
}
