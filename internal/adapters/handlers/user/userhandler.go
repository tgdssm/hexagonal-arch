package handlers

import (
	"hexgo/internal/core/ports"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type userHandler struct {
	userUseCase ports.UserUseCase
}

func NewUserHandler(userUseCase ports.UserUseCase, router *httprouter.Router) {
	handler := &userHandler{
		userUseCase: userUseCase,
	}

	router.POST("/users", handler.Create)
	router.GET("/users", handler.List)
	router.POST("/users/:id", handler.Get)

}

func (uh userHandler) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}

func (uh userHandler) List(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("List users"))

}

func (uh userHandler) Get(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}
