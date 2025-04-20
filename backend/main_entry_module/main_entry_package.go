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
	DevCardStatementController controller.IDevCardStatementController
	CSVHistoryController      controller.ICSVHistoryController
	
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
	entry.initDevCardStatementModule(db)
	entry.initCSVHistoryModule(db)

	return entry
}

// initCardStatementModule はカード明細関連のモジュールを初期化する
func (e *MainEntryPackage) initCardStatementModule(db *gorm.DB) {
	cardStatementRepo := repository.NewCardStatementRepository(db)
	cardStatementValidator := validator.NewCardStatementValidator()
	cardStatementUsecase := usecase.NewCardStatementUsecase(cardStatementRepo, cardStatementValidator)
	e.CardStatementController = controller.NewCardStatementController(cardStatementUsecase)
}

// initDevCardStatementModule は開発環境限定のカード明細関連のモジュールを初期化する
func (e *MainEntryPackage) initDevCardStatementModule(db *gorm.DB) {
	devCardStatementRepo := repository.NewDevCardStatementRepository(db)
	devCardStatementValidator := validator.NewDevCardStatementValidator()
	devCardStatementUsecase := usecase.NewDevCardStatementUsecase(devCardStatementRepo, devCardStatementValidator)
	e.DevCardStatementController = controller.NewDevCardStatementController(devCardStatementUsecase)
}

// initCSVHistoryModule はCSV履歴関連のモジュールを初期化する
func (e *MainEntryPackage) initCSVHistoryModule(db *gorm.DB) {
	csvHistoryRepo := repository.NewCSVHistoryRepository(db)
	csvHistoryValidator := validator.NewCSVHistoryValidator()
	csvHistoryUsecase := usecase.NewCSVHistoryUsecase(csvHistoryRepo, csvHistoryValidator)
	e.CSVHistoryController = controller.NewCSVHistoryController(csvHistoryUsecase)
}