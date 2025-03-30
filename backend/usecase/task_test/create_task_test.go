package task_test

import (
	"monelog/model"
	"monelog/testutils"
	"testing"
)

func TestTaskUsecase_CreateTask(t *testing.T) {
	setupTaskUsecaseTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("新しいタスクを作成できる", func(t *testing.T) {
			// テスト用タスク - 有効なタイトルを生成関数から取得
			validTitle := testutils.GenerateValidTitle()
			validRequest := model.TaskRequest{
				Title:  validTitle,
				UserId: testUser.ID,
			}
			
			t.Logf("タスク作成: Title=%s, UserId=%d", validRequest.Title, validRequest.UserId)
			
			// テスト実行
			response, err := taskUsecase.CreateTask(validRequest)
			
			// 検証
			if err != nil {
				t.Errorf("CreateTask() error = %v", err)
			}
			
			validateTaskResponse(t, response, response.ID, validRequest.Title)
			
			// データベースから直接確認
			assertTaskExists(t, response.ID, validRequest.Title, testUser.ID)
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("バリデーションエラーが発生するタスクは作成できない", func(t *testing.T) {
			// 無効なタスク - 長すぎるタイトルをヘルパー関数で生成
			invalidTitle := testutils.GenerateInvalidTitle()
			invalidRequest := model.TaskRequest{
				Title:  invalidTitle,
				UserId: testUser.ID,
			}
			
			t.Logf("無効なタスク作成を試行: Title=%s (長さ: %d)", 
				invalidRequest.Title, len(invalidRequest.Title))
			
			_, err := taskUsecase.CreateTask(invalidRequest)
			
			// バリデーションエラーが発生するはず
			if err == nil {
				t.Error("無効なタスクでエラーが返されませんでした")
			} else {
				t.Logf("期待通りバリデーションエラーが返されました: %v", err)
			}
			
			// データベースに保存されていないことを確認
			var count int64
			taskDb.Model(&model.Task{}).Where("title = ?", invalidRequest.Title).Count(&count)
			if count > 0 {
				t.Error("バリデーションエラーのタスクがデータベースに保存されています")
			} else {
				t.Log("バリデーションエラーのタスクは保存されていないことを確認")
			}
		})
	})
}
