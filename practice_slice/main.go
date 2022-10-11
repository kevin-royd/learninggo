package main

import "fmt"

func main() {
	/*
		切片底层使用的数组，切片的长度通过len获取，容器通过cap获取
		在函数当中传递切片时，当数据类型数据较大时，使用切片可以有效减少内存占用，提高程序执行效率
	*/
	//	定义切片
	var sliceVariables []int
	//	定义数组
	arrayVariable := [...]int{11, 12, 13, 14, 15}
	for i := 0; i < len(arrayVariable); i++ {
		fmt.Printf("arrayVariable[%d]=%d,地址=%p\n", i, arrayVariable[i], &arrayVariable[i])
	}
	//	若直接定义一个切片去引用数组，那么这个切片默认长度和内存地址一致
	// 这里[:]表示将数组的索引全部赋值给切片
	sliceVariables = arrayVariable[:]
	for i := 0; i < len(sliceVariables); i++ {
		fmt.Printf("sliceVariables[%d]=%d,地址=%p\n", i, sliceVariables[i], &sliceVariables[i])
	}
	// 新的切片使用时必须进行定义
	var sliceVariables1 []int
	// 将数组的部分值赋值给切片时，切片的内存地址还是和数组一致
	sliceVariables1 = arrayVariable[1:3]
	for i := 0; i < len(sliceVariables1); i++ {
		fmt.Printf("sliceVariables1[%d]=%d,地址=%p\n", i, sliceVariables1[i], &sliceVariables1[i])
	}

	//	第二种 通过make方式创建切片，这种方式可以指定切片的类型，大小和容器、未赋值时有默认值，这种方式切片对应的数组不可见，由make来维护
	// 这里创建一个int类型，初始化元素个数为5，容量的长度6
	// 容量的意思为：容器的总容量-左指针走过的索引位置 切片中指针只有左指针，并且索引位置从0开始
	var sliceVariables2 []int = make([]int, 5, 6)
	fmt.Printf("sliceVariables2长度=%d,sliceVariables2容量=%d,sliceVariables2数组地址=%p,sliceVariables2内存地址=%p\n", len(sliceVariables2), cap(sliceVariables2), sliceVariables2, &sliceVariables2)
	//	切片通过append追加元素 当追加的元素超过容量时，go会向内存中新申请一块内存将原始数据拷贝到新内存中，
	//	并指向新的内存空间 意思为新的切片容量增加，底层数组地址改变、但内存地址不变
	sliceVariables2 = append(sliceVariables2, 7, 8, 9)
	fmt.Printf("sliceVariables2长度=%d,sliceVariables2容量=%d,sliceVariables2数组地址=%p,sliceVariables2内存地址=%p\n", len(sliceVariables2), cap(sliceVariables2), sliceVariables2, &sliceVariables2)

	//	 slice的拷贝 源切片sliceVariables3 模板切片sliceVariables4
	sliceVariables3 := []int{1, 2, 3}
	sliceVariables4 := make([]int, 10)
	//将源切片的下标赋值到目标切边中，当源切片的长度小于模板切片时，目标切片剩余下标用默认值填充
	copy(sliceVariables4, sliceVariables3)
	fmt.Println(sliceVariables3)
	fmt.Println(sliceVariables4)

	//	slice做为参数传递，比数组的cup效率更高

	// 如何slice追加的元素为sclie 则需要添加... 表示多个元素
	newSlice := append(sliceVariables3, sliceVariables4[:]...)
	fmt.Println(newSlice)
}
