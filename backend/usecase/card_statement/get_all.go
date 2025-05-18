package card_statement

import (
	"monelog/model"
)

func (csu *cardStatementUsecase) GetAllCardStatements(userId uint) ([]model.CardStatementResponse, error) {
	cardStatements, err := csu.csr.GetAllCardStatements(userId)
	if err != nil {
		return nil, err
	}

	responses := make([]model.CardStatementResponse, len(cardStatements))
	for i, cardStatement := range cardStatements {
		responses[i] = cardStatement.ToResponse()
	}
	return responses, nil
}