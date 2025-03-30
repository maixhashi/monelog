package task_test

import (
	"monelog/model"
	"monelog/testutils"
	"testing"
	"time"
)

func TestTaskUsecase_UpdateTask(t *testing.T) {
	setupTaskUsecaseTest()
	
	// テストデータの作成
	task := createTestTask(t, testutils.GenerateValidTitle(), testUser.ID)
	t.Logf("元のタスク作成: ID=%d, Title=%s", task.ID, task.Title)
	
	// 少し待って更新時間に差をつける
	time.Sleep(10 * time.Millisecond)
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("タスクのタイトルを更新できる", func(t *testing.T) {
			// 更新するタスク - 別の有効なタイトルを使用
			updatedTitle := testutils.GenerateValidTitle() + "2" // 末尾に2を追加して別のタイトルに
			updateRequest := model.TaskRequest{
				Title:  updatedTitle,
				UserId: testUser.ID,
			}
			
			t.Logf("タスク更新リクエスト: ID=%d, 新Title=%s", task.ID, updateRequest.Title)
			
			// テスト実行
			response, err := taskUsecase.UpdateTask(updateRequest, testUser.ID, task.ID)
			
			// 検証
			if err != nil {
				t.Errorf("UpdateTask() error = %v", err)
			} else {
				t.Log("タスク更新成功")
			}
			
			validateTaskResponse(t, response, task.ID, updateRequest.Title)
			
			// データベースから直接確認
			assertTaskExists(t, task.ID, updateRequest.Title, testUser.ID)
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("バリデーションエラーが発生するタスクは更新できない", func(t *testing.T) {
			// 無効な更新 - 長すぎるタイトル
			invalidTitle := testutils.GenerateInvalidTitle()
			invalidRequest := model.TaskRequest{
				Title:  invalidTitle,
				UserId: testUser.ID,
			}
			
			t.Logf("無効なタイトルでの更新を試行: %s (長さ: %d)", 
				invalidRequest.Title, len(invalidRequest.Title))
			
			_, err := taskUsecase.UpdateTask(invalidRequest, testUser.ID, task.ID)
			
			// バリデーションエラーが発生するはず
			if err == nil {
				t.Error("無効なタイトルでエラーが返されませんでした")
			} else {
				t.Logf("期待通りバリデーションエラーが返されました: %v", err)
			}
			
			// データベースに反映されていないことを確認
			var dbTask model.Task
			taskDb.First(&dbTask, task.ID)
			if dbTask.Title == invalidRequest.Title {
				t.Error("バリデーションエラーの更新がデータベースに反映されています")
			} else {
				t.Logf("データベース確認: Title=%s (変更されていない)", dbTask.Title)
			}
		})
		
		t.Run("存在しないタスクIDでの更新はエラーになる", func(t *testing.T) {
			updateAttempt := model.TaskRequest{
				Title:  "Valid Title",
				UserId: testUser.ID,
			}
			t.Logf("存在しないID %d でタスク更新を試行", nonExistentTaskID)
		
			_, err := taskUsecase.UpdateTask(updateAttempt, testUser.ID, nonExistentTaskID)
			if err == nil {
				t.Error("存在しないIDでの更新でエラーが返されませんでした")
			} else {
				t.Logf("期待通りエラーが返されました: %v", err)
			}
		})
	
		t.Run("他のユーザーのタスクは更新できない", func(t *testing.T) {
			// 他のユーザーのタスクを作成
			otherUserTask := createTestTask(t, "Other User's Task", otherUser.ID)
			t.Logf("他ユーザーのタスク: ID=%d, Title=%s, UserId=%d", otherUserTask.ID, otherUserTask.Title, otherUserTask.UserId)
		
			// 他ユーザーのタスクを更新しようとする
			updateAttempt := model.TaskRequest{
				Title:  "Attempted Update",
				UserId: testUser.ID,
			}
			_, err := taskUsecase.UpdateTask(updateAttempt, testUser.ID, otherUserTask.ID)
		
			if err == nil {
				t.Error("他のユーザーのタスク更新でエラーが返されませんでした")
			} else {
				t.Logf("期待通りエラーが返されました: %v", err)
			}
		
			// データベースに反映されていないことを確認
			assertTaskExists(t, otherUserTask.ID, otherUserTask.Title, otherUser.ID)
		})
	})
}