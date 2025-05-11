package card_statement_test

import (
	"monelog/model"
	"monelog/testutils"
	"monelog/validator"

	"gorm.io/gorm"
)

// テスト用の共通変数
var (
	db                     *gorm.DB
	cardStatementValidator validator.ICardStatementValidator
	testUser               model.User
	otherUser              model.User
)

// テスト前の共通セットアップ
func setupCardStatementValidatorTest() {
	// テストごとにデータベースをクリーンアップ
	if db != nil {
		testutils.CleanupTestDB(db)
	} else {
		// 初回のみデータベース接続を作成
		db = testutils.SetupTestDB()
		cardStatementValidator = validator.NewCardStatementValidator()
	}
	
	// テストユーザーを作成
	testUser = testutils.CreateTestUser(db)
	
	// 別のテストユーザーを作成
	otherUser = testutils.CreateOtherUser(db)
}