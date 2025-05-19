package csv_history

import (
	"monelog/model"
)

func (chu *csvHistoryUsecase) GetAllCSVHistories(userId uint) ([]model.CSVHistoryResponse, error) {
	csvHistories, err := chu.chr.GetAllCSVHistories(userId)
	if err != nil {
		return nil, err
	}

	responses := make([]model.CSVHistoryResponse, len(csvHistories))
	for i, csvHistory := range csvHistories {
		responses[i] = csvHistory.ToResponse()
	}
	return responses, nil
}