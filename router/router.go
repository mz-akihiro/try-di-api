package router

import (
	"net/http"
	"try-di-api/controller"
)

func NewRouter(uc controller.IUserController) *http.ServeMux {
	s := http.NewServeMux()
	s.HandleFunc("/login", uc.LogIn)
	return s
}
