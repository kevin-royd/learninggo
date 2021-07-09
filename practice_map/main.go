package main

import "fmt"

func main() {
	/*
		map又称为集合、相当于python中的字典，key-value的形式，并且是无序的
	*/
	//	定义map			map不能改 map[key的类型]value类型
	var mapVariables map[string]string
	//	若定义的时候未进行直接赋值，则需要使用make进行分配数据空间 这里的2什么map的数量
	mapVariables = make(map[string]string, 2)
	mapVariables["monday"] = "周一"
	mapVariables["Tuesday"] = "周二"
	fmt.Println(mapVariables)

	//	第二种初始就赋值，map会根据key-value来确定初始化的长度 第二种更常用
	var mapVariables1 = map[string]int{
		//赋值就以key-value的形式来赋值
		"monday": 1, "Tuesday": 2,
	}
	fmt.Println(mapVariables1)

	//	第三种 通过简短声明 当只需声明一个map时使用make的形式
	mapVariables2 := make(map[string]int)
	mapVariables2["monday"] = 1
	mapVariables2["Tuesday"] = 2
	fmt.Println(mapVariables2)

	//	结构体和对象一样 定义时 type + 结构体名称 + 关键字struct
	type course struct {
		// 定义属性名称 和类型
		courseName    string
		courseTine    float32
		courseTeacher string
	}
	//	定义一个结构体变量并直接赋值
	course1 := course{
		//不指定结构体字段名的方式，严格按照定义结构体时的顺序
		"相信科学", 3.0, "张山",
	}
	fmt.Println(course1)

	//	定义map为结构体
	course2 := make(map[string]course)
	// map中赋值以key-value的形式赋值
	course2["go"] = course1
	course2["美容"] = course1
	fmt.Println(course2)

	//	map+切片
	// 定义
	var mapVariables3 []map[string]interface{}
	mapVariables3 = make([]map[string]interface{},2)
	mapVariables3[0] = make(map[string]interface{},2)
	mapVariables3[0]["name"] = "波哥"
	mapVariables3[0]["age"] = 18

	mapVariables3[1] = make(map[string]interface{},2)
	mapVariables3[1]["name"] = "胡哥"
	mapVariables3[1]["age"] = 16
	for i, m := range mapVariables3 {
		fmt.Println(i, m)
	}

//	判断map中key是否存在
	value,ok := mapVariables2["monday"]
	if ok{
		fmt.Println(value,ok)
	}
//	第二种写法
	if value,ok := mapVariables2["monday"];ok{
		fmt.Println(value,ok)
	}

//	实现map中的key-value交换
//	通过遍历map后重新将key-value赋值
	mapVariables4 := map[int]string{0: "j", 8: "h", 3: "c", 9: "i", 5: "e", 6: "f", 4: "d", 2: "b", 1: "a"}
	mapVariables5 := make(map[string]int)
	for i, s := range mapVariables4 {
		mapVariables5[s] =i
	}
	fmt.Println(mapVariables5)
}

