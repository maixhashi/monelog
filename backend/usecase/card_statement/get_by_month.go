package card_statement

import (
	"monelog/dto"
	"monelog/mapper"
)

func (csu *cardStatementUsecase) GetCardStatementsByMonth(request dto.CardStatementByMonthRequest) ([]dto.CardStatementResponse, error) {
	// Validate the request
	if err := csu.csv.ValidateCardStatementByMonthRequest(request); err != nil {
		return nil, err
	}
	
	// Get card statements for the specified month
	cardStatements, err := csu.csr.GetCardStatementsByMonth(request.UserId, request.Year, request.Month)
	if err != nil {
		return nil, err
	}
	
	// Convert to response format using the mapper
	return mapper.ToCardStatementResponseList(cardStatements), nil
}