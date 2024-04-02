package main

import (
	"fmt"
	"time"
)

/*
通过select完成对所有channel的监控哦、每一个channel返回了都应知道
1、某一个分支就绪了就执行该分支
2、如果都就绪了则随机执行。防止饥饿。（参考锁。防止某个线程释放锁后又重新获取锁。导致一直不释放）
*/

// 定义一个空结构体channel 作用和定义值的类型为bool作用类似

func main() {
	// 创建并初始化多个channel
	g1Channel := make(chan struct{}, 1)
	g2Channel := make(chan struct{}, 1)

	go g1(g1Channel)
	go g2(g2Channel)

	//通过timer进行超时控制
	timer := time.NewTimer(3 * time.Second)
	select {
	case <-g1Channel:
		fmt.Println("g1 channel done")
	case <-g2Channel:
		fmt.Println("g2 channel done")
	case <-timer.C:
		fmt.Println("is time out")
	}

}

func g1(ch chan struct{}) {
	time.Sleep(4 * time.Second)
	ch <- struct{}{}

}

func g2(ch chan struct{}) {
	time.Sleep(4 * time.Second)
	ch <- struct{}{}
}
