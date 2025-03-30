package task_test

import (
    "monelog/model"
    "testing"
)

func TestTaskRepository_GetAllTasks(t *testing.T) {
    setupTaskTest()
    
    tasks := []model.Task{
        {Title: "Task 1", UserId: taskTestUser.ID},
        {Title: "Task 2", UserId: taskTestUser.ID},
        {Title: "Task 3", UserId: taskOtherUser.ID},
    }
    
    for _, task := range tasks {
        taskDB.Create(&task)
    }
    
    t.Run("正常系", func(t *testing.T) {
        t.Run("正しいユーザーIDのタスクのみを取得する", func(t *testing.T) {
            result, err := taskRepo.GetAllTasks(taskTestUser.ID)
            
            if err != nil {
                t.Errorf("GetAllTasks() error = %v", err)
            }
            
            if len(result) != 2 {
                t.Errorf("GetAllTasks() got %d tasks, want 2", len(result))
            }
            
            titles := make(map[string]bool)
            for _, task := range result {
                titles[task.Title] = true
            }
            
            if !titles["Task 1"] || !titles["Task 2"] {
                t.Errorf("期待したタスクが結果に含まれていません: %v", result)
            }
        })
    })
}
