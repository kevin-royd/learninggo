package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
案例通过goroutine的方式监控cpu并且可以主动推出
使用context的方式
WithCancel 自动发起取消通知
WithTimeout 在使用时候直接传递超时时间 context.WithTimeout(context.Background(),6*time.Second)
WithDeadline 时间点cancel
WithValue 可以在链路中进行值的传递 ctx := context.WithValue(context.Background(),"key123","value123")
*/

// 使用锁保证线程安全
var wg sync.WaitGroup

func main() {

	wg.Add(1)
	// 如果你的goroutine，函数中，如果希望被控制、超时、传值，但我不希望影响原来的结果时，参数的第一个参数尽量写context
	ctx, cancelFunc := context.WithCancel(context.Background())
	// 生成的ctx还可以继续WithCancel继续派生。结构类似树状型
	go cpuInfo(ctx)
	// 模拟主动推出
	time.Sleep(6 * time.Second)
	// 主动发出取消信息。 返回的取消函数的作用是通知相关的操作停止执行。
	cancelFunc()
	wg.Done()
	fmt.Println("监控完成")
}

func cpuInfo(ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			//主动推出监控
			fmt.Println("退出cpu监控")
			return
		default:
			time.Sleep(2 * time.Second)
			fmt.Println("监控cpu")
		}
	}
}
