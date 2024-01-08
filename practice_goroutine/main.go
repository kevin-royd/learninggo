package main

/**
控制并发
每一个并发执行的单元都叫做goroutine，类似其他语言中并发执行的最小单位——线程
当一个程序启动时，其主函数即在一个单独的goroutine中运行，我们叫它main goroutine。
新的goroutine会用go语句来创建。在语法上，go语句是一个普通的函数或方法调用前加上关键字go

goroutine退出只能由本身控制，不允许从外部强制结束该goroutine。
只有两种情况例外，那就是main函数结束或者程序崩溃结束运行；
所以，要实现主进程控制子goroutine的开始和结束，必须借助其它工具来实现。
*/

// 通过计数器来控制并发的数量 并发方式前添加wg.add()  之后添加wg.wait() 逻辑处理完成中添加wg.Done
//var wg sync.WaitGroup
//
//func main() {
//	strSlice := [...]string{
//		"张三", "李四", "王五", "赵一", "孙李",
//	}
//
//	wg.Add(len(strSlice))
//	go func() {
//		for _, v := range strSlice {
//			fmt.Printf("输出%v\n", v)
//			wg.Done()
//		}
//	}()
//	wg.Wait()
//}

//并发时钟测试
//func main(){
//	listen, err := net.Listen("tcp", "localhost:8080")
//	if err != nil{
//		log.Fatal(err)
//	}
//	for{
//		conn, err := listen.Accept()
//		if err != nil{
//			log.Fatal(err)
//			continue
//		}
//		handleConn(conn)
//	}
//}
//
//
//func handleConn(c net.Conn){
//	// 延迟执行
//	defer c.Close()
//	for{
//		_, err := io.WriteString(c, time.Now().Format("19:00:05\n"))
//		if err != nil{
//			return
//		}
//		time.Sleep(1 * time.Second)
//	}
//}

import (
	"context"
	"fmt"
	"time"
)

// tracker 结构体
type Tracker struct {
	ch   chan string   // 数据通道
	stop chan struct{} // 停止信号通道
}

// 初始化 Tracker 的函数
func newTracker() *Tracker {
	return &Tracker{
		ch:   make(chan string, 10), // 带有长度为 10 的缓冲区的数据通道
		stop: make(chan struct{}),
	}
}

func (t *Tracker) Event(ctx context.Context, data string) error {
	select {
	case t.ch <- data:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// Run 方法在单独的 goroutine 中处理数据通道的数据
func (t *Tracker) Run() {

	for {
		select {
		case data := <-t.ch:
			// 模拟数据处理
			time.Sleep(1 * time.Second)
			fmt.Printf("data:%s\n", data)
		case <-t.stop:
			return
		}
	}
}

func (t *Tracker) Shutdown(ctx context.Context) {
	close(t.ch)
	select {
	case <-t.stop:
	case <-ctx.Done():

	}
}

func main() {
	// 初始化 tracker
	tr := newTracker()

	// 启动处理数据的 goroutine
	go tr.Run()

	_ = tr.Event(context.Background(), "test")

	// 等待一段时间
	time.Sleep(3 * time.Second)
	//创建了一个带有截止时间的上下文对象，然后在 defer cancel() 中延迟调用 cancel 函数。
	//这样做的目的是，如果 main 函数在执行过程中出现了错误或提前结束，defer cancel() 会确保及时地取消 ctx，防止潜在的资源泄漏。
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	defer cancel()
	tr.Shutdown(ctx)
}
