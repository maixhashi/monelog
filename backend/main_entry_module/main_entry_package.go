package main_entry_module

import (
	"monelog/controller"
	"gorm.io/gorm"
)

// MainEntryPackage はアプリケーションの主要コンポーネントを保持する構造体
type MainEntryPackage struct {
	UserController            controller.IUserController
	TaskController            controller.ITaskController 
	
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

	return entry
}