package main

import (
	"fmt"
	"net"
	"net/rpc"
)

// 定义服务对象
type helloService struct{}

/*
Hello 生产rpc的方法供client调用
request: 客户端传递的参数，
reply: 指向客户端传递的参数的指针
*/
func (h *helloService) Hello(request string, reply *string) error {
	//模拟对数据进行处理后返回
	*reply = "hello " + request
	return nil
}

func main() {
	// 实例化server，监听端口
	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("监控端口失败:", err)
		return
	}

	// 注册一个RPC服务对象
	err = rpc.RegisterName("HelloService", &helloService{})
	if err != nil {
		fmt.Println("注册对象失败:", err)
		return
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
