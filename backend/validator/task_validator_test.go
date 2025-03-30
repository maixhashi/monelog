package validator

import (
	"monelog/model"
	"monelog/testutils"
	"testing"
)

func TestTaskValidate(t *testing.T) {
	// テスト用DBの設定
	db := testutils.SetupTestDB()
	defer testutils.CleanupTestDB(db)
	
	// テストユーザーの作成
	user := testutils.CreateTestUser(db)
	
	validator := NewTaskValidator()

	testCases := []struct {
		name     string
		request  model.TaskRequest
		hasError bool
	}{
		{
			name: "Valid task with valid title",
			request: model.TaskRequest{
				Title:  testutils.GenerateValidTitle(),
				UserId: user.ID,
			},
			hasError: false,
		},
		{
			name: "Empty title",
			request: model.TaskRequest{
				Title:  "",
				UserId: user.ID,
			},
			hasError: true,
		},
		{
			name: "Title too long",
			request: model.TaskRequest{
				Title:  testutils.GenerateInvalidTitle(),
				UserId: user.ID,
			},
			hasError: true,
		},
		{
			name: "Valid title with exact max length",
			request: model.TaskRequest{
				Title:  generateExactMaxLengthTitle(),
				UserId: user.ID,
			},
			hasError: false,
		},
		{
			name: "Zero user ID",
			request: model.TaskRequest{
				Title:  "Valid Title",
				UserId: 0,
			},
			hasError: false, // UserIDはバリデーションしていないので、エラーにならないはず
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validator.ValidateTaskRequest(tc.request)
			if (err != nil) != tc.hasError {
				t.Errorf("ValidateTaskRequest() error = %v, want error: %v", err, tc.hasError)
			}
		})
	}
}

// 追加の別テストケース - テスト用DBを使ったバリデーション
func TestTaskValidateWithDB(t *testing.T) {
	// テスト用DBの設定
	db := testutils.SetupTestDB()
	defer testutils.CleanupTestDB(db)
	
	// 異なるテストユーザーを作成
	user1 := testutils.CreateTestUser(db)
	user2 := testutils.CreateOtherUser(db)
	
	validator := NewTaskValidator()

	// ユーザー1のタスクを作成
	request1 := model.TaskRequest{
		Title:  "User 1 Task",
		UserId: user1.ID,
	}
	
	// ユーザー2のタスクを作成
	request2 := model.TaskRequest{
		Title:  "User 2 Task",
		UserId: user2.ID,
	}
	
	// バリデーションのテスト
	err1 := validator.ValidateTaskRequest(request1)
	if err1 != nil {
		t.Errorf("ValidateTaskRequest() for user1 should not return error, got: %v", err1)
	}
	
	err2 := validator.ValidateTaskRequest(request2)
	if err2 != nil {
		t.Errorf("ValidateTaskRequest() for user2 should not return error, got: %v", err2)
	}
}

// 最大長さちょうどのタイトルを生成する関数
func generateExactMaxLengthTitle() string {
	title := ""
	for i := 0; i < model.TaskTitleMaxLength; i++ {
		title += "x"
	}
	return title
}
