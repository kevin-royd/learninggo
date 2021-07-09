package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		函数
		格式 func 函数名称(参数名称1 类型1,...,参数名称n 类型n)(返回值名称 返回值类型1,...返回值名称n 返回值类型n)
	*/
	variables, i := actionVariables(2, 1)
	fmt.Println(variables, i)
	actionVariables1(1, "string", 3.14, true)
	var args int
	actionVariables2(args)
	fmt.Println(args)
	// 定义匿名函数，定义时并赋值 注意返回值只有一个时不用写括号
	actionVariables3 := func(var3, var4 int) int {
		return var3 + var4
	}(1, 2)
	fmt.Println(actionVariables3)

	// 定义一个变量、将匿名函数赋值给该变量，则该变量具有匿名函数的功能
	actionVariables4 := func(var1, var2 int) int {
		return var1 - var2
	}
	//将一个匿名函数赋值给一个变量，这个变量存储的是这个匿名的地址
	fmt.Printf("func1调用的值=%d,func1=%p\n", actionVariables4(1, 2), &actionVariables4)

	actionVariables6 := actionVariables5()
	fmt.Println(actionVariables6(1))
	fmt.Println(actionVariables6(2))
	fmt.Println(actionVariables6(3))

	i1, i2 := funcAsArgs(actionVariables7, 1, 2)
	fmt.Println(i1, i2)
}

//	如果参数类型一直则可以忽略在最后一个参数名称 添加参数类型
func actionVariables(var1, var2 int) (int, int) {
	return var1 + var2, var1 - var2
}

// 可变参数
func actionVariables1(args ...interface{}) {
	for i, arg := range args {
		fmt.Println(i, arg)
	}
}

// 局部变量
func actionVariables2(args int) {
	args = 100
	fmt.Println(args)

}

// 闭包：函数体内部引用了函数体外的数据
func actionVariables5() func(int) int {
	step := 0
	return func(_step int) int {
		step += _step
		return step
	}
}

// defer 关键字 允许我们将业务逻辑延迟到函数返回之前才执行某个语句或函数 常用于资源的回收
func deferAction() {
	handle, err := os.Open("hello.txt")
	//	若这个过程中代码出现了问题，资源无法得到释放就会一直占用，所以需要使用defer
	defer handle.Close()
	err = err
	//	defer 同时遵循先进先出原则
}

// 函数做为类型在函数中使用
//				类型名称		传入函数的类型	返回的类型
func funcAsArgs(funcName func(int, int) (int, int), var9, var10 int) (int, int) {
	return funcName(var9, var10)
}

func actionVariables7(var9, var10 int) (sum int, minus int) {
	sum = var9 + var10  //3
	minus = var9 - var10	//-1
	//return sum,minus
	return
}
