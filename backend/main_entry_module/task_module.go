package main_entry_module

import (
	"gorm.io/gorm"
	
	"monelog/controller"
	"monelog/repository"
	"monelog/usecase"
	"monelog/validator"
)

// initTaskModule はタスク関連のモジュールを初期化します
// @Summary タスクモジュールの初期化
// @Description タスク関連のリポジトリ、ユースケース、コントローラーを初期化します
func (m *MainEntryPackage) initTaskModule(db *gorm.DB) {
	taskValidator := validator.NewTaskValidator()
	taskRepository := repository.NewTaskRepository(db)
	taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)
	m.TaskController = controller.NewTaskController(taskUsecase)
}
