package main

import (
	"context"
	"fmt"
	"github.com/username/learninggo/rpc/validate_rpc/proto"
	"google.golang.org/grpc"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := proto.NewGreeterClient(conn)
	// WithTimeout client请求响应超时控制
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	p := &proto.Person{
		Id:     1000,
		Email:  "abc@gmail.com",
		Mobile: "13888888888",
	}
	resp, err := client.SayHello(ctx, p)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
