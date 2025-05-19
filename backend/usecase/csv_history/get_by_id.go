package csv_history

import (
	"monelog/model"
)

func (chu *csvHistoryUsecase) GetCSVHistoryById(userId uint, csvHistoryId uint) (model.CSVHistoryDetailResponse, error) {
	csvHistory, err := chu.chr.GetCSVHistoryById(userId, csvHistoryId)
	if err != nil {
		return model.CSVHistoryDetailResponse{}, err
	}
	return csvHistory.ToDetailResponse(), nil
}