package main

import (
	"context"
	"fmt"
	"strconv"
	"sync"
)

/*
	1、及时取消请求
	ctx, cancel := context.WithCancel(context.Background())
    defer cancel() // 在完成工作或函数退出时取消上下文
	2、WithValue对值的传递
	3、// 处理上下文取消的清理工作
	select {
    case <-ctx.Done():
    }
	4、不要在上下文传递nil
		ctx := context.Background()
	5、在函数中传递上下文时。确保传递的是派生出的新上下文，而不是直接传递原始的上下文。这可以避免在不同函数间共享相同上下文引起的不确定性
*/

type KeyType string

const (
	myKey KeyType = "myKey"
)

var (
	Wait sync.WaitGroup
)

type Config struct {
	value sync.Map
}

func main() {

	// 创建一个根上下文
	rootContext := context.Background()
	// 设置初始化key,value
	cfg := &Config{}
	cfg.value.Store("key0", "value0")
	// 对值进行修改
	ctx := context.WithValue(rootContext, myKey, cfg)
	for i := 1; i < 3; i++ {
		Wait.Add(1)
		go readValue(ctx, i)
	}
	Wait.Wait()
	// 直接使用 cfg，不再需要 atomic.Value
	fmt.Printf("最终值为: %v\n", cfg.value)
}

func readValue(ctx context.Context, data int) {
	defer Wait.Done()

	// // 从上下文中获取原子配置
	// 获取当前原子值
	value := ctx.Value(myKey).(*Config)
	strData := strconv.Itoa(data)
	// 新建一个键值对
	newKey := "key" + strData
	newValue := "value" + strData
	// 在这个例子中，我们简单地打印值
	value.value.LoadOrStore(newKey, newValue)

}
