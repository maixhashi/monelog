package csv_history

func (chu *csvHistoryUsecase) DeleteCSVHistory(userId uint, csvHistoryId uint) error {
	return chu.chr.DeleteCSVHistory(userId, csvHistoryId)
}