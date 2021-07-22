package main

import (
	"encoding/json"
	"fmt"
)

/*
将结构体转换为json对象
*/
func main() {
	c := &Class{
		Name:  "101",
		Count: 100,
	}

	for i := 0; i < 10; i++ {
		s := &Student{
			Id:   i,
			Name: fmt.Sprintf("stu%d", i),
			Sex:  "man",
		}
		//	将每个信息添加到切片中
		c.Student = append(c.Student, s)
	}
	data, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}
	//得到的二进制编码需要转为字符串
	fmt.Printf("json in %s\n", string(data))
	/*
		json反序列化
	*/
	err = json.Unmarshal(data, c)
	if err != nil {
		panic(err)
	}
	// 因为student为指针地址、所以切片中存放的为指针地址
	fmt.Printf("c1:%#v\n", c)
	//	所以得到单个对象需要遍历
	for _, v := range c.Student {
		fmt.Printf("student:%#v\n", v)
	}
}

type Class struct {
	Name    string
	Count   int
	Student []*Student
}

type Student struct {
	Id   int
	Name string
	Sex  string
}
