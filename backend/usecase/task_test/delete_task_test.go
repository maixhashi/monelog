package task_test

import (
	"testing"
)

func TestTaskUsecase_DeleteTask(t *testing.T) {
	setupTaskUsecaseTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("自分のタスクを削除できる", func(t *testing.T) {
			// テスト用タスクの作成
			task := createTestTask(t, "Task to Delete", testUser.ID)
			t.Logf("削除対象タスク作成: ID=%d, Title=%s", task.ID, task.Title)
		
			// テスト実行
			err := taskUsecase.DeleteTask(testUser.ID, task.ID)
		
			// 検証
			if err != nil {
				t.Errorf("DeleteTask() error = %v", err)
			} else {
				t.Logf("タスク削除成功: ID=%d", task.ID)
			}
		
			// データベースから直接確認
			assertTaskNotExists(t, task.ID)
		})
	})

	t.Run("異常系", func(t *testing.T) {
		t.Run("存在しないタスクIDでの削除はエラーになる", func(t *testing.T) {
			t.Logf("存在しないID %d でタスク削除を試行", nonExistentTaskID)
		
			err := taskUsecase.DeleteTask(testUser.ID, nonExistentTaskID)
			if err == nil {
				t.Error("存在しないIDでの削除でエラーが返されませんでした")
			} else {
				t.Logf("期待通りエラーが返されました: %v", err)
			}
		})
	
		t.Run("他のユーザーのタスクは削除できない", func(t *testing.T) {
			// 他のユーザーのタスクを作成
			otherUserTask := createTestTask(t, "Other User's Task", otherUser.ID)
			t.Logf("他ユーザーのタスク作成: ID=%d, Title=%s, UserId=%d", otherUserTask.ID, otherUserTask.Title, otherUserTask.UserId)
		
			// 他ユーザーのタスクを削除しようとする
			err := taskUsecase.DeleteTask(testUser.ID, otherUserTask.ID)
			if err == nil {
				t.Error("他のユーザーのタスク削除でエラーが返されませんでした")
			} else {
				t.Logf("期待通りエラーが返されました: %v", err)
			}
		
			// データベースに残っていることを確認
			assertTaskExists(t, otherUserTask.ID, otherUserTask.Title, otherUser.ID)
		})
	})
}
