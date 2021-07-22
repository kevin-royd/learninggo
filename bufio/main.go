package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
bufio相关
*/
func main() {
	//	Scanf只能读取空格之间的字符串、他以空格做为分隔符
	var strs string
	fmt.Scanf("%s", &strs)
	fmt.Printf("this is from scanf %s\n", strs)

	reader := bufio.NewReader(os.Stdin)
	// 读取什么字符串结束 注意单引号为字节、双引号为字符串
	str, _ := reader.ReadString('\n')
	fmt.Printf("this is from bufio %s\n", str)
}
