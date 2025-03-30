package user_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestUserController_CsrfToken(t *testing.T) {
	setupUserControllerTest()

	t.Run("正常系", func(t *testing.T) {
		t.Run("CSRFトークンを取得できる", func(t *testing.T) {
			// テスト実行
			_, c, rec := setupEchoContextWithCSRF(http.MethodGet, "/csrf-token")
			err := userController.CsrfToken(c)

			// 検証
			if err != nil {
				t.Errorf("CsrfToken() error = %v", err)
			}

			if rec.Code != http.StatusOK {
				t.Errorf("CsrfToken() status code = %d, want %d", rec.Code, http.StatusOK)
			}

			// レスポンスをパース
			var response map[string]string
			err = json.Unmarshal(rec.Body.Bytes(), &response)
			if err != nil {
				t.Errorf("Failed to unmarshal response: %v", err)
			}

			// CSRFトークンが含まれていることを確認
			token, exists := response["csrf_token"]
			if !exists {
				t.Error("CsrfToken() response does not contain csrf_token field")
			}

			if token != "test-csrf-token" {
				t.Errorf("CsrfToken() returned token = %s, want %s", token, "test-csrf-token")
			}
		})
	})

	t.Run("異常系", func(t *testing.T) {
		t.Run("CSRFトークンが設定されていない場合はエラーを返す", func(t *testing.T) {
			// CSRFトークンなしでコンテキストを設定
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/csrf-token", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			
			// テスト実行 - CSRFトークンがないのでパニックが発生するはず
			// パニックをキャッチするためにdeferとrecoverを使用
			defer func() {
				if r := recover(); r == nil {
					t.Error("CsrfToken() did not panic when csrf token is missing")
				}
			}()
			
			// このコードはパニックを起こすはず
			userController.CsrfToken(c)
		})
	})
}
