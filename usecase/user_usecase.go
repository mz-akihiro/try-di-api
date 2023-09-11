package usecase

import (
	"log"
	"os"
	"time"
	"try-di-api/model"
	"try-di-api/repository"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	LogIn(user model.User) (string, error)
}

type userUsecase struct {
	ur repository.IUserRepository
}

func NewUserUsecase(ur repository.IUserRepository) IUserUsecase {
	return &userUsecase{ur}
}

func (uu *userUsecase) LogIn(user model.User) (string, error) {
	var storedPassword string
	var userId int
	if err := uu.ur.GetUserByEmail(&storedPassword, &userId, user.Email); err != nil {
		log.Println("Failed GetUser")
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(user.Password)); err != nil {
		log.Println("Failed Compare")
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,                               // 後々taskを追加・削除する際などに使う用に、ペイロードにuserIdを入れる
		"exp":     time.Now().Add(time.Hour * 6).Unix(), // jwtトークンの有効期限(6時間)
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
