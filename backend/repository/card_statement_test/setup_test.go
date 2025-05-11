package card_statement_test

import (
	"monelog/model"
	"monelog/repository"
	"monelog/testutils"
	"gorm.io/gorm"
)

var (
	cardStatementDB *gorm.DB
	cardStatementRepo repository.ICardStatementRepository
	cardStatementTestUser model.User
	cardStatementOtherUser model.User
	nonExistentCardStatementID uint = 9999
)

func setupCardStatementTest() {
	cardStatementDB = testutils.SetupTestDB()
	cardStatementRepo = repository.NewCardStatementRepository(cardStatementDB)
	
	cardStatementTestUser = testutils.CreateTestUser(cardStatementDB)
	cardStatementOtherUser = testutils.CreateOtherUser(cardStatementDB)
}
