package card_statement_test

import (
	"monelog/model"
	"monelog/repository"
	"monelog/testutils"
	"gorm.io/gorm"
)

var (
	csDB *gorm.DB
	csRepo repository.ICardStatementRepository
	csTestUser model.User
	csOtherUser model.User
	nonExistentCardStatementID uint = 9999
)

func setupCardStatementTest() {
	csDB = testutils.SetupTestDB()
	
	// テストデータベースにマイグレーションを実行
	csDB.AutoMigrate(&model.CardStatement{})
	
	csRepo = repository.NewCardStatementRepository(csDB)
	
	csTestUser = testutils.CreateTestUser(csDB)
	csOtherUser = testutils.CreateOtherUser(csDB)
}
