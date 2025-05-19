package csv_history_test

import (
	"monelog/model"
	"testing"
)

func TestGetCSVHistoriesByMonth(t *testing.T) {
	// テスト環境のセットアップ
	setupCSVHistoryTest(t)
	defer cleanupCSVHistoryTest(t)

	// テストデータの作成
	testHistories := []model.CSVHistory{
		{
			UserId:    csvHistoryTestUser,
			FileName:  "jan_2023.csv",
			CardType:  "rakuten",
			FileData:  []byte("jan data 1"),
			Year:      2023,
			Month:     1,
		},
		{
			UserId:    csvHistoryTestUser,
			FileName:  "jan_2023_2.csv",
			CardType:  "rakuten",
			FileData:  []byte("jan data 2"),
			Year:      2023,
			Month:     1,
		},
		{
			UserId:    csvHistoryTestUser,
			FileName:  "feb_2023.csv",
			CardType:  "rakuten",
			FileData:  []byte("feb data"),
			Year:      2023,
			Month:     2,
		},
		{
			UserId:    2, // 別のユーザー
			FileName:  "jan_2023_other.csv",
			CardType:  "rakuten",
			FileData:  []byte("other user data"),
			Year:      2023,
			Month:     1,
		},
	}

	// テストデータをデータベースに挿入
	for i := range testHistories {
		if err := csvHistoryDB.Create(&testHistories[i]).Error; err != nil {
			t.Fatalf("Failed to create test data: %v", err)
		}
	}

	// テストケース
	testCases := []struct {
		name          string
		userId        uint
		year          int
		month         int
		expectedCount int
		wantErr       bool
	}{
		{
			name:          "2023年1月のCSV履歴を取得",
			userId:        csvHistoryTestUser,
			year:          2023,
			month:         1,
			expectedCount: 2,
			wantErr:       false,
		},
		{
			name:          "2023年2月のCSV履歴を取得",
			userId:        csvHistoryTestUser,
			year:          2023,
			month:         2,
			expectedCount: 1,
			wantErr:       false,
		},
		{
			name:          "データが存在しない月のCSV履歴を取得",
			userId:        csvHistoryTestUser,
			year:          2023,
			month:         3,
			expectedCount: 0,
			wantErr:       false,
		},
		{
			name:          "別のユーザーの月別CSV履歴を取得",
			userId:        2,
			year:          2023,
			month:         1,
			expectedCount: 1,
			wantErr:       false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// テスト実行
			histories, err := csvHistoryRepo.GetCSVHistoriesByMonth(tc.userId, tc.year, tc.month)

			// 結果の検証
			if (err != nil) != tc.wantErr {
				t.Errorf("GetCSVHistoriesByMonth() error = %v, wantErr %v", err, tc.wantErr)
				return
			}

			if !tc.wantErr {
				if len(histories) != tc.expectedCount {
					t.Errorf("GetCSVHistoriesByMonth() returned %d histories, expected %d", len(histories), tc.expectedCount)
				}

				// 取得したデータの検証
				for _, history := range histories {
					if history.UserId != tc.userId {
						t.Errorf("GetCSVHistoriesByMonth() returned history with wrong UserId: got = %v, want = %v", history.UserId, tc.userId)
					}
					if history.Year != tc.year {
						t.Errorf("GetCSVHistoriesByMonth() returned history with wrong Year: got = %v, want = %v", history.Year, tc.year)
					}
					if history.Month != tc.month {
						t.Errorf("GetCSVHistoriesByMonth() returned history with wrong Month: got = %v, want = %v", history.Month, tc.month)
					}
				}

				// 順序の検証（created_at DESCでソートされているか）
				if len(histories) >= 2 && histories[0].CreatedAt.Before(histories[1].CreatedAt) {
					t.Errorf("GetCSVHistoriesByMonth() results are not sorted by created_at DESC")
				}
			}
		})
	}
}