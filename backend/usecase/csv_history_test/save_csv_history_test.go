package csv_history_test

import (
	"monelog/model"
	"testing"
)

// テスト用のCSV履歴保存リクエストを作成
func createSaveCSVHistoryRequest(userId uint) model.CSVHistorySaveRequest {
	return model.CSVHistorySaveRequest{
		FileName: "test.csv",
		CardType: "楽天カード",
		UserId:   userId,
		Year:     2023,
		Month:    1,
	}
}

func TestCSVHistoryUsecase_SaveCSVHistory(t *testing.T) {
	t.Skip("このテストはファイルアップロードをモックするのが難しいため、スキップします")
	
	// 注意: このテストは実際のファイルアップロードをシミュレートするのが難しいため、
	// 実際の実装では、リポジトリとバリデータをモックして単体テストを行うか、
	// 統合テストとして実装することをお勧めします。
}