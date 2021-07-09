package main

import (
	"io"
	"log"
	"net"
	"time"
)

/**
控制并发
每一个并发执行的单元都叫做goroutine，类似其他语言中并发执行的最小单位——线程
当一个程序启动时，其主函数即在一个单独的goroutine中运行，我们叫它main goroutine。
新的goroutine会用go语句来创建。在语法上，go语句是一个普通的函数或方法调用前加上关键字go

goroutine退出只能由本身控制，不允许从外部强制结束该goroutine。
只有两种情况例外，那就是main函数结束或者程序崩溃结束运行；
所以，要实现主进程控制子goroutine的开始和结束，必须借助其它工具来实现。
*/

//并发时钟测试
func main(){
	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil{
		log.Fatal(err)
	}
	for{
		conn, err := listen.Accept()
		if err != nil{
			log.Fatal(err)
			continue
		}
		handleConn(conn)
	}
}

func handleConn(c net.Conn){
	// 延迟执行
	defer c.Close()
	for{
		_, err := io.WriteString(c, time.Now().Format("19:00:05\n"))
		if err != nil{
			return
		}
		time.Sleep(1 * time.Second)
	}
}
