package main

import (
	"context"
	"github.com/username/learninggo/rpc/validate_rpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"
)

type server struct {
}

type interceptorInter interface {
	Validate() error
}

func (s *server) SayHello(ctx context.Context, req *proto.Person) (*proto.Person, error) {
	return &proto.Person{
		Id:     req.Id,
		Email:  req.Email,
		Mobile: req.Mobile,
	}, nil
}

func main() {
	// 注册拦截器
	interceptor := func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		if r, ok := req.(interceptorInter); ok {
			if err = r.Validate(); err != nil {
				return nil, status.Error(codes.InvalidArgument, err.Error())
			}
		}
		return handler(ctx, req)
	}
	var opts []grpc.ServerOption
	opts = append(opts, grpc.UnaryInterceptor(interceptor))
	g := grpc.NewServer(opts...)
	proto.RegisterGreeterServer(g, &server{})
	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	g.Serve(lis)
}
