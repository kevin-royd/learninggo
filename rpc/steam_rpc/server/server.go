package main

import (
	"fmt"
	"github.com/username/learninggo/rpc/steam_rpc/proto"
	"google.golang.org/grpc"
	"net"
	"sync"
	"time"
)

// 定义对象
type Server struct {
}

// 服务流模式。服务端一直发送数据。基于sock
func (s *Server) GetSteam(req *proto.StreamReqData, rsp proto.Greeter_GetSteamServer) error {
	i := 0
	for {
		i++
		if i > 10 {
			break
		}
		// 写入数据
		err := rsp.Send(&proto.StreamRspData{Data: fmt.Sprintf("%v", time.Now().Unix())})
		if err != nil {
			return err
		}
		time.Sleep(time.Second)
	}
	return nil
}

// 客户端流模式和服务器流模式代码反过来即可
func (s *Server) PutSteam(cliStr proto.Greeter_PutSteamServer) error {
	return nil
}

// 双向流模式可能同时写和同时发 所以需要携程并发处理
func (s *Server) AllSteam(allStr proto.Greeter_AllSteamServer) error {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			if data, err := allStr.Recv(); err != nil {
				fmt.Println(err)
				break
			} else {
				fmt.Println(data.Data)
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
			if err := allStr.Send(&proto.StreamRspData{Data: fmt.Sprintf("%v", time.Now().Unix())}); err != nil {
				fmt.Println(err)
				break
			}
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
	return nil
}

func main() {
	// 监听端口
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err.Error())
	}
	//初始化对象
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &Server{})
	err = s.Serve(listener)
	if err != nil {
		panic(err.Error())
	}
}
