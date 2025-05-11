package card_statement_test

import (
	"monelog/model"
	"testing"
)

func TestCardStatementValidatePreviewRequest(t *testing.T) {
	setupCardStatementValidatorTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("すべてのフィールドが有効な場合", func(t *testing.T) {
			request := model.CardStatementPreviewRequest{
				CardType: "rakuten",
				Year:     2023,
				Month:    4,
				UserId:   testUser.ID,
			}
			
			err := cardStatementValidator.ValidateCardStatementPreviewRequest(request)
			assertValidationResult(t, err, false, "有効なプレビューリクエスト")
		})
		
		t.Run("年月が指定されていなくても有効な場合", func(t *testing.T) {
			request := model.CardStatementPreviewRequest{
				CardType: "rakuten",
				UserId:   testUser.ID,
			}
			
			err := cardStatementValidator.ValidateCardStatementPreviewRequest(request)
			assertValidationResult(t, err, false, "年月なしのプレビューリクエスト")
		})
		
		t.Run("月が0でも有効な場合（任意項目）", func(t *testing.T) {
			request := model.CardStatementPreviewRequest{
				CardType: "rakuten",
				Year:     2023,
				Month:    0,
				UserId:   testUser.ID,
			}
			
			err := cardStatementValidator.ValidateCardStatementPreviewRequest(request)
			assertValidationResult(t, err, false, "月が0のプレビューリクエスト")
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("カードタイプが空の場合", func(t *testing.T) {
			request := model.CardStatementPreviewRequest{
				CardType: "",
				Year:     2023,
				Month:    4,
				UserId:   testUser.ID,
			}
			
			err := cardStatementValidator.ValidateCardStatementPreviewRequest(request)
			assertValidationResult(t, err, true, "空のカードタイプ")
		})
		
		t.Run("無効なカードタイプの場合", func(t *testing.T) {
			request := model.CardStatementPreviewRequest{
				CardType: "invalid_card",
				Year:     2023,
				Month:    4,
				UserId:   testUser.ID,
			}
			
			err := cardStatementValidator.ValidateCardStatementPreviewRequest(request)
			assertValidationResult(t, err, true, "無効なカードタイプ")
		})
		
		t.Run("月が範囲外の場合", func(t *testing.T) {
			request := model.CardStatementPreviewRequest{
				CardType: "rakuten",
				Year:     2023,
				Month:    13,
				UserId:   testUser.ID,
			}
			
			err := cardStatementValidator.ValidateCardStatementPreviewRequest(request)
			assertValidationResult(t, err, true, "月が大きすぎる")
		})
	})
}