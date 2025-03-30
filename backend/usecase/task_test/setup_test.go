package task_test

import (
	"monelog/model"
	"monelog/repository"
	"monelog/testutils"
	"monelog/usecase"
	"monelog/validator"
	"testing"

	"gorm.io/gorm"
)

// テスト用の共通変数
var (
	taskDb          *gorm.DB
	taskRepo        repository.ITaskRepository
	taskValidator   validator.ITaskValidator
	taskUsecase     usecase.ITaskUsecase
	testUser        model.User
	otherUser       model.User
)

const nonExistentTaskID uint = 9999

// テスト前の共通セットアップ
func setupTaskUsecaseTest() {
	// テストごとにデータベースをクリーンアップ
	if taskDb != nil {
		testutils.CleanupTestDB(taskDb)
	} else {
		// 初回のみデータベース接続を作成
		taskDb = testutils.SetupTestDB()
		taskRepo = repository.NewTaskRepository(taskDb)
		taskValidator = validator.NewTaskValidator()
		taskUsecase = usecase.NewTaskUsecase(taskRepo, taskValidator)
	}
	
	// テストユーザーを作成
	testUser = testutils.CreateTestUser(taskDb)
	
	// 別のテストユーザーを作成
	otherUser = testutils.CreateOtherUser(taskDb)
}

// テスト用のタスクを作成
func createTestTask(t *testing.T, title string, userId uint) model.Task {
	task := model.Task{
		Title:  title,
		UserId: userId,
	}
	
	result := taskDb.Create(&task)
	if result.Error != nil {
		t.Fatalf("テストタスクの作成に失敗しました: %v", result.Error)
	}
	
	return task
}
