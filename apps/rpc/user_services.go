package rpc

import (
	"context"
	"github.com/agung96tm/go-learn-proto/lib"
	"github.com/agung96tm/go-learn-proto/resources/protos"
)

type UserService struct {
	protos.UnimplementedUsersServer
	handler lib.RpcHandler
}

func NewUserService(handler lib.RpcHandler) UserService {
	return UserService{
		handler: handler,
	}
}

func (s UserService) ListUser(ctx context.Context, in *protos.ListUserIn) (*protos.ListUserOut, error) {
	return &protos.ListUserOut{
		Users: []*protos.User{
			{Id: 1, Name: "Agung Yuliyanto", Email: "agung.96tm@gmail.com"},
			{Id: 1, Name: "Agung Y", Email: "agung.57tm@gmail.com"},
		},
	}, nil
}

func (s UserService) Setup() {
	protos.RegisterUsersServer(s.handler.Engine, s)
}
