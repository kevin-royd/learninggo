package main

import (
	"fmt"
	"os"
)

func main() {
	/**
	流程控制
	*/
	//	if判断多个条件进行赋值，
	if fileHeader, err := os.Open("./hello.text"); err != nil {
		fmt.Println(fileHeader)
		fmt.Println(err.Error())
	} else {
		fmt.Println(fileHeader)
	}
	//	 switch case select 判断时直接选择
	switch switchVariables := 100; {
	case switchVariables > 90:
		fmt.Println("成绩优秀")
		// fallthrough 的作用是允许贯穿继续执行
		//fallthrough
	case switchVariables > 80:
		fmt.Println("成绩良好")
	}

	//	switch if else
	switchVariables1 := 100
	switch {
	//优先匹配，若都瞒住条件则匹配第一个
	case switchVariables1 > 90:
		fmt.Println("成绩优秀")
	case switchVariables1 > 80:
		fmt.Println("成绩良好")
	case switchVariables1 >= 60:
		fmt.Println("成绩合格")
	default:
		fmt.Println("成绩不及格")
	}
	
//	类型断言
	switchVariables2 := "笔记本"
	switch switchVariables2 {
	case "笔记本","ipad","台式":
		fmt.Println("数码产品")
	case "饼干", "巧克力", "苏打":
		fmt.Println("零食产品")
	}
//	跳转控制	goto
	//跳转到某个标签
	//break退出当前循环
	//continue 忽略当前循环剩余的代码
}
