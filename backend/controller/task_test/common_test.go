package task_test

import (
	"encoding/json"
	"fmt"
	"monelog/model"
	"testing"
	"time"
	"strings"
)

// テスト用タスクを作成するヘルパー関数
func createTestTask(title string, userId uint) *model.Task {
	task := &model.Task{
		Title:  title,
		UserId: userId,
	}
	taskDB.Create(task)
	return task
}

// レスポンスボディをパースするヘルパー関数
func parseTaskResponse(t *testing.T, responseBody []byte) model.TaskResponse {
	var response model.TaskResponse
	err := json.Unmarshal(responseBody, &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}
	return response
}

// 複数タスクのレスポンスボディをパースするヘルパー関数
func parseTasksResponse(t *testing.T, responseBody []byte) []model.TaskResponse {
	var response []model.TaskResponse
	err := json.Unmarshal(responseBody, &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}
	return response
}

// 有効なタスクタイトルを生成するヘルパー関数
func generateValidTaskTitle() string {
	return fmt.Sprintf("Test Task %d", time.Now().UnixNano() % 10000)
}

// 無効なタスクタイトル（最大長超過）を生成するヘルパー関数
func generateInvalidTaskTitle() string {
	// TaskTitleMaxLength = 50 を超える長さのタイトル
	return strings.Repeat("X", model.TaskTitleMaxLength + 1)
}