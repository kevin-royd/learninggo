package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
数据竞态（Data Race）是一种并发编程中常见的问题，它发生在多个goroutine并发访问共享数据且至少其中一个是写操作的情况下。
当两个或更多的goroutine同时访问同一片内存区域，其中至少有一个是写操作，而且没有使用同步机制（如互斥锁、信号量等）来保护共享数据时，就可能发生数据竞态。
*/

var Wait sync.WaitGroup
var counter int = 0
var v atomic.Value

type Config struct {
	a []int
}

func main() {
	for routine := 0; routine < 2; routine++ {
		Wait.Add(1)
		go work()
	}
	Wait.Wait()

	// 由于 counter 存储在 Config 中，可以在 main 中通过 atomic.Value 安全地读取
	cfg := v.Load().(*Config)
	fmt.Printf("counter = %d\n", cfg.a[0])
}

func work() {
	defer Wait.Done()

	for i := 0; i < 2; i++ {
		counter++
		// 创建新的 Config 对象并存储到 atomic.Value
		cfg := &Config{a: []int{counter}}
		// 存在数据 该操作是原子的不会发生  date race
		v.Store(cfg)
	}
}
