package main

import "fmt"

func main() {
	/**
	指针:存储的是另一个变量的内存地址
	获取变量在内存中的地址 & + 变量名称
	*/
	intVariables := 100
	fmt.Printf("intVariables=%d,地址=%v\n", intVariables, &intVariables)
	//	定义指针变量 注：定义指针时不能使用缩减写法
	var pointerVariable *int = &intVariables
	fmt.Printf("pointerVariable=%v,地址=%v\n", pointerVariable, &pointerVariable)
	//	注：定义什么类型的变量就会在内存中生成类型的指针、不同类型的指针进行访问时会报错的
	//	指针的作用：
	//	1. 节省内存空间、提高程序的执行效率
	//	2. 间接访问与修改变量的值
	//	注：
	//	1.操作指针都不能使用：=
	//	2.不能把变量的值直接赋值为指针类型,指针
	//	3.指针可以指向另一个指针，不推荐
	*pointerVariable = 200
	fmt.Println(intVariables)

	/*
	值类型：整型，浮点型，bool，array，string， 栈中分配
	引用类型：指针，slice，map，chan，interface，堆中分配
	 */
}
