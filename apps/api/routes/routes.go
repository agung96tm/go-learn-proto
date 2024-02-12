package routes

import "go.uber.org/fx"

var Module = fx.Module(
	"routes",
	fx.Provide(NewUserRouter),
	fx.Provide(NewRoutes),
)

type Router interface {
	Setup()
}

type Routes []Router

func NewRoutes(userRouter UserRouter) Routes {
	return Routes{
		userRouter,
	}
}

func (r Routes) Setup() {
	for _, router := range r {
		router.Setup()
	}
}
