package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

/*
闭包	1.必须需要一个函数体 这里的func(v int)就为函数体
	2.必须具有局部变量 这里的v就为局部变量
	3.必须要有自由变量、意思为必有函数体外的参数
*/
func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

/*
正统式函数编程
*/
// 定义类型
type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

func main() {
	//a := adder()
	a1 := adder2(0)
	var s int
	for i := 0; i < 10; i++ {
		s, a1 = a1(i)
		//fmt.Printf("0+ ... + %d=%d\n", i, a(i))
		fmt.Printf("0+ ... + %d=%d\n", i, s)
	}
	f := feiBoNaQi()
	printFileContents(f)
}

/*
将斐波那契数列当成文件来读
*/

// 定义类型 只要是类型就能实现接口 类型：intGen 值就为 func() int
type intGen func() int

func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

/*
斐波那契数列
*/
func feiBoNaQi() intGen {
	j, k := 0, 1
	return func() int {
		j, k = k, j+k
		return j
	}
}
