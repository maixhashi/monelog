package task_test

import (
	"monelog/model"
	"testing"
)

func TestTaskUsecase_GetAllTasks(t *testing.T) {
	setupTaskUsecaseTest()
	
	// テストデータの作成
	tasks := []model.Task{
		{Title: "Task 1", UserId: testUser.ID},
		{Title: "Task 2", UserId: testUser.ID},
		{Title: "Task 3", UserId: otherUser.ID}, // 別ユーザーのタスク
	}
	
	for _, task := range tasks {
		taskDb.Create(&task)
	}
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("正しいユーザーIDのタスクのみを取得する", func(t *testing.T) {
			t.Logf("ユーザーID %d のタスクを取得します", testUser.ID)
			
			taskResponses, err := taskUsecase.GetAllTasks(testUser.ID)
			
			if err != nil {
				t.Errorf("GetAllTasks() error = %v", err)
			}
			
			if len(taskResponses) != 2 {
				t.Errorf("GetAllTasks() got %d tasks, want 2", len(taskResponses))
			}
			
			// タスクタイトルの確認
			titles := make(map[string]bool)
			for _, task := range taskResponses {
				titles[task.Title] = true
				t.Logf("取得したタスク: ID=%d, Title=%s", task.ID, task.Title)
			}
			
			if !titles["Task 1"] || !titles["Task 2"] {
				t.Errorf("期待したタスクが結果に含まれていません: %v", taskResponses)
			}
			
			// レスポンス形式の検証
			for _, task := range taskResponses {
				if task.ID == 0 || task.Title == "" || task.CreatedAt.IsZero() || task.UpdatedAt.IsZero() {
					t.Errorf("GetAllTasks() returned invalid task: %+v", task)
				}
			}
		})
	})
}

func TestTaskUsecase_GetTaskById(t *testing.T) {
	setupTaskUsecaseTest()
	
	// テストデータの作成
	task := createTestTask(t, "Test Task", testUser.ID)
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("存在するタスクを正しく取得する", func(t *testing.T) {
			t.Logf("タスクID %d を取得します", task.ID)
			
			response, err := taskUsecase.GetTaskById(testUser.ID, task.ID)
			
			if err != nil {
				t.Errorf("GetTaskById() error = %v", err)
			}
			
			validateTaskResponse(t, response, task.ID, task.Title)
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("存在しないIDを指定した場合はエラーを返す", func(t *testing.T) {
			t.Logf("存在しないID %d を指定してタスクを取得しようとします", nonExistentTaskID)
			
			_, err := taskUsecase.GetTaskById(testUser.ID, nonExistentTaskID)
			
			if err == nil {
				t.Error("存在しないIDを指定したときにエラーが返されませんでした")
			} else {
				t.Logf("期待通りエラーが返されました: %v", err)
			}
		})
		
		t.Run("他のユーザーのタスクは取得できない", func(t *testing.T) {
			// 他のユーザーのタスクを作成
			otherUserTask := createTestTask(t, "Other User's Task", otherUser.ID)
			t.Logf("他ユーザーのタスク(ID=%d)を別ユーザー(ID=%d)として取得しようとします", otherUserTask.ID, testUser.ID)
			
			_, err := taskUsecase.GetTaskById(testUser.ID, otherUserTask.ID)
			
			if err == nil {
				t.Error("他のユーザーのタスクを取得できてしまいました")
			} else {
				t.Logf("期待通りエラーが返されました: %v", err)
			}
		})
	})
}
