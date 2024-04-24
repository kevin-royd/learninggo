package main

import (
	"context"
	proto2 "github.com/username/learninggo/rpc/proto_rpc/proto"
	"google.golang.org/grpc"
	"net"
)

// 初始化对象 实现接收类型的方法
type Server struct {
}

func (s *Server) SayHello(ctx context.Context, request *proto2.HelloRequest) (*proto2.HelloReply, error) {
	return &proto2.HelloReply{Message: "hello " + request.Name}, nil
}

func main() {
	// 实例化对象
	g := grpc.NewServer()
	// 注册实例
	proto2.RegisterGreeterServer(g, &Server{})
	listener, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		panic("监听端口失败:" + err.Error())
	}
	err = g.Serve(listener)
	if err != nil {
		panic("启动服务失败:" + err.Error())

	}
}
