package csv_history_test

import (
	"monelog/model"
	"testing"
)

func TestGetAllCSVHistories(t *testing.T) {
	// テスト環境のセットアップ
	setupCSVHistoryTest(t)
	defer cleanupCSVHistoryTest(t)

	// テストデータの作成
	testHistories := []model.CSVHistory{
		{
			UserId:    csvHistoryTestUser,
			FileName:  "test_file1.csv",
			CardType:  "rakuten",
			FileData:  []byte("test data 1"),
			Year:      2023,
			Month:     1,
		},
		{
			UserId:    csvHistoryTestUser,
			FileName:  "test_file2.csv",
			CardType:  "rakuten",
			FileData:  []byte("test data 2"),
			Year:      2023,
			Month:     2,
		},
		{
			UserId:    2, // 別のユーザー
			FileName:  "other_user.csv",
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

	// テスト実行
	histories, err := csvHistoryRepo.GetAllCSVHistories(csvHistoryTestUser)
	if err != nil {
		t.Fatalf("GetAllCSVHistories() error = %v", err)
	}

	// 結果の検証
	expectedCount := 2 // csvHistoryTestUserに属するデータのみ
	if !validateCSVHistories(t, histories, expectedCount) {
		t.Errorf("GetAllCSVHistories() returned invalid data")
	}

	// 順序の検証（created_at DESCでソートされているか）
	if len(histories) >= 2 && histories[0].CreatedAt.Before(histories[1].CreatedAt) {
		t.Errorf("GetAllCSVHistories() results are not sorted by created_at DESC")
	}
}