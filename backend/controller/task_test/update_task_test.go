package task_test

import (
	"fmt"
	"monelog/model"
	"net/http"
	"testing"
)

func TestTaskController_UpdateTask(t *testing.T) {
	setupTaskControllerTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("タスクのタイトルを更新できる", func(t *testing.T) {
			// テスト用タスクの作成
			task := createTestTask("Original Title", taskTestUser.ID)
			
			// 更新リクエストの準備
			updatedTitle := "Updated Title"
			reqBody := fmt.Sprintf(`{"title":"%s"}`, updatedTitle)
			
			// テスト実行
			_, c, rec := setupEchoWithTaskId(
				taskTestUser.ID, 
				task.ID, 
				http.MethodPut, 
				fmt.Sprintf("/tasks/%d", task.ID), 
				reqBody,
			)
			err := taskController.UpdateTask(c)
			
			// 検証
			if err != nil {
				t.Errorf("UpdateTask() error = %v", err)
			}
			
			if rec.Code != http.StatusOK {
				t.Errorf("UpdateTask() status code = %d, want %d", rec.Code, http.StatusOK)
			}
			
			// レスポンスボディをパース
			response := parseTaskResponse(t, rec.Body.Bytes())
			
			if response.Title != updatedTitle {
				t.Errorf("UpdateTask() = %v, want title=%s", response, updatedTitle)
			}
			
			// データベースから直接確認
			var dbTask model.Task
			taskDB.First(&dbTask, task.ID)
			if dbTask.Title != updatedTitle {
				t.Errorf("UpdateTask() did not update task correctly, got title=%s, want=%s", dbTask.Title, updatedTitle)
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("存在しないタスクIDでの更新はエラーになる", func(t *testing.T) {
			// 更新リクエストの準備
			reqBody := `{"title":"Invalid Update"}`
			
			// テスト実行
			_, c, rec := setupEchoWithTaskId(
				taskTestUser.ID, 
				nonExistentTaskID, 
				http.MethodPut, 
				fmt.Sprintf("/tasks/%d", nonExistentTaskID), 
				reqBody,
			)
			err := taskController.UpdateTask(c)
			
			// コントローラーはエラーを返さずにJSONレスポンスとしてエラーメッセージを返す
			if err != nil {
				t.Errorf("UpdateTask() returned unexpected error: %v", err)
			}
			
			// ステータスコードの確認
			if rec.Code != http.StatusInternalServerError {
				t.Errorf("UpdateTask() with non-existent ID status code = %d, want %d", rec.Code, http.StatusInternalServerError)
			}
		})

		t.Run("他のユーザーのタスクは更新できない", func(t *testing.T) {
			// 他のユーザーのタスクを作成
			otherUserTask := createTestTask("Other User's Task", taskOtherUser.ID)
			
			// 更新リクエストの準備
			reqBody := `{"title":"Attempted Update"}`
			
			// テスト実行
			_, c, rec := setupEchoWithTaskId(
				taskTestUser.ID, 
				otherUserTask.ID, 
				http.MethodPut, 
				fmt.Sprintf("/tasks/%d", otherUserTask.ID), 
				reqBody,
			)
			err := taskController.UpdateTask(c)
			
			// コントローラーはエラーを返さずにJSONレスポンスとしてエラーメッセージを返す
			if err != nil {
				t.Errorf("UpdateTask() returned unexpected error: %v", err)
			}
			
			// ステータスコードの確認
			if rec.Code != http.StatusInternalServerError {
				t.Errorf("UpdateTask() with other user's task status code = %d, want %d", rec.Code, http.StatusInternalServerError)
			}
			
			// データベースから直接確認して、タスクが更新されていないことを確認
			var dbTask model.Task
			taskDB.First(&dbTask, otherUserTask.ID)
			if dbTask.Title != "Other User's Task" {
				t.Errorf("UpdateTask() should not update other user's task, got title=%s", dbTask.Title)
			}
		})
		
		t.Run("バリデーションエラーが発生するタスクは更新できない", func(t *testing.T) {
			// テスト用タスクの作成
			task := createTestTask("Valid Title", taskTestUser.ID)
			
			// 無効なタイトル（最大長超過）
			invalidTitle := generateInvalidTaskTitle()
			reqBody := fmt.Sprintf(`{"title":"%s"}`, invalidTitle)
			
			// テスト実行
			_, c, rec := setupEchoWithTaskId(
				taskTestUser.ID, 
				task.ID, 
				http.MethodPut, 
				fmt.Sprintf("/tasks/%d", task.ID), 
				reqBody,
			)
			err := taskController.UpdateTask(c)
			
			// コントローラーはエラーを返さずにJSONレスポンスとしてエラーメッセージを返す
			if err != nil {
				t.Errorf("UpdateTask() returned unexpected error: %v", err)
			}
			
			// ステータスコードの確認
			if rec.Code != http.StatusInternalServerError {
				t.Errorf("UpdateTask() with invalid title status code = %d, want %d", rec.Code, http.StatusInternalServerError)
			}
		})
		
		t.Run("JSONデコードエラーでバッドリクエストを返す", func(t *testing.T) {
			// テスト用タスクの作成
			task := createTestTask("Valid Title", taskTestUser.ID)
			
			// 無効なJSON
			invalidJSON := `{"title": Invalid JSON`
			
			// テスト実行
			_, c, rec := setupEchoWithTaskId(
				taskTestUser.ID, 
				task.ID, 
				http.MethodPut, 
				fmt.Sprintf("/tasks/%d", task.ID), 
				invalidJSON,
			)
			err := taskController.UpdateTask(c)
			
			// この場合はコントローラーがJSONレスポンスを返すので、
			// エラーオブジェクトではなくレスポンスのステータスコードを確認
			if err != nil {
				t.Errorf("UpdateTask() unexpected error: %v", err)
			}
			
			if rec.Code != http.StatusBadRequest {
				t.Errorf("UpdateTask() with invalid JSON status code = %d, want %d", 
					rec.Code, http.StatusBadRequest)
			}
		})
	})
}
