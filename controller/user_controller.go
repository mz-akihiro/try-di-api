package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
	"try-di-api/model"
	"try-di-api/usecase"
)

type IUserController interface {
	LogIn(w http.ResponseWriter, r *http.Request)
}

type userController struct {
	uu usecase.IUserUsecase
}

func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &userController{uu}
}

func (uc *userController) LogIn(w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid JSON format")
		return
	}
	tokenString, err := uc.uu.LogIn(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Failed to create token")
		return
	}
	cookie := &http.Cookie{}
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(12 * time.Hour)
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	http.SetCookie(w, cookie)
}
