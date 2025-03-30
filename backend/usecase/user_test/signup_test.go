package user_test

import (
	"monelog/model"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestUserUsecase_SignUp(t *testing.T) {
	setupUserUsecaseTest()

	t.Run("正常系", func(t *testing.T) {
		t.Run("新規ユーザーを登録できる", func(t *testing.T) {
			// 一意のメールアドレスを生成
			testUserEmail := generateUniqueEmail()
			testUserPwd := "password123"

			// テスト用ユーザー
			signupReq := model.UserSignupRequest{
				Email:    testUserEmail,
				Password: testUserPwd,
			}
		
			t.Logf("ユーザー登録: Email=%s", signupReq.Email)
		
			// テスト実行
			userRes, err := userUsecase.SignUp(signupReq)
		
			// 検証
			if err != nil {
				t.Errorf("SignUp() error = %v", err)
			}
		
			if userRes.ID == 0 || userRes.Email != signupReq.Email {
				t.Errorf("SignUp() = %v, want email=%s and ID > 0", userRes, signupReq.Email)
			} else {
				t.Logf("生成されたユーザーID: %d", userRes.ID)
			}
		
			// データベースから直接確認
			var savedUser model.User
			userDb.First(&savedUser, userRes.ID)
		
			if savedUser.Email != signupReq.Email {
				t.Errorf("SignUp() saved email = %v, want %v", savedUser.Email, signupReq.Email)
			}
		
			// パスワードがハッシュ化されていることを確認
			err = bcrypt.CompareHashAndPassword([]byte(savedUser.Password), []byte(signupReq.Password))
			if err != nil {
				t.Errorf("パスワードが正しくハッシュ化されていません: %v", err)
			} else {
				t.Log("パスワードが正しくハッシュ化されています")
			}
		})
	})

	t.Run("異常系", func(t *testing.T) {
		t.Run("バリデーションエラーが発生する場合はユーザー登録に失敗する", func(t *testing.T) {
			// 無効なメールアドレス
			invalidUser := model.UserSignupRequest{
				Email:    "invalid-email",
				Password: "password123",
			}
		
			t.Logf("無効なユーザー登録を試行: Email=%s", invalidUser.Email)
		
			_, err := userUsecase.SignUp(invalidUser)
		
			// バリデーションエラーが発生するはず
			if err == nil {
				t.Error("無効なメールアドレスでエラーが返されませんでした")
			} else {
				t.Logf("期待通りバリデーションエラーが返されました: %v", err)
			}
		})
	
		t.Run("すでに登録されているメールアドレスで登録に失敗する", func(t *testing.T) {
			// 一意のメールアドレスを生成
			duplicateEmail := generateUniqueEmail()
		
			// 最初に1人目のユーザーを登録
			firstUser := model.UserSignupRequest{
				Email:    duplicateEmail,
				Password: "password123",
			}
			_, err := userUsecase.SignUp(firstUser)
			if err != nil {
				t.Fatalf("最初のユーザー登録に失敗しました: %v", err)
			}
		
			// 同じメールアドレスで2人目のユーザーを登録
			duplicateUser := model.UserSignupRequest{
				Email:    duplicateEmail,
				Password: "different_password",
			}
		
			t.Logf("重複するメールアドレスでユーザー登録を試行: Email=%s", duplicateUser.Email)
		
			_, err = userUsecase.SignUp(duplicateUser)
		
			// エラーが発生するはず
			if err == nil {
				t.Error("重複するメールアドレスでエラーが返されませんでした")
			} else {
				t.Logf("期待通りエラーが返されました: %v", err)
			}
		})
	})
}
