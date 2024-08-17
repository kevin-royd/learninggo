package main

import (
	"github.com/BurntSushi/toml"
	"log"
	"os"
	"path/filepath"
)

func main() {
	conf := Instance()
	// 尽量减少使用反射。因为反射内存消耗较大
	name := conf.test.Name
	log.Printf("从toml中配置test域下name=%v\n", name)

}

// 1、toml对象结构体
type config struct {
	test Test
}

type Test struct {
	Name string `toml:"name"`
}

// Instance 初始化配置文件对象
func Instance() *config {
	var conf *config
	// 获取程序执行的根目录
	dir, _ := os.Getwd()
	//拼接配置文件目录
	filePath := filepath.Join(dir, "/toml/conf/config.toml")
	if _, err := toml.DecodeFile(filePath, &conf); err != nil {
		panic(err)
	}
	return conf
}
