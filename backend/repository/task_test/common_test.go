package task_test

import (
    "monelog/model"
    "testing"
)

func createTestTask(title string, userId uint) *model.Task {
    task := &model.Task{
        Title:  title,
        UserId: userId,
    }
    taskDB.Create(task)
    return task
}

func validateTask(t *testing.T, task *model.Task) {
    if task.ID == 0 {
        t.Error("Task ID should not be zero")
    }
    if task.CreatedAt.IsZero() {
        t.Error("CreatedAt should not be zero")
    }
    if task.UpdatedAt.IsZero() {
        t.Error("UpdatedAt should not be zero")
    }
}
