package main

import "fmt"

func main() {
	/*
		字符 字节
		byte(uint8)
		rune unicode int32 汉字在unicode中占3个字节
		流程 字符--> ascii编码--> 二进制--> ascii编码--> 字符
	*/

	/*
		字符串 string
	*/
	var stringVariables1 string = "hello world,张山"
	// ``这里调用print stringVariables2 会将赋值 原格式输出
	var stringVariables2 = `sdfsadfsadfasdfasfd`
	fmt.Printf("stringVariables1=%T,stringVariables2=%T\n", stringVariables1, stringVariables2)
	fmt.Printf("stringVariables1=%v,stringVariables2=%v", stringVariables1, stringVariables2)

	//	for循环便利和其他语法一样 底层使用的Unicode编码 遍历字符串出现中文就会出现问题
	for i := 0; i < len(stringVariables1); i++ {
		fmt.Printf("	编码值=%d,值=%c\n", stringVariables1[i], stringVariables1[i])
	}

	for _, value := range stringVariables1 {
		//字符串处理时含有中文 则数据类型为Int32
		fmt.Printf("%s编码值=%d,值=%c,类型=%T\n", stringVariables1, value, value, value)

	}

}
