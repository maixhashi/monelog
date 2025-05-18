package card_statement

import (
	"monelog/model"
)

func (csu *cardStatementUsecase) SaveCardStatements(request model.CardStatementSaveRequest) ([]model.CardStatementResponse, error) {
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
		cardStatements[i] = summary.ToModel(request.UserId, request.Year, request.Month)
	}

	if err := csu.csr.CreateCardStatements(cardStatements); err != nil {
		return nil, err
	}

	// レスポンスを作成
	responses := make([]model.CardStatementResponse, len(cardStatements))
	for i, cardStatement := range cardStatements {
		responses[i] = cardStatement.ToResponse()
	}

	return responses, nil
}