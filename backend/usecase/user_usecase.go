package usecase

import (
	"monelog/model"
	"monelog/repository"
	"monelog/validator"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	SignUp(req model.UserSignupRequest) (*model.UserResponse, error)
	Login(req model.UserLoginRequest) (string, error)
	VerifyAuth(userId uint) (*model.UserResponse, error) // 追加
}

type userUsecase struct {
	ur repository.IUserRepository
	uv validator.IUserValidator
}

func NewUserUsecase(ur repository.IUserRepository, uv validator.IUserValidator) IUserUsecase {
	return &userUsecase{ur, uv}
}

func (uu *userUsecase) SignUp(req model.UserSignupRequest) (*model.UserResponse, error) {
	// リクエストをUserモデルに変換
	user := req.ToUser()
	
	// バリデーション
	if err := uu.uv.UserValidate(user); err != nil {
		return nil, err
	}
	
	// パスワードハッシュ化
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return nil, err
	}
	user.Password = string(hash)
	
	// ユーザー作成
	if err := uu.ur.CreateUser(&user); err != nil {
		return nil, err
	}
	
	// レスポンス作成
	response := user.ToUserResponse()
	return &response, nil
}

func (uu *userUsecase) Login(req model.UserLoginRequest) (string, error) {
	// バリデーション用にUserに変換
	user := model.User{Email: req.Email, Password: req.Password}
	if err := uu.uv.UserValidate(user); err != nil {
		return "", err
	}
	
	// ユーザー取得
	storedUser, err := uu.ur.GetUserByEmail(req.Email)
	if err != nil {
		return "", err
	}
	
	// パスワード検証
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(req.Password)); err != nil {
		return "", err
	}
	
	// JWTトークン生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})
	
	return token.SignedString([]byte(os.Getenv("SECRET")))
}

// VerifyAuth は認証状態を確認し、ユーザー情報を返します
func (uu *userUsecase) VerifyAuth(userId uint) (*model.UserResponse, error) {
	// ユーザーIDからユーザー情報を取得
	user, err := uu.ur.GetUserById(userId)
	if err != nil {
		return nil, err
	}
	
	// レスポンス作成
	response := user.ToUserResponse()
	return &response, nil
}