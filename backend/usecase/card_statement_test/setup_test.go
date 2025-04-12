package card_statement_test

import (
	"monelog/model"
	"monelog/repository"
	"monelog/testutils"
	"monelog/usecase"
	"monelog/validator"
	"testing"

	"gorm.io/gorm"
)

// テスト用の共通変数
var (
	cardStatementDb        *gorm.DB
	cardStatementRepo      repository.ICardStatementRepository
	cardStatementValidator validator.ICardStatementValidator
	cardStatementUsecase   usecase.ICardStatementUsecase
	testUser               model.User
	otherUser              model.User
)

const nonExistentCardStatementID uint = 9999

// テスト前の共通セットアップ
func setupCardStatementUsecaseTest() {
	// テストごとにデータベースをクリーンアップ
	if cardStatementDb != nil {
		testutils.CleanupTestDB(cardStatementDb)
	} else {
		// 初回のみデータベース接続を作成
		cardStatementDb = testutils.SetupTestDB()
		cardStatementRepo = repository.NewCardStatementRepository(cardStatementDb)
		cardStatementValidator = validator.NewCardStatementValidator()
		cardStatementUsecase = usecase.NewCardStatementUsecase(cardStatementRepo, cardStatementValidator)
	}
	
	// テストユーザーを作成
	testUser = testutils.CreateTestUser(cardStatementDb)
	
	// 別のテストユーザーを作成
	otherUser = testutils.CreateOtherUser(cardStatementDb)
}

// テスト用のカード明細を作成
func createTestCardStatement(t *testing.T, description string, userId uint) model.CardStatement {
	cardStatement := model.CardStatement{
		Type:              "発生",
		StatementNo:       1,
		CardType:          "テストカード",
		Description:       description,
		UseDate:           "2023/01/01",
		PaymentDate:       "2023/02/27",
		PaymentMonth:      "2023年02月",
		Amount:            10000,
		TotalChargeAmount: 10000,
		ChargeAmount:      0,
		RemainingBalance:  10000,
		PaymentCount:      0,
		InstallmentCount:  1,
		AnnualRate:        0.0,
		MonthlyRate:       0.0,
		UserId:            userId,
	}
	
	result := cardStatementDb.Create(&cardStatement)
	if result.Error != nil {
		t.Fatalf("テストカード明細の作成に失敗しました: %v", result.Error)
	}
	
	return cardStatement
}
