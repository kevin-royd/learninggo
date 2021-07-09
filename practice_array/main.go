package main

import "fmt"

func main() {
	/*
		go中数据和其他语言区别不大
	*/
	//	 1.数组的定义 同样需要var字符 并且需要声明数组的长度
	var arrayVariable = [5]int{1, 2, 3, 4}
	fmt.Println(arrayVariable)
	// 定义一个数组，不设置长度，同时进行赋值，赋值完后go能推导出数组的长度
	arrayVariables1 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arrayVariables1, len(arrayVariables1))
	//	数据中存在2个值，一个下标，一个索引
	//	指定索引
	arrayVariable2 := [...]int{1: 100, 2: 200, 3: 300}
	fmt.Println(arrayVariable2)
	//	二维数组		 第一个[表示数组的长度],[表示每一个索引位置的数量]
	arrayVariable3 := [4][2]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}
	fmt.Println(arrayVariable3)
	//	遍历数组的方式 同理两种 第一种获取的是数组的索引位置，第二种是可以选择数组下标的值或者索引+下标
	for i := 0; i < len(arrayVariable3); i++ {
		fmt.Println(i)
	}
	for i, n := range arrayVariable3 {
		fmt.Println(i, n)
	}
}

func changArrByPointer(arr *[10]int) {
	(*arr)[0] = 100

}
