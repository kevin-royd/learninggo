package handler

/*
解决注册名称冲突问题
*/

// 注册名称
var HelloServiceName = "HelloService"

// 定义服务对象
type HelloService struct{}

func (h *HelloService) Hello(request string, reply *string) error {
	//模拟对数据进行处理后返回
	*reply = "hello " + request
	return nil
}
