package routes

import (
	"github.com/agung96tm/go-learn-proto/apps/api/controllers"
	"github.com/agung96tm/go-learn-proto/lib"
	"net/http"
)

type UserRouter struct {
	handler    lib.HttpHandler
	controller controllers.UserController
}

func NewUserRouter(handler lib.HttpHandler, controller controllers.UserController) UserRouter {
	return UserRouter{
		handler:    handler,
		controller: controller,
	}
}

func (r UserRouter) Setup() {
	r.handler.Engine.HandlerFunc(http.MethodGet, "/users", r.controller.Index)
	r.handler.Engine.HandlerFunc(http.MethodGet, "/users/:id", r.controller.Detail)
}
