package card_statement

import (
	"monelog/dto"
	"monelog/mapper"
	"monelog/model"
)

func (csu *cardStatementUsecase) SaveCardStatements(request dto.CardStatementSaveRequest) ([]dto.CardStatementResponse, error) {
	if err := csu.csv.ValidateCardStatementSaveRequest(request); err != nil {
		return nil, err
	}

	// 既存のデータを削除
	if err := csu.csr.DeleteCardStatements(request.UserId); err != nil {
		return nil, err
	}

	// 新しいデータを保存
	cardStatements := make([]model.CardStatement, len(request.CardStatements))
	for i, summary := range request.CardStatements {
		// 年月を引数として渡す
		cardStatements[i] = mapper.ToCardStatementModel(&summary, request.UserId, request.Year, request.Month)
	}

	if err := csu.csr.CreateCardStatements(cardStatements); err != nil {
		return nil, err
	}

	// レスポンスを作成
	return mapper.ToCardStatementResponseList(cardStatements), nil
}