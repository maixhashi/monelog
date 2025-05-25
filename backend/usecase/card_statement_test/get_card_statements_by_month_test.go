package card_statement_test

import (
	"monelog/dto"
	"testing"
)

func TestCardStatementUsecase_GetCardStatementsByMonth(t *testing.T) {
	setupCardStatementUsecaseTest()
	
	// テストデータの作成 - 異なる年月の明細
	createTestCardStatement(t, "1月の明細1", testUser.ID, 2023, 1)
	createTestCardStatement(t, "1月の明細2", testUser.ID, 2023, 1)
	createTestCardStatement(t, "2月の明細", testUser.ID, 2023, 2)
	createTestCardStatement(t, "前年の明細", testUser.ID, 2022, 1)
	createTestCardStatement(t, "他ユーザーの1月明細", otherUser.ID, 2023, 1)
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("指定した年月の明細のみを取得する", func(t *testing.T) {
			request := dto.CardStatementByMonthRequest{
				Year:   2023,
				Month:  1,
				UserId: testUser.ID,
			}
			
			t.Logf("ユーザーID %d の2023年1月の明細を取得します", testUser.ID)
			
			responses, err := cardStatementUsecase.GetCardStatementsByMonth(request)
			
			if err != nil {
				t.Errorf("GetCardStatementsByMonth() error = %v", err)
			}
			
			if !validateCardStatementResponses(t, responses, 2) {
				return
			}
			
			// 明細の説明の確認
			descriptions := make(map[string]bool)
			for _, response := range responses {
				descriptions[response.Description] = true
				t.Logf("取得した明細: ID=%d, Description=%s, Year=%d, Month=%d", 
					response.ID, response.Description, response.Year, response.Month)
				
				// 年月が正しいことを確認
				if response.Year != 2023 || response.Month != 1 {
					t.Errorf("年月が一致しません: got=%d年%d月, want=2023年1月", response.Year, response.Month)
				}
			}
			
			if !descriptions["1月の明細1"] || !descriptions["1月の明細2"] {
				t.Errorf("期待した明細が結果に含まれていません: %v", responses)
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("バリデーションエラー - 無効な月", func(t *testing.T) {
			request := dto.CardStatementByMonthRequest{
				Year:   2023,
				Month:  13, // 無効な月
				UserId: testUser.ID,
			}
			
			t.Logf("無効な月(13月)を指定して明細を取得しようとします")
			
			_, err := cardStatementUsecase.GetCardStatementsByMonth(request)
			
			if err == nil {
				t.Error("無効な月を指定したときにエラーが返されませんでした")
			} else {
				t.Logf("期待通りエラーが返されました: %v", err)
			}
		})
	})
}
