package task_test

import (
    "monelog/model"
    "testing"
)

func TestTaskRepository_CreateTask(t *testing.T) {
    setupTaskTest()
    
    t.Run("正常系", func(t *testing.T) {
        t.Run("新しいタスクを作成できる", func(t *testing.T) {
            task := model.Task{
                Title:  "New Task",
                UserId: taskTestUser.ID,
            }
            
            err := taskRepo.CreateTask(&task)
            
            if err != nil {
                t.Errorf("CreateTask() error = %v", err)
            }
            
            validateTask(t, &task)
        })
    })
}
