package task_test

import (
    "monelog/model"
    "testing"
)

func TestTaskRepository_DeleteTask(t *testing.T) {
    setupTaskTest()
    
    t.Run("正常系", func(t *testing.T) {
        t.Run("自分のタスクを削除できる", func(t *testing.T) {
            task := model.Task{
                Title:  "Task to Delete",
                UserId: taskTestUser.ID,
            }
            taskDB.Create(&task)
            
            err := taskRepo.DeleteTask(taskTestUser.ID, task.ID)
            
            if err != nil {
                t.Errorf("DeleteTask() error = %v", err)
            }
            
            var count int64
            taskDB.Model(&model.Task{}).Where("id = ?", task.ID).Count(&count)
            
            if count != 0 {
                t.Error("DeleteTask() did not delete the task from database")
            }
        })
    })
    
    t.Run("異常系", func(t *testing.T) {
        t.Run("存在しないタスクIDでの削除はエラーになる", func(t *testing.T) {
            err := taskRepo.DeleteTask(taskTestUser.ID, nonExistentTaskID)
            
            if err == nil {
                t.Error("DeleteTask() with non-existent ID should return error")
            }
        })
        
        t.Run("他のユーザーのタスクは削除できない", func(t *testing.T) {
            otherUserTask := model.Task{
                Title:  "Other User's Task",
                UserId: taskOtherUser.ID,
            }
            taskDB.Create(&otherUserTask)
            
            err := taskRepo.DeleteTask(taskTestUser.ID, otherUserTask.ID)
            
            if err == nil {
                t.Error("DeleteTask() should not allow deleting other user's task")
            }
            
            var count int64
            taskDB.Model(&model.Task{}).Where("id = ?", otherUserTask.ID).Count(&count)
            if count == 0 {
                t.Error("他ユーザーのタスクが削除されています")
            }
        })
    })
}
