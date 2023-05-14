package main

import (
	handlers "hexgo/internal/adapters/handlers/user"
	repositories "hexgo/internal/adapters/repositories/user"
	"hexgo/internal/core/ports"
	"hexgo/internal/core/usecases"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	var userRepo ports.UserRepository = repositories.NewUserMysqlRepository()
	userUseCase := usecases.NewUserUseCase(userRepo)
	handlers.NewUserHandler(userUseCase, router)

	log.Fatal(http.ListenAndServe(":8080", router))
}
