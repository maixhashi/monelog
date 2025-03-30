package main

import (
	"monelog/db"
	"monelog/main_entry_module"
	"log"
	
	_ "monelog/docs" // Swaggerドキュメントのインポート（重要）
)

// @title Blog CMS API
// @version 1.0
// @description ブログCMSのバックエンドAPI
// @host localhost:8080
// @BasePath /
// @schemes http
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @openapi 3.0.0
func main() {
	// データベース接続
	db := db.NewDB()
	
	// アプリケーション初期化
	app := main_entry_module.NewMainEntryPackage(db)
	
	// サーバー起動（エラーがあれば終了）
	if err := app.StartServer(); err != nil {
		log.Fatalf("サーバー起動エラー: %v", err)
	}
}