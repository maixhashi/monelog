package card_statement

import (
	"monelog/usecase"
)

type Handler struct {
	csu usecase.ICardStatementUsecase
	chu usecase.ICSVHistoryUsecase
}

func NewHandler(csu usecase.ICardStatementUsecase, chu usecase.ICSVHistoryUsecase) *Handler {
	return &Handler{
		csu: csu,
		chu: chu,
	}
}