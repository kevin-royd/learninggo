package main

import (
	"context"
	"fmt"
	"github.com/username/learninggo/proto_rpc/proto"
	"google.golang.org/grpc"
)

func main() {
	// 建立连接 WithInsecure使用不安全的连接
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		panic("连接失败：" + err.Error())
	}
	// 适应返回的是个连接 需要主动关闭
	defer conn.Close()
	// 初始化grpc客户端
	client := proto.NewGreeterClient(conn)
	// 调用接口
	reply, err := client.SayHello(context.Background(), &proto.HelloRequest{Name: "evan"})
	if err != nil {
		panic("调用失败：" + err.Error())

	}
	// 输出返回结果
	fmt.Println(reply)
}
