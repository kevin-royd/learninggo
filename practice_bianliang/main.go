package main

import "fmt"

// 包名需要改成 main 项目以包为单位、每一个文件都隶属于一个包

//申请全局变量
var (
	variablesx int
	//	切片
	slicex []int
	//	结构体
	interfacex interface{}
)

func main() {
	// 定义变量注意事项
	//大写的变量名是能被其他包读取的 是共有变量 函数同理
	//	 第一种请求变量用 var
	// var +变量名 +数据类型 不赋值就为默认值
	var variables1 int
	// 第二种使用变量 变量名称右侧使用:= 右侧为变量值 要求改变量没有被声明过
	// 简短声明 不能在函数体外使用
	variables2 := "张山"
	// 特殊符号_ 特殊变量名 任何赋值给他值的都会被丢弃
	_, x := 22, 222
	println("x:", x)
	// 注意声明变量必须使用
	// 声明多个变量 方式：
	var variables3, variables4, variables5 = 1, "dd", true
	// 两个数据进行交换
	var i, j = 100, 200
	i, j = j, i
	println(i, j)
	//注意 任何声明的变量名未使用都会报错
	fmt.Println(variables1, variables2, variables3, variables4, variables5)
	changLiang()
}

// 变量格式化输出
// 1.	%v 值的默认格式表示
//	2.	%+v 站类似%v，但输出的结构体时会添加字段名
//	3.	%#v 值的GO语言表示
//	4.	%T 值的类型的GO语言表示
//	5.	%% 百分号
//	6.	%t 单纯true或false 布尔值

func changLiang() {
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
