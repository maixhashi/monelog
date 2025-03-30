package task_test

import (
    "monelog/model"
    "testing"
    "time"
)

func TestTaskRepository_UpdateTask(t *testing.T) {
    setupTaskTest()
    
    task := model.Task{
        Title:  "Original Title",
        UserId: taskTestUser.ID,
    }
    taskDB.Create(&task)
    
    time.Sleep(10 * time.Millisecond)
    
    t.Run("正常系", func(t *testing.T) {
        t.Run("タスクのタイトルを更新できる", func(t *testing.T) {
            updatedTask := model.Task{
                Title: "Updated Title",
            }
            
            err := taskRepo.UpdateTask(&updatedTask, taskTestUser.ID, task.ID)
            
            if err != nil {
                t.Errorf("UpdateTask() error = %v", err)
            }
            
            if updatedTask.Title != "Updated Title" {
                t.Errorf("UpdateTask() returned task title = %v, want %v", updatedTask.Title, "Updated Title")
            }
            
            var taskDBTask model.Task
            taskDB.First(&taskDBTask, task.ID)
            
            if taskDBTask.Title != "Updated Title" {
                t.Errorf("UpdateTask() database task title = %v, want %v", taskDBTask.Title, "Updated Title")
            }
        })
    })

    t.Run("異常系", func(t *testing.T) {
        t.Run("存在しないタスクIDでの更新はエラーになる", func(t *testing.T) {
            invalidTask := model.Task{Title: "Invalid Update"}
            err := taskRepo.UpdateTask(&invalidTask, taskTestUser.ID, nonExistentTaskID)
            
            if err == nil {
                t.Error("UpdateTask() with non-existent ID should return error")
            }
        })

        t.Run("他のユーザーのタスクは更新できない", func(t *testing.T) {
            otherUserTask := model.Task{
                Title:  "Other User's Task",
                UserId: taskOtherUser.ID,
            }
            taskDB.Create(&otherUserTask)
            
            updateAttempt := model.Task{Title: "Attempted Update"}
            err := taskRepo.UpdateTask(&updateAttempt, taskTestUser.ID, otherUserTask.ID)
            
            if err == nil {
                t.Error("UpdateTask() should not allow updating other user's task")
            }
        })
    })
}
