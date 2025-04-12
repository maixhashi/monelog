package main

import (
	"fmt"
	"monelog/db"
	"monelog/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(
		&model.User{},
		&model.Task{},
		&model.CardStatement{},
	)
}