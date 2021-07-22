package main

import (
	"flag"
	"fmt"
	"github.com/urfave/cli"
	"os"
)

/*
命令行参数
*/

//使用flog添加命令行参数
var (
	recusive bool
	test     string
	level    int
)

// 初始化方法会在main函数之前执行
func init() {
	// 第一个为内存地址，第二个为传入名称，第三个为默认值，第四个位解释说明
	// 若在命令行传入了改命名名称 则不使用默认值 使用传入的值
	flag.BoolVar(&recusive, "r", false, "recusive xxx")
	flag.StringVar(&test, "t", "defult string", "string option")
	flag.IntVar(&level, "l", 1, "level of xxx")
	//	设置完成后需要执行 flag.Parse() 添加其中
	flag.Parse()
}

//func main() {
//	// os.Args[0] 为程序名称
//	fmt.Printf("args[0]=%s", os.Args[0])
//	if len(os.Args) > 1 {
//		for i, v := range os.Args {
//			if i == 0 {
//				continue
//			}
//			fmt.Printf("args[%d]=%v", i, v)
//		}
//	}
//}

// 通过cli获取用户传入的参数
func main() {
	//创建一个app
	app := cli.NewApp()
	app.Name = "greet"                 //指定app的名称
	app.Usage = "fight the loneliness" //程序说明
	// 执行命令时要执行的操作
	app.Action = func(c *cli.Context) error {
		var cmd cli.Args
		//判断执行命令的参数是大于0
		if c.NArg() > 0 {
			cmd = c.Args()
		}
		fmt.Println("hello friend cmd:", cmd)
		return nil
	}
	app.Run(os.Args)
}
