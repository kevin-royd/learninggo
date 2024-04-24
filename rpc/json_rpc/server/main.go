package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type helloService struct {
}

func (h *helloService) Hello(request string, reply *string) error {
	*reply = "hello " + request
	return nil
}

func main() {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("监控端口失败")
	}
	// 注册方式
	err = rpc.RegisterName("HelloService", &helloService{})
	if err != nil {
		fmt.Println("注册对象失败")
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("获取连接失败")

		}
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}

}
