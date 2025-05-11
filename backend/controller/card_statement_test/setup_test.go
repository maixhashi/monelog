package card_statement_test

import (
	"monelog/controller"
	"monelog/model"
	"monelog/repository"
	"monelog/testutils"
	"monelog/usecase"
	"monelog/validator"

	"gorm.io/gorm"
)

// テスト用の共通変数
var (
	cardStatementDB          *gorm.DB
	cardStatementRepo        repository.ICardStatementRepository
	csvHistoryRepo           repository.ICSVHistoryRepository
	cardStatementValidator   validator.ICardStatementValidator
	csvHistoryValidator      validator.ICSVHistoryValidator
	cardStatementUsecase     usecase.ICardStatementUsecase
	csvHistoryUsecase        usecase.ICSVHistoryUsecase
	cardStatementController  controller.ICardStatementController
	cardStatementTestUser    model.User
	cardStatementOtherUser   model.User
	nonExistentCardStatementID uint = 9999
)

// テストセットアップ関数
func setupCardStatementControllerTest() {
	// テストごとにデータベースをクリーンアップ
	if cardStatementDB != nil {
		testutils.CleanupTestDB(cardStatementDB)
	} else {
		// 初回のみデータベース接続を作成
		cardStatementDB = testutils.SetupTestDB()
		cardStatementRepo = repository.NewCardStatementRepository(cardStatementDB)
		csvHistoryRepo = repository.NewCSVHistoryRepository(cardStatementDB)
		cardStatementValidator = validator.NewCardStatementValidator()
		csvHistoryValidator = validator.NewCSVHistoryValidator()
		cardStatementUsecase = usecase.NewCardStatementUsecase(cardStatementRepo, cardStatementValidator)
		csvHistoryUsecase = usecase.NewCSVHistoryUsecase(csvHistoryRepo, csvHistoryValidator)
		cardStatementController = controller.NewCardStatementController(cardStatementUsecase, csvHistoryUsecase)
	}
	
	// テストユーザーを作成
	cardStatementTestUser = testutils.CreateTestUser(cardStatementDB)
	cardStatementOtherUser = testutils.CreateOtherUser(cardStatementDB)
}
