package rpc

import "go.uber.org/fx"

var Module = fx.Module(
	"rpc_services",
	fx.Provide(NewUserService),
	fx.Provide(NewServices),
)

type Service interface {
	Setup()
}

type Services []Service

func NewServices(userService UserService) Services {
	return Services{
		userService,
	}
}

func (s Services) Setup() {
	for _, service := range s {
		service.Setup()
	}
}
