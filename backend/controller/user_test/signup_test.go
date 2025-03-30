package user_test

import (
	"fmt"
	"net/http"
	"testing"
)

func TestUserController_SignUp(t *testing.T) {
	setupUserControllerTest()

	t.Run("正常系", func(t *testing.T) {
		t.Run("有効なユーザー情報でアカウント作成に成功する", func(t *testing.T) {
			// テスト用の一意なメールアドレスを生成
			validEmail := generateTestEmail()
			validPassword := "password123"
			reqBody := fmt.Sprintf(`{"email":"%s","password":"%s"}`, validEmail, validPassword)

			t.Logf("テスト実行: email=%s, password=%s", validEmail, validPassword)

			// テスト実行
			_, c, rec := setupEchoContextWithBody(http.MethodPost, "/signup", reqBody)
			err := userController.SignUp(c)

			// レスポンスの詳細をログ出力
			t.Logf("レスポンス: status=%d, body=%s", rec.Code, rec.Body.String())

			// 検証
			if err != nil {
				t.Errorf("SignUp() error = %v", err)
			}

			if rec.Code != http.StatusCreated {
				t.Errorf("SignUp() status code = %d, want %d", rec.Code, http.StatusCreated)
			}

			// レスポンスボディをパース
			response := parseUserResponse(t, rec.Body.Bytes())

			if response.Email != validEmail {
				t.Errorf("SignUp() = %v, want email=%s", response, validEmail)
			}

			// IDが設定されていることを確認
			if response.ID == 0 {
				t.Error("SignUp() did not set ID in response")
			}
		})
	})

	t.Run("異常系", func(t *testing.T) {
		t.Run("無効なJSONリクエストでバッドリクエストを返す", func(t *testing.T) {
			// 無効なJSON
			invalidJSON := `{"email": "invalid@example.com", "password": Invalid JSON`

			// テスト実行
			_, c, rec := setupEchoContextWithBody(http.MethodPost, "/signup", invalidJSON)
			err := userController.SignUp(c)

			// エラーがあるはずだが、コントローラーはJSONレスポンスを返す
			if err != nil {
				t.Errorf("SignUp() unexpected error: %v", err)
			}

			if rec.Code != http.StatusBadRequest {
				t.Errorf("SignUp() with invalid JSON status code = %d, want %d",
					rec.Code, http.StatusBadRequest)
			}
		})

		t.Run("バリデーションエラーが発生する場合はエラーを返す", func(t *testing.T) {
			// 短すぎるパスワード
			invalidReqBody := `{"email":"test@example.com","password":"12"}`

			// テスト実行
			_, c, rec := setupEchoContextWithBody(http.MethodPost, "/signup", invalidReqBody)
			err := userController.SignUp(c)

			// コントローラーはエラーを返さずにJSONレスポンスとしてエラーメッセージを返す
			if err != nil {
				t.Errorf("SignUp() returned unexpected error: %v", err)
			}

			// ステータスコードの確認
			if rec.Code != http.StatusInternalServerError {
				t.Errorf("SignUp() with validation error status code = %d, want %d", 
                   rec.Code, http.StatusInternalServerError)
			}
		})

		t.Run("既に存在するユーザーは登録できない", func(t *testing.T) {
			// 最初のユーザー登録
			email := "duplicate@example.com"
			password := "password123"
			reqBody := fmt.Sprintf(`{"email":"%s","password":"%s"}`, email, password)

			_, c, _ := setupEchoContextWithBody(http.MethodPost, "/signup", reqBody)
			userController.SignUp(c)

			// 同じメールアドレスで再登録を試みる
			_, c2, rec2 := setupEchoContextWithBody(http.MethodPost, "/signup", reqBody)
			err := userController.SignUp(c2)

			if err != nil {
				t.Errorf("SignUp() returned unexpected error: %v", err)
			}

			// 重複登録はエラーになるはず
			if rec2.Code != http.StatusInternalServerError {
				t.Errorf("SignUp() with duplicate email status code = %d, want %d",
					rec2.Code, http.StatusInternalServerError)
			}
		})
	})
}
