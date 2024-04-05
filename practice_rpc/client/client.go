package main

import (
	"fmt"
	"net/rpc"
)

// main函数是程序的入口
func main() {
	// 建立请求
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("建立连接失败:", err)
		return
	}

	// 初始化用于接收返回值的变量
	var reply string
	// 调用远程RPC函数，向服务端发送请求
	// 第一个参数是服务名称和函数名的组合，格式为 "服务名.函数名"
	// 第二个参数是传递给远程函数的参数，这里是字符串 "evan"
	// 第三个参数是用于接收远程函数的返回值的指针
	err = client.Call("HelloService.Hello", "evan", &reply)
	if err != nil {
		fmt.Println("调用失败:", err)
		return
	}

	// 输出远程函数的返回值
	fmt.Println(reply)
}
