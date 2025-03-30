package task_test

import (
	"fmt"
	"monelog/model"
	"net/http"
	"testing"
)

func TestTaskController_DeleteTask(t *testing.T) {
	setupTaskControllerTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("自分のタスクを削除できる", func(t *testing.T) {
			// テスト用タスクの作成
			task := createTestTask("Task to Delete", taskTestUser.ID)
			
			// テスト実行
			_, c, rec := setupEchoWithTaskId(
				taskTestUser.ID, 
				task.ID, 
				http.MethodDelete, 
				fmt.Sprintf("/tasks/%d", task.ID), 
				"",
			)
			err := taskController.DeleteTask(c)
			
			// 検証
			if err != nil {
				t.Errorf("DeleteTask() error = %v", err)
			}
			
			if rec.Code != http.StatusNoContent {
				t.Errorf("DeleteTask() status code = %d, want %d", rec.Code, http.StatusNoContent)
			}
			
			// データベースから直接確認
			var count int64
			taskDB.Model(&model.Task{}).Where("id = ?", task.ID).Count(&count)
			
			if count != 0 {
				t.Error("DeleteTask() did not delete the task from database")
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("存在しないタスクIDでの削除はエラーになる", func(t *testing.T) {
			// テスト実行
			_, c, rec := setupEchoWithTaskId(
				taskTestUser.ID, 
				nonExistentTaskID, 
				http.MethodDelete, 
				fmt.Sprintf("/tasks/%d", nonExistentTaskID), 
				"",
			)
			err := taskController.DeleteTask(c)
			
			// コントローラーはエラーを返さずにJSONレスポンスとしてエラーメッセージを返す
			if err != nil {
				t.Errorf("DeleteTask() returned unexpected error: %v", err)
			}
			
			// ステータスコードの確認
			if rec.Code != http.StatusInternalServerError {
				t.Errorf("DeleteTask() with non-existent ID status code = %d, want %d", rec.Code, http.StatusInternalServerError)
			}
		})
		
		t.Run("他のユーザーのタスクは削除できない", func(t *testing.T) {
			// 他のユーザーのタスクを作成
			otherUserTask := createTestTask("Other User's Task", taskOtherUser.ID)
			
			// テスト実行
			_, c, rec := setupEchoWithTaskId(
				taskTestUser.ID, 
				otherUserTask.ID, 
				http.MethodDelete, 
				fmt.Sprintf("/tasks/%d", otherUserTask.ID), 
				"",
			)
			err := taskController.DeleteTask(c)
			
			// コントローラーはエラーを返さずにJSONレスポンスとしてエラーメッセージを返す
			if err != nil {
				t.Errorf("DeleteTask() returned unexpected error: %v", err)
			}
			
			// ステータスコードの確認
			if rec.Code != http.StatusInternalServerError {
				t.Errorf("DeleteTask() with other user's task status code = %d, want %d", rec.Code, http.StatusInternalServerError)
			}
			
			// データベースから直接確認して、タスクが削除されていないことを確認
			var count int64
			taskDB.Model(&model.Task{}).Where("id = ?", otherUserTask.ID).Count(&count)
			if count == 0 {
				t.Error("DeleteTask() should not delete other user's task")
			}
		})
		
		t.Run("無効なタスクIDパラメータでエラーを返す", func(t *testing.T) {
			// 無効なIDパラメータ
			_, c, rec := setupEchoWithJWTAndBody(
				taskTestUser.ID, 
				http.MethodDelete, 
				"/tasks/invalid", 
				"",
			)
			c.SetParamNames("taskId")
			c.SetParamValues("invalid")
			
			err := taskController.DeleteTask(c)
			
			// コントローラーはエラーを返さずにJSONレスポンスとしてエラーメッセージを返す
			if err != nil {
				t.Errorf("DeleteTask() returned unexpected error: %v", err)
			}
			
			// ステータスコードの確認
			if rec.Code != http.StatusBadRequest {
				t.Errorf("DeleteTask() with invalid ID parameter status code = %d, want %d", rec.Code, http.StatusBadRequest)
			}
		})
	})
}