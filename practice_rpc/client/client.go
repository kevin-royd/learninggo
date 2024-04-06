package main

import (
	"fmt"
	"github.com/username/learninggo/practice_rpc/client_proxy"
)

// main函数是程序的入口
func main() {
	// 建立请求
	client := client_proxy.NewHelloServiceClient("tcp", "localhost:1234")
	// 初始化用于接收返回值的变量
	var reply string
	err := client.Hello("evan", &reply)
	if err != nil {
		fmt.Println("调用失败:", err)
		return
	}
	// 输出远程函数的返回值
	fmt.Println(reply)
}
