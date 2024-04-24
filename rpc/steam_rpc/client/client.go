package main

import (
	"context"
	"fmt"
	"github.com/username/learninggo/rpc/steam_rpc/proto"
	"google.golang.org/grpc"
	"sync"
	"time"
)

func main() {
	// 拨号
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()
	client := proto.NewGreeterClient(conn)
	//getSteam, err := client.GetSteam(context.Background(), &proto.StreamReqData{Data: "evan"})
	allSteam, err := client.AllSteam(context.Background())
	if err != nil {
		panic(err.Error())

	}
	wg := sync.WaitGroup{}
	wg.Add(2)
	// 因为是一直接收所以需要循环
	go func() {
		defer wg.Done()
		for {
			if result, err := allSteam.Recv(); err != nil {
				fmt.Println(err)
			} else {
				//fmt.Println(result.Data)
				fmt.Printf("收到服务端消息:%s\n", result.Data)
			}
		}
	}()

	go func() {
		defer wg.Done()
		i := 0
		for {
			i++
			if i > 10 {
				break
			}
			if err := allSteam.Send(&proto.StreamReqData{Data: fmt.Sprintf("客户端第 %d 发送消息", i)}); err != nil {
				fmt.Println(err)
				break
			}
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
}
