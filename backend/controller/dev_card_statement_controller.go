package controller

import (
	"monelog/model"
	"monelog/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IDevCardStatementController interface {
	DeleteAllCardStatements(c echo.Context) error
}

type devCardStatementController struct {
	dcsu usecase.IDevCardStatementUsecase
}

func NewDevCardStatementController(dcsu usecase.IDevCardStatementUsecase) IDevCardStatementController {
	return &devCardStatementController{dcsu}
}

// DeleteAllCardStatements 開発環境限定で全カード明細を削除
// @Summary 開発環境限定で全カード明細を削除
// @Description 開発環境限定で全カード明細レコードを削除する
// @Tags dev
// @Accept json
// @Produce json
// @Success 200 {object} model.DevCardStatementResponse
// @Failure 403 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /dev/card-statements/delete-all [post]
func (dcsc *devCardStatementController) DeleteAllCardStatements(c echo.Context) error {
	// 空のリクエストを作成
	request := model.DevCardStatementRequest{}

	response, err := dcsc.dcsu.DeleteAllCardStatements(request)
	if err != nil {
		// エラーの種類に応じてステータスコードを変える
		if err.Error() == "this operation is only allowed in development environment" {
			return c.JSON(http.StatusForbidden, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, response)
}
