package main

import (
	"fmt"
	"sync"
)

// 通过结构体包装发送chan和通信chan
type woWork struct {
	in chan int
	//wg *sync.WaitGroup //引用 使用指针类型
//	函数式编程完成
	done func()
}

/*
通过通信共享内存
*/
func chanDemo() {
	/*
		wg.add 告知有多少个任务、每个任务做完执行done、wait等待任务做完
	*/
	var wg sync.WaitGroup
	var worker [10]woWork
	//wg.Add(20) //明确是可以直接写入任务数量 未知时放入循环中
	for i := 0; i < 10; i++ {
		worker[i] = createdWorker(i, &wg)
	}
	for i, work := range worker {
		work.in <- 'a' + i
		wg.Add(1)
	}
	for i, work := range worker {
		work.in <- 'A' + i
		wg.Add(1)
	}
	wg.Wait()
	/*
		接收done
		方式一：连续2个发送者若接收者没有并发就会出现死锁
		方法二：将接收done进行分离 进入2个for循环中，并且插入连续发送者之间也不出现死锁
		方法三：使用WaitGroup
	*/
	//for _, work := range worker {
	//	<-work.done
	//	<-work.done
	//}

}

//接收者
func work(i int,w woWork) {
	for n := range w.in {
		fmt.Printf("work %d recevied %c\n", i, n)
		//	告诉发送者数据接收完成
		//go func() {
		//	// 若不加并发 则会出现死锁
		//	done <- true
		//}()
		w.done()
	}
}

// 创建工作线程
func createdWorker(id int, wg *sync.WaitGroup) woWork {
	w := woWork{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go work(id, w)
	return w
}

func main() {
	chanDemo()
}
