package csv_history

import (
	"monelog/model"
)

func (chu *csvHistoryUsecase) GetCSVHistoriesByMonth(userId uint, year int, month int) ([]model.CSVHistoryResponse, error) {
	// リポジトリから指定された年月のCSV履歴を取得
	csvHistories, err := chu.chr.GetCSVHistoriesByMonth(userId, year, month)
	if err != nil {
		return nil, err
	}

	// レスポンス形式に変換
	responses := make([]model.CSVHistoryResponse, len(csvHistories))
	for i, csvHistory := range csvHistories {
		responses[i] = csvHistory.ToResponse()
	}
	
	return responses, nil
}