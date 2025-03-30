package user_test

import (
	"monelog/model"
	"net/http"
	"testing"
)

func TestUserController_LogIn(t *testing.T) {
	setupUserControllerTest()

	// テスト用ユーザーを事前に登録
	validEmail := "logintest@example.com"
	validPassword := "password123"
	userUsecase.SignUp(model.UserSignupRequest{
		Email:    validEmail,
		Password: validPassword,
	})

	t.Run("正常系", func(t *testing.T) {
		t.Run("有効な認証情報でログインに成功する", func(t *testing.T) {
			// テストリクエストの準備
			reqBody := `{"email":"` + validEmail + `","password":"` + validPassword + `"}`

			// テスト実行
			_, c, rec := setupEchoContextWithBody(http.MethodPost, "/login", reqBody)
			err := userController.LogIn(c)

			// 検証
			if err != nil {
				t.Errorf("LogIn() error = %v", err)
			}

			if rec.Code != http.StatusOK {
				t.Errorf("LogIn() status code = %d, want %d", rec.Code, http.StatusOK)
			}

			// Cookieが設定されていることを確認
			checkCookie(t, rec, "token", true)
		})
	})

	t.Run("異常系", func(t *testing.T) {
		t.Run("無効なJSONリクエストでバッドリクエストを返す", func(t *testing.T) {
			// 無効なJSON
			invalidJSON := `{"email": "invalid@example.com", "password": Invalid JSON`

			// テスト実行
			_, c, rec := setupEchoContextWithBody(http.MethodPost, "/login", invalidJSON)
			err := userController.LogIn(c)

			if err != nil {
				t.Errorf("LogIn() unexpected error: %v", err)
			}

			if rec.Code != http.StatusBadRequest {
				t.Errorf("LogIn() with invalid JSON status code = %d, want %d",
					rec.Code, http.StatusBadRequest)
			}
		})

		t.Run("存在しないユーザーでログインするとエラーを返す", func(t *testing.T) {
			// 存在しないユーザー
			nonExistentUserReqBody := `{"email":"nonexistent@example.com","password":"password123"}`

			// テスト実行
			_, c, rec := setupEchoContextWithBody(http.MethodPost, "/login", nonExistentUserReqBody)
			err := userController.LogIn(c)

			if err != nil {
				t.Errorf("LogIn() unexpected error: %v", err)
			}

			if rec.Code != http.StatusInternalServerError {
				t.Errorf("LogIn() with non-existent user status code = %d, want %d",
					rec.Code, http.StatusInternalServerError)
			}
		})

		t.Run("誤ったパスワードでログインするとエラーを返す", func(t *testing.T) {
			// 誤ったパスワード
			wrongPasswordReqBody := `{"email":"` + validEmail + `","password":"wrongpassword"}`

			// テスト実行
			_, c, rec := setupEchoContextWithBody(http.MethodPost, "/login", wrongPasswordReqBody)
			err := userController.LogIn(c)

			if err != nil {
				t.Errorf("LogIn() unexpected error: %v", err)
			}

			if rec.Code != http.StatusInternalServerError {
				t.Errorf("LogIn() with wrong password status code = %d, want %d",
					rec.Code, http.StatusInternalServerError)
			}
		})
	})
}
