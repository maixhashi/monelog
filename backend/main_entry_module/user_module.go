package main_entry_module

import (
	"gorm.io/gorm"
	
	"monelog/controller"
	"monelog/repository"
	"monelog/usecase"
	"monelog/validator"
)

func (m *MainEntryPackage) initUserModule(db *gorm.DB) {
	userValidator := validator.NewUserValidator()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	m.UserController = controller.NewUserController(userUsecase)
}
