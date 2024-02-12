package lib

import (
	"google.golang.org/grpc"
	"net"
)

type RpcHandler struct {
	Engine *grpc.Server
}

func NewRpcHandler() RpcHandler {
	return RpcHandler{
		Engine: grpc.NewServer(),
	}
}

func (l *RpcHandler) Run() error {
	lis, err := net.Listen("tcp", ":4000")
	if err != nil {
		panic(err)
	}
	return l.Engine.Serve(lis)
}
