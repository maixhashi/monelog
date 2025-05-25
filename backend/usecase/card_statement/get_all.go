package card_statement

import (
	"monelog/dto"
	"monelog/mapper"
)

func (csu *cardStatementUsecase) GetAllCardStatements(userId uint) ([]dto.CardStatementResponse, error) {
	cardStatements, err := csu.csr.GetAllCardStatements(userId)
	if err != nil {
		return nil, err
	}

	return mapper.ToCardStatementResponseList(cardStatements), nil
}