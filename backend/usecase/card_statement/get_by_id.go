package card_statement

import (
	"monelog/model"
)

func (csu *cardStatementUsecase) GetCardStatementById(userId uint, cardStatementId uint) (model.CardStatementResponse, error) {
	cardStatement, err := csu.csr.GetCardStatementById(userId, cardStatementId)
	if err != nil {
		return model.CardStatementResponse{}, err
	}
	return cardStatement.ToResponse(), nil
}