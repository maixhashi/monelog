package main_entry_module

import (
	"fmt"
	"os"
	"log"
)

// StartServer はサーバーを起動します
// ポート番号を引数で受け取るか、デフォルトの8080を使用します
func (m *MainEntryPackage) StartServer(port ...string) error {
	serverPort := "8080" // デフォルトポート
	
	// 環境変数からポートを取得（オプション）
	if envPort := os.Getenv("PORT"); envPort != "" {
		serverPort = envPort
	}
	
	// 引数でポートが指定されていれば、それを使用
	if len(port) > 0 && port[0] != "" {
		serverPort = port[0]
	}
	
	// ルーターを初期化
	e := m.InitRouter()
	
	// Swaggerが有効な場合、ログに表示
	if m.SwaggerEnabled {
		log.Printf("Swagger UI available at http://localhost:%s/swagger/index.html", serverPort)
	}
	
	// サーバー起動
	return e.Start(fmt.Sprintf(":%s", serverPort))
}
