package task_test

import (
    "monelog/model"
    "monelog/repository"
    "monelog/testutils"
    "gorm.io/gorm"
)

var (
    taskDB *gorm.DB
    taskRepo repository.ITaskRepository
    taskTestUser model.User
    taskOtherUser model.User
    nonExistentTaskID uint = 9999
)

func setupTaskTest() {
    taskDB = testutils.SetupTestDB()
    taskRepo = repository.NewTaskRepository(taskDB)
    
    taskTestUser = testutils.CreateTestUser(taskDB)
    taskOtherUser = testutils.CreateOtherUser(taskDB)
}
