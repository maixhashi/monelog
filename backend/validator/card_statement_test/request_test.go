package card_statement_test

import (
	"monelog/dto"
	"testing"
)

func TestCardStatementValidateRequest(t *testing.T) {
	setupCardStatementValidatorTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("すべてのフィールドが有効な場合", func(t *testing.T) {
			request := dto.CardStatementRequest{
				CardType: "rakuten",
				Year:     2023,
				Month:    4,
				UserId:   testUser.ID,
			}
			
			err := cardStatementValidator.ValidateCardStatementRequest(request)
			assertValidationResult(t, err, false, "有効なリクエスト")
		})
		
		t.Run("異なるカードタイプでも有効な場合", func(t *testing.T) {
			// mufgカードタイプ
			request1 := dto.CardStatementRequest{
				CardType: "mufg",
				Year:     2023,
				Month:    12,
				UserId:   testUser.ID,
			}
			
			err1 := cardStatementValidator.ValidateCardStatementRequest(request1)
			assertValidationResult(t, err1, false, "mufgカードタイプ")
			
			// eposカードタイプ
			request2 := dto.CardStatementRequest{
				CardType: "epos",
				Year:     2023,
				Month:    1,
				UserId:   testUser.ID,
			}
			
			err2 := cardStatementValidator.ValidateCardStatementRequest(request2)
			assertValidationResult(t, err2, false, "eposカードタイプ")
		})
		
		t.Run("異なるユーザーでも有効な場合", func(t *testing.T) {
			request := dto.CardStatementRequest{
				CardType: "rakuten",
				Year:     2023,
				Month:    4,
				UserId:   otherUser.ID,
			}
			
			err := cardStatementValidator.ValidateCardStatementRequest(request)
			assertValidationResult(t, err, false, "別ユーザーのリクエスト")
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("カードタイプが空の場合", func(t *testing.T) {
			request := dto.CardStatementRequest{
				CardType: "",
				Year:     2023,
				Month:    4,
				UserId:   testUser.ID,
			}
			
			err := cardStatementValidator.ValidateCardStatementRequest(request)
			assertValidationResult(t, err, true, "空のカードタイプ")
		})
		
		t.Run("無効なカードタイプの場合", func(t *testing.T) {
			request := dto.CardStatementRequest{
				CardType: "invalid_card",
				Year:     2023,
				Month:    4,
				UserId:   testUser.ID,
			}
			
			err := cardStatementValidator.ValidateCardStatementRequest(request)
			assertValidationResult(t, err, true, "無効なカードタイプ")
		})
		
		t.Run("年が指定されていない場合", func(t *testing.T) {
			request := dto.CardStatementRequest{
				CardType: "rakuten",
				Year:     0,
				Month:    4,
				UserId:   testUser.ID,
			}
			
			err := cardStatementValidator.ValidateCardStatementRequest(request)
			assertValidationResult(t, err, true, "年が未指定")
		})
		
		t.Run("月が範囲外の場合", func(t *testing.T) {
			// 月が小さすぎる
			request1 := dto.CardStatementRequest{
				CardType: "rakuten",
				Year:     2023,
				Month:    0,
				UserId:   testUser.ID,
			}
			
			err1 := cardStatementValidator.ValidateCardStatementRequest(request1)
			assertValidationResult(t, err1, true, "月が小さすぎる")
			
			// 月が大きすぎる
			request2 := dto.CardStatementRequest{
				CardType: "rakuten",
				Year:     2023,
				Month:    13,
				UserId:   testUser.ID,
			}
			
			err2 := cardStatementValidator.ValidateCardStatementRequest(request2)
			assertValidationResult(t, err2, true, "月が大きすぎる")
		})
	})
}