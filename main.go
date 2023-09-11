package main

import (
	"log"
	"net/http"
	"try-di-api/controller"
	"try-di-api/db"
	"try-di-api/repository"
	"try-di-api/router"
	"try-di-api/usecase"
)

func main() {
	db := db.Newdb()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)
	s := router.NewRouter(userController)
	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", s))
}
