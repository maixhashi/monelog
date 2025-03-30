package task_test

import (
	"fmt"
	"monelog/model"
	"net/http"
	"testing"
)

func TestTaskController_CreateTask(t *testing.T) {
	setupTaskControllerTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("新しいタスクを作成できる", func(t *testing.T) {
			// テストリクエストの準備
			validTitle := generateValidTaskTitle()
			reqBody := fmt.Sprintf(`{"title":"%s"}`, validTitle)
			
			// テスト実行
			_, c, rec := setupEchoWithJWTAndBody(
				taskTestUser.ID, 
				http.MethodPost, 
				"/tasks", 
				reqBody,
			)
			err := taskController.CreateTask(c)
			
			// 検証
			if err != nil {
				t.Errorf("CreateTask() error = %v", err)
			}
			
			if rec.Code != http.StatusCreated {
				t.Errorf("CreateTask() status code = %d, want %d", rec.Code, http.StatusCreated)
			}
			
			// レスポンスボディをパース
			response := parseTaskResponse(t, rec.Body.Bytes())
			
			if response.Title != validTitle {
				t.Errorf("CreateTask() = %v, want title=%s", response, validTitle)
			}
			
			// データベースから直接確認
			var dbTask model.Task
			taskDB.First(&dbTask, response.ID)
			if dbTask.Title != validTitle || dbTask.UserId != taskTestUser.ID {
				t.Errorf("CreateTask() did not save task correctly, got title=%s, userId=%d", dbTask.Title, dbTask.UserId)
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("バリデーションエラーが発生するタスクは作成できない", func(t *testing.T) {
			// 無効なタイトル（最大長超過）
			invalidTitle := generateInvalidTaskTitle()
			reqBody := fmt.Sprintf(`{"title":"%s"}`, invalidTitle)
			
			// テスト実行
			_, c, rec := setupEchoWithJWTAndBody(
				taskTestUser.ID, 
				http.MethodPost, 
				"/tasks", 
				reqBody,
			)
			err := taskController.CreateTask(c)
			
			// コントローラーはエラーを返さずにJSONレスポンスとしてエラーメッセージを返す
			if err != nil {
				t.Errorf("CreateTask() returned unexpected error: %v", err)
			}
			
			// ステータスコードの確認
			if rec.Code != http.StatusInternalServerError {
				t.Errorf("CreateTask() with invalid title status code = %d, want %d", rec.Code, http.StatusInternalServerError)
			}
		})
		
		t.Run("JSONデコードエラーでバッドリクエストを返す", func(t *testing.T) {
			// 無効なJSON
			invalidJSON := `{"title": Invalid JSON`
			
			// テスト実行
			_, c, rec := setupEchoWithJWTAndBody(
				taskTestUser.ID, 
				http.MethodPost, 
				"/tasks", 
				invalidJSON,
			)
			err := taskController.CreateTask(c)
			
			// この場合はコントローラーがJSONレスポンスを返すので、
			// エラーオブジェクトではなくレスポンスのステータスコードを確認
			if err != nil {
				t.Errorf("CreateTask() unexpected error: %v", err)
			}
			
			if rec.Code != http.StatusBadRequest {
				t.Errorf("CreateTask() with invalid JSON status code = %d, want %d", 
					rec.Code, http.StatusBadRequest)
			}
		})
	})
}