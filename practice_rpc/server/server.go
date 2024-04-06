package main

import (
	"fmt"
	"github.com/username/learninggo/practice_rpc/handler"
	"github.com/username/learninggo/practice_rpc/server_proxy"
	"net"
	"net/rpc"
)

func main() {
	// 实例化server，监听端口
	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("监控端口失败:", err)
		return
	}

	err = server_proxy.RegisterName(&handler.HelloService{})
	if err != nil {
		panic(err)
	}

	// 循环监听客户端连接，并为每个连接提供RPC服务
	for {
		// 接收新的客户端连接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("获取sock失败:", err)
			continue
		}
		// 为客户端连接提供RPC服务 如果没有指明序列化的传输协议。go的序列化和反序列化都是使用的默认gob格式。客户端也是一样
		// 为了处理并发。需要通过携程的方式进行处理
		go rpc.ServeConn(conn)
	}
}
