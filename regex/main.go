package main

import (
	"fmt"
	regexp "regexp"
)

const text = "Myemailisgolang@org.com  email1 is python org@.com email2 is java org@.com"

func main() {
	// 传入的字符串中可以写入正规规则 .* .+等等 转移需要\\ 如果使用``就不会发生任何的转义
	// 同样可以使用[] [a-zA-Z0-9]
	re := regexp.MustCompile(`(.*)@([a-zA-Z0-9]+)\.([a-zA-Z0-9]+)`)
	match := re.FindString(text) //在text中找到符合正则规则的字串
	// FindString 只能获得第一个匹配、若想获得多个则需要使用FindAllString
	//第一个为目标str 第二个为要匹配的数量、若全部匹配则使用-1
	re.FindAllString(text,-1)
	//进行子匹配，将需要的内容用括号包裹起来 第一个参数目标str 第二要匹配的数量 返回的是一个二维数组
	submatch := re.FindAllStringSubmatch(text, -1)
	fmt.Println(match)
	for _, val := range submatch {
		// 输出sclie 第零个为完整的地址 第一个为第一个括号的内容一次类推
		fmt.Println(val)
	}


}
