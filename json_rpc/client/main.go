package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	// 发起请求
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("建立连接失败")
	}
	var reply string
	// 序列化协议变为json
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	err = client.Call("HelloService.Hello", "evan", &reply)
	if err != nil {
		fmt.Println("发送请求失败")
	}
	fmt.Println(reply)
}
