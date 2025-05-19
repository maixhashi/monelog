package csv_history_test

import (
	"monelog/model"
	"testing"
)

func TestGetCSVHistoryById(t *testing.T) {
	// テスト環境のセットアップ
	setupCSVHistoryTest(t)
	defer cleanupCSVHistoryTest(t)

	// テストデータの作成
	testHistory := model.CSVHistory{
		UserId:    csvHistoryTestUser,
		FileName:  "test_file.csv",
		CardType:  "rakuten",
		FileData:  []byte("test data"),
		Year:      2023,
		Month:     1,
	}

	otherUserHistory := model.CSVHistory{
		UserId:    2, // 別のユーザー
		FileName:  "other_user.csv",
		CardType:  "rakuten",
		FileData:  []byte("other user data"),
		Year:      2023,
		Month:     1,
	}

	// テストデータをデータベースに挿入
	if err := csvHistoryDB.Create(&testHistory).Error; err != nil {
		t.Fatalf("Failed to create test data: %v", err)
	}

	if err := csvHistoryDB.Create(&otherUserHistory).Error; err != nil {
		t.Fatalf("Failed to create test data: %v", err)
	}

	// テストケース
	testCases := []struct {
		name        string
		userId      uint
		historyId   uint
		wantErr     bool
		expectedMsg string
	}{
		{
			name:      "存在するCSV履歴の取得",
			userId:    csvHistoryTestUser,
			historyId: testHistory.ID,
			wantErr:   false,
		},
		{
			name:        "存在しないCSV履歴の取得",
			userId:      csvHistoryTestUser,
			historyId:   9999,
			wantErr:     true,
			expectedMsg: "record not found",
		},
		{
			name:        "他のユーザーのCSV履歴の取得",
			userId:      csvHistoryTestUser,
			historyId:   otherUserHistory.ID,
			wantErr:     true,
			expectedMsg: "record not found",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// テスト実行
			history, err := csvHistoryRepo.GetCSVHistoryById(tc.userId, tc.historyId)

			// 結果の検証
			if (err != nil) != tc.wantErr {
				t.Errorf("GetCSVHistoryById() error = %v, wantErr %v", err, tc.wantErr)
				return
			}

			if tc.wantErr {
				if err != nil && tc.expectedMsg != "" && err.Error() != tc.expectedMsg {
					t.Errorf("GetCSVHistoryById() error message = %v, expected %v", err.Error(), tc.expectedMsg)
				}
				return
			}

			// 取得したデータの検証
			if history.ID != tc.historyId {
				t.Errorf("GetCSVHistoryById() returned wrong ID: got = %v, want = %v", history.ID, tc.historyId)
			}

			if history.UserId != tc.userId {
				t.Errorf("GetCSVHistoryById() returned wrong UserId: got = %v, want = %v", history.UserId, tc.userId)
			}
		})
	}
}