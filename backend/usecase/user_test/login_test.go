package user_test

import (
	"monelog/model"
	"testing"
)

func TestUserUsecase_Login(t *testing.T) {
	setupUserUsecaseTest()

	// テスト用ユーザーを作成
	testUserEmail, testUserPwd := createTestUser(t)

	t.Run("正常系", func(t *testing.T) {
		t.Run("有効な認証情報でログインに成功する", func(t *testing.T) {
			loginReq := model.UserLoginRequest{
				Email:    testUserEmail,
				Password: testUserPwd,
			}
		
			t.Logf("ログイン試行: Email=%s", loginReq.Email)
		
			// テスト実行
			tokenString, err := userUsecase.Login(loginReq)
		
			// 検証
			if err != nil {
				t.Errorf("Login() error = %v", err)
			}
		
			if validateJWTToken(t, tokenString) {
				t.Log("JWTトークンが正常に生成されました")
			}
		})
	})

	t.Run("異常系", func(t *testing.T) {
		t.Run("バリデーションエラーが発生する場合はログインに失敗する", func(t *testing.T) {
			// 無効なメールアドレス
			invalidUser := model.UserLoginRequest{
				Email:    "invalid-email",
				Password: testUserPwd,
			}
		
			t.Logf("無効なユーザーでログインを試行: Email=%s", invalidUser.Email)
		
			_, err := userUsecase.Login(invalidUser)
		
			// バリデーションエラーが発生するはず
			if err == nil {
				t.Error("無効なメールアドレスでエラーが返されませんでした")
			} else {
				t.Logf("期待通りバリデーションエラーが返されました: %v", err)
			}
		})
	
		t.Run("存在しないユーザーでログインに失敗する", func(t *testing.T) {
			nonExistUser := model.UserLoginRequest{
				Email:    wrongUserEmail,
				Password: testUserPwd,
			}
		
			t.Logf("存在しないユーザーでログインを試行: Email=%s", nonExistUser.Email)
		
			_, err := userUsecase.Login(nonExistUser)
		
			if err == nil {
				t.Error("存在しないユーザーでエラーが返されませんでした")
			} else {
				t.Logf("期待通りエラーが返されました: %v", err)
			}
		})
	
		t.Run("パスワードが間違っている場合はログインに失敗する", func(t *testing.T) {
			wrongPwdUser := model.UserLoginRequest{
				Email:    testUserEmail,
				Password: wrongUserPwd,
			}
		
			t.Logf("間違ったパスワードでログインを試行: Email=%s", wrongPwdUser.Email)
		
			_, err := userUsecase.Login(wrongPwdUser)
		
			if err == nil {
				t.Error("間違ったパスワードでエラーが返されませんでした")
			} else {
				t.Logf("期待通りエラーが返されました: %v", err)
			}
		})
	})
}
