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
	db                 *gorm.DB
	cardStatementRepo  repository.ICardStatementRepository
	cardStatementValidator validator.ICardStatementValidator
	cardStatementUsecase   usecase.ICardStatementUsecase
	testUser           model.User
	otherUser          model.User
)

const nonExistentCardStatementID uint = 9999

// テスト前の共通セットアップ
func setupCardStatementUsecaseTest() {
	// テストごとにデータベースをクリーンアップ
	if db != nil {
		testutils.CleanupTestDB(db)
	} else {
		// 初回のみデータベース接続を作成
		db = testutils.SetupTestDB()
		cardStatementRepo = repository.NewCardStatementRepository(db)
		cardStatementValidator = validator.NewCardStatementValidator()
		cardStatementUsecase = usecase.NewCardStatementUsecase(cardStatementRepo, cardStatementValidator)
	}
	
	// テストユーザーを作成
	testUser = testutils.CreateTestUser(db)
	
	// 別のテストユーザーを作成
	otherUser = testutils.CreateOtherUser(db)
}

// テスト用のカード明細を作成
func createTestCardStatement(t *testing.T, description string, userId uint, year int, month int) model.CardStatement {
	cardStatement := model.CardStatement{
		Type:              "発生",
		StatementNo:       1,
		CardType:          "楽天カード",
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
		Year:              year,
		Month:             month,
	}
	
	result := db.Create(&cardStatement)
	if result.Error != nil {
		t.Fatalf("テストカード明細の作成に失敗しました: %v", result.Error)
	}
	
	return cardStatement
}

// テスト用のカード明細サマリーを作成
func createTestCardStatementSummary(description string) model.CardStatementSummary {
	return model.CardStatementSummary{
		Type:              "発生",
		StatementNo:       1,
		CardType:          "楽天カード",
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
	}
}
