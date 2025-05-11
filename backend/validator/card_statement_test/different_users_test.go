package card_statement_test

import (
	"monelog/model"
	"testing"
)

func TestCardStatementValidateWithDifferentUsers(t *testing.T) {
	setupCardStatementValidatorTest()
	
	t.Run("異なるユーザーでのバリデーション", func(t *testing.T) {
		// ユーザー1のリクエスト
		request1 := model.CardStatementRequest{
			CardType: "rakuten",
			Year:     2023,
			Month:    4,
			UserId:   testUser.ID,
		}
		
		// ユーザー2のリクエスト
		request2 := model.CardStatementRequest{
			CardType: "mufg",
			Year:     2023,
			Month:    5,
			UserId:   otherUser.ID,
		}
		
		// バリデーションのテスト
		err1 := cardStatementValidator.ValidateCardStatementRequest(request1)
		assertValidationResult(t, err1, false, "ユーザー1のリクエスト")
		
		err2 := cardStatementValidator.ValidateCardStatementRequest(request2)
		assertValidationResult(t, err2, false, "ユーザー2のリクエスト")
	})
	
	t.Run("異なるユーザーでの保存リクエスト", func(t *testing.T) {
		validSummaries := createValidCardStatementSummaries()
		
		// ユーザー1の保存リクエスト
		saveRequest1 := model.CardStatementSaveRequest{
			CardType:       "rakuten",
			Year:           2023,
			Month:          4,
			UserId:         testUser.ID,
			CardStatements: validSummaries,
		}
		
		// ユーザー2の保存リクエスト
		saveRequest2 := model.CardStatementSaveRequest{
			CardType:       "mufg",
			Year:           2023,
			Month:          5,
			UserId:         otherUser.ID,
			CardStatements: validSummaries,
		}
		
		// バリデーションのテスト
		err1 := cardStatementValidator.ValidateCardStatementSaveRequest(saveRequest1)
		assertValidationResult(t, err1, false, "ユーザー1の保存リクエスト")
		
		err2 := cardStatementValidator.ValidateCardStatementSaveRequest(saveRequest2)
		assertValidationResult(t, err2, false, "ユーザー2の保存リクエスト")
	})
	
	t.Run("異なるユーザーでの月別リクエスト", func(t *testing.T) {
		// ユーザー1の月別リクエスト
		byMonthRequest1 := model.CardStatementByMonthRequest{
			Year:   2023,
			Month:  4,
			UserId: testUser.ID,
		}
		
		// ユーザー2の月別リクエスト
		byMonthRequest2 := model.CardStatementByMonthRequest{
			Year:   2023,
			Month:  5,
			UserId: otherUser.ID,
		}
		
		// バリデーションのテスト
		err1 := cardStatementValidator.ValidateCardStatementByMonthRequest(byMonthRequest1)
		assertValidationResult(t, err1, false, "ユーザー1の月別リクエスト")
		
		err2 := cardStatementValidator.ValidateCardStatementByMonthRequest(byMonthRequest2)
		assertValidationResult(t, err2, false, "ユーザー2の月別リクエスト")
	})
}