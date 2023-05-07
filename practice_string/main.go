package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	/*
		字符 字节
		byte(uint8)
		rune unicode int32 汉字在unicode中占3个字节
		流程 字符--> ascii编码--> 二进制--> ascii编码--> 字符
		如果str.len返回的是字节数。查看字符的长度需要使用
	*/

	/*
		字符串 string
	*/
	var stringVariables1 string = "hello world,张山"
	// ``这里调用print stringVariables2 会将赋值 原格式输出
	var stringVariables2 = `sdfsadfsadfasdfasfd`
	fmt.Printf("stringVariables1=%T,stringVariables2=%T\n", stringVariables1, stringVariables2)
	fmt.Printf("stringVariables1=%v,stringVariables2=%v\n", stringVariables1, stringVariables2)
	//输出字符的长度
	fmt.Printf("字符长度为:%d\n", utf8.RuneCountInString(stringVariables1))
	//	for循环便利和其他语法一样 底层使用的Unicode编码 遍历字符串出现中文就会出现问题
	for i := 0; i < len(stringVariables1); i++ {
		fmt.Printf("	编码值=%d,值=%c\n", stringVariables1[i], stringVariables1[i])
	}

	for _, value := range stringVariables1 {
		//字符串处理时含有中文 则数据类型为Int32
		fmt.Printf("%s编码值=%d,值=%c,类型=%T\n", stringVariables1, value, value, value)

	}

	/*
		字符串拼接官方推荐使用strings.Builder
	*/
	var b strings.Builder
	b.WriteString("拼接字符1")
	b.WriteString("拼接字符2")
	// 最终合成一个string
	b.String()

}
