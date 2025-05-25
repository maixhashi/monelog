package card_statement

import (
	"monelog/dto"
	"monelog/mapper"
)

func (csu *cardStatementUsecase) GetCardStatementById(userId uint, cardStatementId uint) (dto.CardStatementResponse, error) {
	cardStatement, err := csu.csr.GetCardStatementById(userId, cardStatementId)
	if err != nil {
		return dto.CardStatementResponse{}, err
	}
	return mapper.ToCardStatementResponse(&cardStatement), nil
}