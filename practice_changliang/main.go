package main


func main() {
	//	赋值也和变量相同 但关键词为const
	//注：常量在编译器就明确的值不可进行更改
	const constVariables1 float64 = 3.1415926
	const constVariables2, constVariables3 = 3, "张山"
	println(constVariables1, constVariables2, constVariables3)
	//	预定的常量 true/false iota iota比较特殊，认为可以修改的值
	//iota使用：遇到const,iota的值被重置为0，在遇到下一个const之前每次都会+1
	const (
		iotaVariables1 = iota
		iotaVariables2 = iota
		iotaVariables3 = iota
	)

	const iotaVariables4 = iota
	println(iotaVariables1, iotaVariables2, iotaVariables3, iotaVariables4)

	const (
		// const 第二种使用方法,不进行赋值则默认继续使用iota
		iotaVariables5 = iota
		iotaVariables6
		iotaVariables7
	)
	println(iotaVariables5, iotaVariables6, iotaVariables7)

	const (
		// const中使用iota 如果常量有确定的值则iota在当前跳过，但值也进行累加
		iotaVariables8  = iota   //0
		iotaVariables9  = "Bobo" //Bobo
		iotaVariables10 = iota   //2
	)
	println(iotaVariables8, iotaVariables9, iotaVariables10)

}
