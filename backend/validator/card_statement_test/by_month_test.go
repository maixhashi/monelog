package card_statement_test

import (
	"monelog/dto"
	"testing"
)

func TestCardStatementValidateByMonthRequest(t *testing.T) {
	setupCardStatementValidatorTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("すべてのフィールドが有効な場合", func(t *testing.T) {
			request := dto.CardStatementByMonthRequest{
				Year:   2023,
				Month:  4,
				UserId: testUser.ID,
			}
			
			err := cardStatementValidator.ValidateCardStatementByMonthRequest(request)
			assertValidationResult(t, err, false, "有効な月別リクエスト")
		})
		
		t.Run("最小値の月でも有効な場合", func(t *testing.T) {
			request := dto.CardStatementByMonthRequest{
				Year:   2023,
				Month:  1,
				UserId: testUser.ID,
			}
			
			err := cardStatementValidator.ValidateCardStatementByMonthRequest(request)
			assertValidationResult(t, err, false, "1月のリクエスト")
		})
		
		t.Run("最大値の月でも有効な場合", func(t *testing.T) {
			request := dto.CardStatementByMonthRequest{
				Year:   2023,
				Month:  12,
				UserId: testUser.ID,
			}
			
			err := cardStatementValidator.ValidateCardStatementByMonthRequest(request)
			assertValidationResult(t, err, false, "12月のリクエスト")
		})
		
		t.Run("異なるユーザーでも有効な場合", func(t *testing.T) {
			request := dto.CardStatementByMonthRequest{
				Year:   2023,
				Month:  4,
				UserId: otherUser.ID,
			}
			
			err := cardStatementValidator.ValidateCardStatementByMonthRequest(request)
			assertValidationResult(t, err, false, "別ユーザーのリクエスト")
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("年が指定されていない場合", func(t *testing.T) {
			request := dto.CardStatementByMonthRequest{
				Year:   0,
				Month:  4,
				UserId: testUser.ID,
			}
			
			err := cardStatementValidator.ValidateCardStatementByMonthRequest(request)
			assertValidationResult(t, err, true, "年が未指定")
		})
		
		t.Run("月が範囲外の場合", func(t *testing.T) {
			// 月が小さすぎる
			request1 := dto.CardStatementByMonthRequest{
				Year:   2023,
				Month:  0,
				UserId: testUser.ID,
			}
			
			err1 := cardStatementValidator.ValidateCardStatementByMonthRequest(request1)
			assertValidationResult(t, err1, true, "月が小さすぎる")
			
			// 月が大きすぎる
			request2 := dto.CardStatementByMonthRequest{
				Year:   2023,
				Month:  13,
				UserId: testUser.ID,
			}
			
			err2 := cardStatementValidator.ValidateCardStatementByMonthRequest(request2)
			assertValidationResult(t, err2, true, "月が大きすぎる")
		})
	})
}