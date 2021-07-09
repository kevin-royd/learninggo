package main

import (
	"fmt"
	"time"
)

//定义一个channel chan后面为类型
func chanDemo() {
	//var c chan int // 此时的c 为nil
	//	在函数中使用使用make进行创建使用
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		//将返回的channel进行储存
		channels[i] = createdWorker(i)
	}
	for i := 0; i < 10; i++ {
		//	对channels进行分发数据
		channels[i] <- 'a' + i
	}
	time.Sleep(time.Millisecond)

}

func createdWorker(i int) chan<- int {
	c := make(chan int)
	go func() {
		for {
			fmt.Printf("work %d recevied %c\n", i, <-c)
		}
	}()
	return c
}

func main() {
	chanDemo()
}
