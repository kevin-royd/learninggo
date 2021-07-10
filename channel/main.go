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

func work(i int, c chan int) {
	//for {
	//	// 当发送方存在close时 需要用2个参数进行接收chan 第一个值为正常的值、第二个为是否还有值
	//	n, ok := <-c
	//	if !ok {
	//		break
	//	}
	//	fmt.Printf("work %d recevied %c\n", i, n)
	//}
	//	简便写法 发送方没有close一样可以接收
	for n := range c {
		fmt.Printf("work %d recevied %c\n", i, n)
	}
}

func createdWorker(i int) chan<- int {
	c := make(chan int)
	go work(i, c)
	return c
}

func bufferedChannel() {
	c := make(chan int, 3) //创建并初始化chan
	//	channel创建好了过后需要进行发数据 没人来接受channel就会死掉 就会死锁
	go work(0, c)
	c <- 'a' //发送了数据之后了就一定会进行协程的切换 比较耗费资源
	c <- 'b'
	c <- 'c'
	// 设置接收时间 1毫秒
	time.Sleep(time.Millisecond) //需要加线程随便 因为main也是goroutine 长时间未接受到数据main就关闭了
}

/*
通知接收方 数据已发送完成 没有新的数据要发送
*/
func channelClose() {
	//	注意发送方close 但接收方还是能接收到数据 是否有缓存都不重要
	c := make(chan int)
	go work(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	close(c)
	time.Sleep(time.Millisecond)

}

func main() {
	chanDemo()
	//bufferedChannel()
	//	channelClose chan有明确的结尾
	//channelClose()
}
