package card_statement_test

import (
	"monelog/model"
	"testing"
)

func TestCardStatementValidateSaveRequest(t *testing.T) {
	setupCardStatementValidatorTest()
	
	// 有効なカード明細サマリーを取得
	validSummaries := createValidCardStatementSummaries()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("すべてのフィールドが有効な場合", func(t *testing.T) {
			request := model.CardStatementSaveRequest{
				CardType:       "rakuten",
				Year:           2023,
				Month:          4,
				UserId:         testUser.ID,
				CardStatements: validSummaries,
			}
			
			err := cardStatementValidator.ValidateCardStatementSaveRequest(request)
			assertValidationResult(t, err, false, "有効な保存リクエスト")
		})
		
		t.Run("異なるカードタイプでも有効な場合", func(t *testing.T) {
			// mufgカードタイプ
			request1 := model.CardStatementSaveRequest{
				CardType:       "mufg",
				Year:           2023,
				Month:          12,
				UserId:         testUser.ID,
				CardStatements: validSummaries,
			}
			
			err1 := cardStatementValidator.ValidateCardStatementSaveRequest(request1)
			assertValidationResult(t, err1, false, "mufgカードタイプ")
			
			// eposカードタイプ
			request2 := model.CardStatementSaveRequest{
				CardType:       "epos",
				Year:           2023,
				Month:          1,
				UserId:         testUser.ID,
				CardStatements: validSummaries,
			}
			
			err2 := cardStatementValidator.ValidateCardStatementSaveRequest(request2)
			assertValidationResult(t, err2, false, "eposカードタイプ")
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("カードタイプが空の場合", func(t *testing.T) {
			request := model.CardStatementSaveRequest{
				CardType:       "",
				Year:           2023,
				Month:          4,
				UserId:         testUser.ID,
				CardStatements: validSummaries,
			}
			
			err := cardStatementValidator.ValidateCardStatementSaveRequest(request)
			assertValidationResult(t, err, true, "空のカードタイプ")
		})
		
		t.Run("無効なカードタイプの場合", func(t *testing.T) {
			request := model.CardStatementSaveRequest{
				CardType:       "invalid_card",
				Year:           2023,
				Month:          4,
				UserId:         testUser.ID,
				CardStatements: validSummaries,
			}
			
			err := cardStatementValidator.ValidateCardStatementSaveRequest(request)
			assertValidationResult(t, err, true, "無効なカードタイプ")
		})
		
		t.Run("年が指定されていない場合", func(t *testing.T) {
			request := model.CardStatementSaveRequest{
				CardType:       "rakuten",
				Year:           0,
				Month:          4,
				UserId:         testUser.ID,
				CardStatements: validSummaries,
			}
			
			err := cardStatementValidator.ValidateCardStatementSaveRequest(request)
			assertValidationResult(t, err, true, "年が未指定")
		})
		
		t.Run("月が範囲外の場合", func(t *testing.T) {
			// 月が小さすぎる
			request1 := model.CardStatementSaveRequest{
				CardType:       "rakuten",
				Year:           2023,
				Month:          0,
				UserId:         testUser.ID,
				CardStatements: validSummaries,
			}
			
			err1 := cardStatementValidator.ValidateCardStatementSaveRequest(request1)
			assertValidationResult(t, err1, true, "月が小さすぎる")
			
			// 月が大きすぎる
			request2 := model.CardStatementSaveRequest{
				CardType:       "rakuten",
				Year:           2023,
				Month:          13,
				UserId:         testUser.ID,
				CardStatements: validSummaries,
			}
			
			err2 := cardStatementValidator.ValidateCardStatementSaveRequest(request2)
			assertValidationResult(t, err2, true, "月が大きすぎる")
		})
		
		t.Run("カード明細が空の場合", func(t *testing.T) {
			// 空の配列
			request1 := model.CardStatementSaveRequest{
				CardType:       "rakuten",
				Year:           2023,
				Month:          4,
				UserId:         testUser.ID,
				CardStatements: []model.CardStatementSummary{},
			}
			
			err1 := cardStatementValidator.ValidateCardStatementSaveRequest(request1)
			assertValidationResult(t, err1, true, "空のカード明細配列")
			
			// nil
			request2 := model.CardStatementSaveRequest{
				CardType:       "rakuten",
				Year:           2023,
				Month:          4,
				UserId:         testUser.ID,
				CardStatements: nil,
			}
			
			err2 := cardStatementValidator.ValidateCardStatementSaveRequest(request2)
			assertValidationResult(t, err2, true, "nilのカード明細")
		})
	})
}