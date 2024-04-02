package main

import (
	"fmt"
	"time"
)

/*
channel 默认是双向的。在使用时期望他是是单行的
案例：单向读 var read <- chan int
	单向写 var write chan <- int
*/

func main() {
	//在初始化时创建在双向channel但在传递时会自动进行转换(需要在接收的位置进行定义)
	//
	c := make(chan int, 3)
	//var send <-chan int = c // 将双向的channel改为单向的
	//var read chan<- int = c
	go producer(c)
	go consumer(c)
	time.Sleep(10 * time.Second)
}

// 只读案例
func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i * i
	}
	close(out)
}

// 只写案例
func consumer(in <-chan int) {
	for num := range in {
		fmt.Printf("num is %d\r\n", num)

	}
}
