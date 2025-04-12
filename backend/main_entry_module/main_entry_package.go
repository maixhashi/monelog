package main_entry_module

import (
	"monelog/controller"
	"monelog/repository"
	"monelog/usecase"
	"monelog/validator"
	"gorm.io/gorm"
)

// MainEntryPackage はアプリケーションの主要コンポーネントを保持する構造体
type MainEntryPackage struct {
	UserController            controller.IUserController
	TaskController            controller.ITaskController 
	CardStatementController   controller.ICardStatementController
	
	// Swaggerハンドラーを追加（オプション）
	SwaggerEnabled            bool
}

// NewMainEntryPackage は新しいMainEntryPackageインスタンスを作成する
func NewMainEntryPackage(db *gorm.DB) *MainEntryPackage {
	entry := &MainEntryPackage{
		SwaggerEnabled: true, // デフォルトで有効
	}
	
	// 各モジュールの初期化
	entry.initUserModule(db)
	entry.initTaskModule(db)
	entry.initCardStatementModule(db)

	return entry
}

// initCardStatementModule はカード明細関連のモジュールを初期化する
func (e *MainEntryPackage) initCardStatementModule(db *gorm.DB) {
	cardStatementRepo := repository.NewCardStatementRepository(db)
	cardStatementValidator := validator.NewCardStatementValidator()
	cardStatementUsecase := usecase.NewCardStatementUsecase(cardStatementRepo, cardStatementValidator)
	e.CardStatementController = controller.NewCardStatementController(cardStatementUsecase)
}