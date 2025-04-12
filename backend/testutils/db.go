package testutils

import (
	"fmt"
	"monelog/model"
	"log"
	"math/rand"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ランダムなDB名を生成し、テスト間で競合しないようにする
func generateRandomDBName() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("file:memdb%d?mode=memory&cache=shared", rand.Int())
}

// テスト用のデータベース接続を設定する
func SetupTestDB() *gorm.DB {
	// 毎回ユニークなインメモリデータベースを使用
	dbName := generateRandomDBName()
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		log.Fatalf("テストデータベースの接続に失敗しました: %v", err)
	}

	// テーブルのマイグレーション
	db.AutoMigrate(
		&model.User{}, 
		&model.Task{},
		&model.CardStatement{},
	)

	return db
}

// データベースをクリーンアップする
func CleanupTestDB(db *gorm.DB) {
	// テーブルの全レコードを削除
	db.Exec("DELETE FROM tasks")
	db.Exec("DELETE FROM users")
}

// テスト用ユーザーを作成する
func CreateTestUser(db *gorm.DB) model.User {
	// ユニークなメールアドレスを生成
	rand.Seed(time.Now().UnixNano())
	email := fmt.Sprintf("test%d@example.com", rand.Int())
	
	user := model.User{
		Email:    email,
		Password: "password",
	}
	db.Create(&user)
	return user
}

// 別のテスト用ユーザーを作成する
func CreateOtherUser(db *gorm.DB) model.User {
	// ユニークなメールアドレスを生成
	rand.Seed(time.Now().UnixNano())
	email := fmt.Sprintf("other%d@example.com", rand.Int())
	
	user := model.User{
		Email:    email,
		Password: "password",
	}
	db.Create(&user)
	return user
}
