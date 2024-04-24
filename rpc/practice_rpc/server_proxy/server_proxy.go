package server_proxy

import (
	"fmt"
	"github.com/username/learninggo/rpc/practice_rpc/handler"
	"net/rpc"
)

// 接收类型需要结偶 所以定义为接口类型
type HelloServiceInterface interface {
	Hello(request string, reply *string) error
}

func RegisterName(srv HelloServiceInterface) error {
	// 注册一个RPC服务对象
	err := rpc.RegisterName(handler.HelloServiceName, srv)
	if err != nil {
		return fmt.Errorf("注册服务失败")
	}
	return nil
}
