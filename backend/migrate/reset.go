package main

import (
	"fmt"
	"monelog/db"
	"monelog/model"
	"gorm.io/gorm"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Reset")
	defer db.CloseDB(dbConn)

	resetDatabase(dbConn)
}

func resetDatabase(dbConn *gorm.DB) {
	// Drop tables
	dbConn.Migrator().DropTable(
		&model.User{},
		&model.Task{},
		&model.CardStatement{},
		&model.CSVHistory{},
	)

	// Recreate tables
	dbConn.AutoMigrate(
		&model.User{},
		&model.Task{},
		&model.CardStatement{},
		&model.CSVHistory{},
	)
}