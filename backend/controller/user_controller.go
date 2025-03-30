package controller

import (
	"monelog/model"
	"monelog/usecase"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

type IUserController interface {
	SignUp(c echo.Context) error
	LogIn(c echo.Context) error
	LogOut(c echo.Context) error
	CsrfToken(c echo.Context) error
}

type userController struct {
	uu usecase.IUserUsecase
}

func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &userController{uu}
}

// SignUp 新規ユーザー登録
// @Summary 新規ユーザー登録
// @Description 新しいユーザーアカウントを作成する
// @Tags users
// @Accept json
// @Produce json
// @Param user body model.UserSignupRequest true "ユーザー登録情報"
// @Success 201 {object} model.UserResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /signup [post]
func (uc *userController) SignUp(c echo.Context) error {
	req := model.UserSignupRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	
	userRes, err := uc.uu.SignUp(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	
	return c.JSON(http.StatusCreated, userRes)
}

// LogIn ユーザーログイン
// @Summary ユーザーログイン
// @Description 既存ユーザーのログイン処理
// @Tags users
// @Accept json
// @Produce json
// @Param user body model.UserLoginRequest true "ログイン情報"
// @Success 200 {string} string "OK"
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /login [post]
func (uc *userController) LogIn(c echo.Context) error {
	req := model.UserLoginRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	
	tokenString, err := uc.uu.Login(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
	
	return c.NoContent(http.StatusOK)
}

// LogOut ユーザーログアウト
// @Summary ユーザーログアウト
// @Description ユーザーのログアウト処理
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {string} string "OK"
// @Router /logout [post]
func (uc *userController) LogOut(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Expires = time.Now()
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
	
	return c.NoContent(http.StatusOK)
}

// CsrfToken CSRFトークン取得
// @Summary CSRFトークン取得
// @Description CSRFトークンを取得する
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} model.CsrfTokenResponse
// @Router /csrf-token [get]
func (uc *userController) CsrfToken(c echo.Context) error {
	token := c.Get("csrf").(string)
	return c.JSON(http.StatusOK, map[string]string{
		"csrf_token": token,
	})
}