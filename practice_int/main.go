package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//	定义了一个变量 并赋予他值，若没有给他类型 go也能推导他的类型 uint是int的一倍，uint没有负数
	var intVariables1 = 100 //整型为int
	intVariables2 := 200
	// 注意 直接println方法为输出，和fmt.println相同，但fmt.print有多个输出类型
	fmt.Printf("intVariables1=%T,intVariables2=%T\n", intVariables1, intVariables2)
	//	注：int64=int int16 int32 int64所占的空间大小不同，范围区间也不同，依次2的次方-1
	//	注2：使用int8以上时需要带上变量特殊字 var
	var intVariables3 int64 = 123456
	//同理需要注意大转小时数据精度的问题
	var intVariables4 = int32(intVariables2)
	var intVariables5 int8 = 100
	fmt.Println(unsafe.Sizeof(intVariables1), unsafe.Sizeof(intVariables3), unsafe.Sizeof(intVariables4), unsafe.Sizeof(intVariables5))

}
