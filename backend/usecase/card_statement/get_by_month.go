package card_statement

import (
	"monelog/model"
)

func (csu *cardStatementUsecase) GetCardStatementsByMonth(request model.CardStatementByMonthRequest) ([]model.CardStatementResponse, error) {
	// Validate the request
	if err := csu.csv.ValidateCardStatementByMonthRequest(request); err != nil {
		return nil, err
	}
	
	// Get card statements for the specified month
	cardStatements, err := csu.csr.GetCardStatementsByMonth(request.UserId, request.Year, request.Month)
	if err != nil {
		return nil, err
	}
	
	// Convert to response format
	responses := make([]model.CardStatementResponse, len(cardStatements))
	for i, cardStatement := range cardStatements {
		responses[i] = cardStatement.ToResponse()
	}
	
	return responses, nil
}