package csv_history_test

import (
	"monelog/model"
	"testing"
)

func TestCreateCSVHistory(t *testing.T) {
	// テスト環境のセットアップ
	setupCSVHistoryTest(t)
	defer cleanupCSVHistoryTest(t)

	// テストデータの準備
	testCases := []struct {
		name       string
		csvHistory model.CSVHistory
		wantErr    bool
	}{
		{
			name: "正常なCSV履歴の作成",
			csvHistory: model.CSVHistory{
				UserId:    csvHistoryTestUser,
				FileName:  "test_file.csv",
				CardType:  "rakuten",
				FileData:  []byte("test data"),
				Year:      2023,
				Month:     1,
			},
			wantErr: false,
		},
		{
			name: "FileDataがnilのCSV履歴の作成",
			csvHistory: model.CSVHistory{
				UserId:    csvHistoryTestUser,
				FileName:  "test_file.csv",
				CardType:  "rakuten",
				FileData:  nil, // FileDataがnil
				Year:      2023,
				Month:     1,
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// テスト実行
			err := csvHistoryRepo.CreateCSVHistory(&tc.csvHistory)

			// 結果の検証
			if (err != nil) != tc.wantErr {
				t.Errorf("CreateCSVHistory() error = %v, wantErr %v", err, tc.wantErr)
				return
			}

			// 作成されたデータの検証
			if !tc.wantErr {
				assertCSVHistoryExists(t, tc.csvHistory.ID, tc.csvHistory.FileName, tc.csvHistory.UserId)
			}
		})
	}
}