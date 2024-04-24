package client_proxy

import (
	"github.com/username/learninggo/rpc/practice_rpc/handler"
	"net/rpc"
)

type HelloServiceStub struct {
	*rpc.Client
}

// NewHelloServiceClient 实例化 传入 传输协议和address
func NewHelloServiceClient(protocol string, address string) HelloServiceStub {
	client, err := rpc.Dial(protocol, address)
	if err != nil {
		panic("connection err")

	}
	return HelloServiceStub{client}
}

// Hello 封装方法
func (c *HelloServiceStub) Hello(request string, reply *string) error {
	err := c.Call(handler.HelloServiceName+".Hello", request, reply)
	if err != nil {
		return err
	}
	return nil
}
