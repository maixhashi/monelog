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

func (csu *cardStatementUsecase) GetCardStatementById(userId uint, cardStatementId uint) (model.CardStatementResponse, error) {
	cardStatement, err := csu.csr.GetCardStatementById(userId, cardStatementId)
	if err != nil {
		return model.CardStatementResponse{}, err
	}
	return cardStatement.ToResponse(), nil
}

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