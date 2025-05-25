package mapper

import (
	"monelog/dto"
	"monelog/mapper/card_statement"
	"monelog/model"
)

var (
	modelToDTOMapper = card_statement.NewModelToDTO()
	dtoToModelMapper = card_statement.NewDTOToModel()
)

// ToCardStatementResponse モデルからレスポンスDTOへの変換
func ToCardStatementResponse(cs *model.CardStatement) dto.CardStatementResponse {
	return modelToDTOMapper.ToResponse(cs)
}

// ToCardStatementModel サマリーDTOからモデルへの変換
func ToCardStatementModel(css *dto.CardStatementSummary, userId uint, year int, month int) model.CardStatement {
	return dtoToModelMapper.FromSummary(css, userId, year, month)
}

// ToCardStatementResponseList モデルのスライスからレスポンスDTOのスライスへの変換
func ToCardStatementResponseList(statements []model.CardStatement) []dto.CardStatementResponse {
	return modelToDTOMapper.ToResponseList(statements)
}