package user_test

import (
	"monelog/model"
	"monelog/repository"
	"monelog/testutils"
	"monelog/usecase"
	"monelog/validator"
	"os"
	"testing"

	"gorm.io/gorm"
)

// テスト用の共通変数
var (
	userDb          *gorm.DB
	userRepo        repository.IUserRepository
	userValidator   validator.IUserValidator
	userUsecase     usecase.IUserUsecase
	wrongUserEmail  = "wrong@example.com"
	wrongUserPwd    = "wrongpassword"
)

// テスト前の共通セットアップ
func setupUserUsecaseTest() {
	// テストごとにデータベースをクリーンアップ
	if userDb != nil {
		testutils.CleanupTestDB(userDb)
	} else {
		// 初回のみデータベース接続を作成
		userDb = testutils.SetupTestDB()
		userRepo = repository.NewUserRepository(userDb)
		userValidator = validator.NewUserValidator()
		userUsecase = usecase.NewUserUsecase(userRepo, userValidator)
		
		// JWT用のSECRET環境変数を設定
		os.Setenv("SECRET", "test-secret-key")
	}

	// 既存のテストユーザーを明示的に削除（念のため）
	userDb.Exec("DELETE FROM users WHERE email LIKE 'test%@example.com'")
}

// テスト用のユーザーを作成
func createTestUser(t *testing.T) (string, string) {
	// 一意のメールアドレスを生成
	testUserEmail := generateUniqueEmail()
	testUserPwd := "password123"

	// テスト用ユーザーを登録
	testUser := model.UserSignupRequest{
		Email:    testUserEmail,
		Password: testUserPwd,
	}
	_, err := userUsecase.SignUp(testUser)
	if err != nil {
		t.Fatalf("テストユーザーの登録に失敗しました: %v", err)
	}
	
	return testUserEmail, testUserPwd
}
