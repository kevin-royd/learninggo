package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//	浮点数只有float32和64 进行计算时同样需要进行类型转换，用法和int一直
	// 默认float64
	var floatVariables1 = 3.1415926
	var floatVariables2 = .1415926        //自动推导
	var floatVariables3 float32 = 1415926 //自动推导

	fmt.Printf("floatVariables1=%T,floatVariables1=%d\n", floatVariables1, unsafe.Sizeof(floatVariables1))
	fmt.Printf("floatVariables2=%T,floatVariables2=%d\n", floatVariables2, unsafe.Sizeof(floatVariables2))
	fmt.Printf("floatVariables3=%T,floatVariables3=%d\n", floatVariables3, unsafe.Sizeof(floatVariables3))

	//	科学计数法表示
	var floatVariables4 float32 = 3.1415926e2 //值乘以10的2次方 -e变为除
	fmt.Println(floatVariables4)

	//	复数  实数"+"虚数i组成
	//第一种
	var complexVariable1 = 3.14 + 12i
	//	第二种
	complexVariable2 := complex(3.14, 12)
	fmt.Printf("complexVariable1=%T,complexVariable1=%v\n", complexVariable1, complexVariable1)
	fmt.Printf("complexVariable2=%T,complexVariable2=%v\n", complexVariable2, complexVariable2)

//	打印复数的实数和虚数部分
	fmt.Println(real(complexVariable1), imag(complexVariable1))

}
