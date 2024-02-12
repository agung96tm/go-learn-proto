package controllers

import "net/http"

type UserController struct {
}

func NewUserController() UserController {
	return UserController{}
}

func (c UserController) Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("INDEX"))
}

func (c UserController) Detail(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("INDEX"))
}
