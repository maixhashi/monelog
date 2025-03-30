package task_test

import (
	"monelog/model"
	"net/http"
	"testing"
)

func TestTaskController_GetAllTasks(t *testing.T) {
	setupTaskControllerTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("ユーザーのタスクを全て取得する", func(t *testing.T) {
			// テスト用タスクの作成
			tasks := []model.Task{
				{Title: "Task 1", UserId: taskTestUser.ID},
				{Title: "Task 2", UserId: taskTestUser.ID},
				{Title: "Task 3", UserId: taskOtherUser.ID}, // 別ユーザーのタスク
			}
			
			for _, task := range tasks {
				taskDB.Create(&task)
			}
			
			// テスト実行
			_, c, rec := setupEchoWithJWT(taskTestUser.ID)
			err := taskController.GetAllTasks(c)
			
			// 検証
			if err != nil {
				t.Errorf("GetAllTasks() error = %v", err)
			}
			
			if rec.Code != http.StatusOK {
				t.Errorf("GetAllTasks() status code = %d, want %d", rec.Code, http.StatusOK)
			}
			
			// レスポンスボディをパース
			response := parseTasksResponse(t, rec.Body.Bytes())
			
			if len(response) != 2 {
				t.Errorf("GetAllTasks() returned %d tasks, want 2", len(response))
			}
			
			// タスクタイトルの確認
			titles := make(map[string]bool)
			for _, task := range response {
				titles[task.Title] = true
			}
			
			if !titles["Task 1"] || !titles["Task 2"] {
				t.Errorf("期待したタスクが結果に含まれていません: %v", response)
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		// データベース接続エラーなどのケースをモックして追加可能
		// 現在の実装では直接テストできないため省略
	})
}
