package task_test

import (
	"fmt"
	"net/http"
	"testing"
)

func TestTaskController_GetTaskById(t *testing.T) {
	setupTaskControllerTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("存在するタスクを正しく取得する", func(t *testing.T) {
			// テスト用タスクの作成
			task := createTestTask("Test Task", taskTestUser.ID)
			
			// テスト実行
			_, c, rec := setupEchoWithTaskId(
				taskTestUser.ID, 
				task.ID, 
				http.MethodGet, 
				fmt.Sprintf("/tasks/%d", task.ID), 
				"",
			)
			err := taskController.GetTaskById(c)
			
			// 検証
			if err != nil {
				t.Errorf("GetTaskById() error = %v", err)
			}
			
			if rec.Code != http.StatusOK {
				t.Errorf("GetTaskById() status code = %d, want %d", rec.Code, http.StatusOK)
			}
			
			// レスポンスボディをパース
			response := parseTaskResponse(t, rec.Body.Bytes())
			
			if response.ID != task.ID || response.Title != task.Title {
				t.Errorf("GetTaskById() = %v, want id=%d, title=%s", response, task.ID, task.Title)
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("他のユーザーのタスクは取得できない", func(t *testing.T) {
			// 他のユーザーのタスクを作成
			otherUserTask := createTestTask("Other User's Task", taskOtherUser.ID)
			
			// テスト実行 - testUserとして他のユーザーのタスクにアクセス
			_, c, rec := setupEchoWithTaskId(
				taskTestUser.ID, 
				otherUserTask.ID, 
				http.MethodGet, 
				fmt.Sprintf("/tasks/%d", otherUserTask.ID), 
				"",
			)
			err := taskController.GetTaskById(c)
			
			// コントローラーはエラーを返さずにJSONレスポンスとしてエラーメッセージを返す
			if err != nil {
				t.Errorf("GetTaskById() returned unexpected error: %v", err)
			}
			
			// ステータスコードの確認
			if rec.Code != http.StatusInternalServerError {
				t.Errorf("GetTaskById() with other user's task status code = %d, want %d", rec.Code, http.StatusInternalServerError)
			}
		})
		
		t.Run("存在しないタスクIDでエラーを返す", func(t *testing.T) {
			// テスト実行 - 存在しないタスクID
			_, c, rec := setupEchoWithTaskId(
				taskTestUser.ID, 
				nonExistentTaskID, 
				http.MethodGet, 
				fmt.Sprintf("/tasks/%d", nonExistentTaskID), 
				"",
			)
			err := taskController.GetTaskById(c)
			
			// コントローラーはエラーを返さずにJSONレスポンスとしてエラーメッセージを返す
			if err != nil {
				t.Errorf("GetTaskById() returned unexpected error: %v", err)
			}
			
			// ステータスコードの確認
			if rec.Code != http.StatusInternalServerError {
				t.Errorf("GetTaskById() with non-existent ID status code = %d, want %d", rec.Code, http.StatusInternalServerError)
			}
		})
	})
}