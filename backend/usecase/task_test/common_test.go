package task_test

import (
	"monelog/model"
	"testing"
)

// タスクの存在を確認するヘルパー関数
func assertTaskExists(t *testing.T, taskId uint, expectedTitle string, expectedUserId uint) bool {
	var task model.Task
	result := taskDb.First(&task, taskId)
	
	if result.Error != nil {
		t.Errorf("タスク(ID=%d)がデータベースに存在しません: %v", taskId, result.Error)
		return false
	}
	
	if task.Title != expectedTitle {
		t.Errorf("タスクのタイトルが一致しません: got=%s, want=%s", task.Title, expectedTitle)
		return false
	}
	
	if task.UserId != expectedUserId {
		t.Errorf("タスクのユーザーIDが一致しません: got=%d, want=%d", task.UserId, expectedUserId)
		return false
	}
	
	return true
}

// タスクが存在しないことを確認するヘルパー関数
func assertTaskNotExists(t *testing.T, taskId uint) bool {
	var count int64
	taskDb.Model(&model.Task{}).Where("id = ?", taskId).Count(&count)
	
	if count != 0 {
		t.Errorf("タスク(ID=%d)がデータベースに存在します", taskId)
		return false
	}
	
	return true
}

// タスクレスポンスの検証ヘルパー関数
func validateTaskResponse(t *testing.T, task model.TaskResponse, expectedId uint, expectedTitle string) bool {
	if task.ID != expectedId {
		t.Errorf("タスクIDが一致しません: got=%d, want=%d", task.ID, expectedId)
		return false
	}
	
	if task.Title != expectedTitle {
		t.Errorf("タスクタイトルが一致しません: got=%s, want=%s", task.Title, expectedTitle)
		return false
	}
	
	if task.CreatedAt.IsZero() || task.UpdatedAt.IsZero() {
		t.Errorf("タスクのタイムスタンプが正しく設定されていません: %+v", task)
		return false
	}
	
	return true
}
