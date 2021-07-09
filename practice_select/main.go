package main

import "fmt"

/*
select和switch作用类似
select 中虽然也有多个 case，但是这些 case 中的表达式必须都是 Channel 的收发操作
*/

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		//<- 管道符 将右侧的值取出赋值给左侧 将x赋值为c
		case c <- x:
			//0,1
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	//第二部
	go func() {
		for i := 0; i < 6; i++ {
			//0,1,2
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	//第一步
	fibonacci(c, quit)
}
