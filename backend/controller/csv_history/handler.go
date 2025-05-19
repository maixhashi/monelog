package csv_history

import (
	"monelog/model"
	"monelog/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	chu usecase.ICSVHistoryUsecase
}

func NewHandler(chu usecase.ICSVHistoryUsecase) *Handler {
	return &Handler{chu}
}

// ユーザーIDをトークンから取得するヘルパー関数
func getUserIdFromToken(c echo.Context) (uint, error) {
	user := c.Get("user")
	if user == nil {
		return 0, echo.NewHTTPError(http.StatusUnauthorized, "ユーザー情報が見つかりません")
	}
	
	// ユーザーの型を確認して適切に処理
	switch u := user.(type) {
	case uint:
		return u, nil
	case model.User:
		return u.ID, nil
	default:
		return 0, echo.NewHTTPError(http.StatusInternalServerError, "ユーザー情報の型が不正です")
	}
}