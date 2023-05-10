package main

import "fmt"

func main() {
	/*
		结构体	是值类型，用于自定义数据类型与实现面向对象
	*/
	//	结构的使用：第一种
	//	首先需要定义一个结构体
	type userinfo struct {
		name      string
		age       int
		height    float32
		eduschool string
		hobby     []string
		moreinfo  map[string]interface{}
	}
	//	在通过引用进行赋值使用
	var bobo userinfo
	bobo.name = "张山"
	bobo.age = 18
	bobo.height = 180.3
	bobo.eduschool = "伯明翰"
	bobo.hobby = []string{"coding", "运动"}
	bobo.moreinfo = map[string]interface{}{
		"work": "百度", "duty": "产品",
	}
	fmt.Println(bobo)

	//	第二种 简短式
	zhangShan := userinfo{
		name:      "张山",
		age:       19,
		height:    169.9,
		eduschool: "剑桥",
		hobby:     []string{"coding", "篮球"},
		moreinfo: map[string]interface{}{
			"work": "百度", "duty": "产品",
		},
	}
	fmt.Println(zhangShan)

	//	第三种 不指定字段名直接赋值 需要严格按照声明时字段添加
	liSi := userinfo{
		"张山", 20, 179.9, "牛津", []string{"coding", "足球"}, map[string]interface{}{"work": "百度", "duty": "产品"},
	}
	fmt.Println(liSi)

	//	结构体的使用方式3 new 使用new new(int),new(string),new(T) 返回结构体指针
	var xiaoming *userinfo
	xiaoming = new(userinfo)
	(*xiaoming).name = "小明"
	(*xiaoming).age = 18
	(*xiaoming).eduschool = "伦敦大学"
	//xiaoming->(*xiaoming) go语言编译器自动转换
	xiaoming.hobby = []string{"coding", "绘画"}
	fmt.Println(xiaoming)

	//	结构体使用方式4 地址类型 同样返回结构体指针
	var xiaohong *userinfo = &userinfo{
		"小红", 12, 120, "小学", []string{"学习", "玩", "打游戏"}, map[string]interface{}{"年级": "五年级"},
	}
	fmt.Println(xiaohong)

	//	结构体还可以做为另一个结构体的类型
	type role struct {
		user          userinfo
		authorization int
	}

	superadmin := role{
		user: userinfo{
			name:      "超级管理员",
			age:       0,
			height:    0,
			eduschool: "",
			hobby:     nil,
			moreinfo:  nil,
		},
		authorization: 1,
	}
	admin := role{
		user: userinfo{
			name:      "管理员",
			age:       0,
			height:    0,
			eduschool: "",
			hobby:     nil,
			moreinfo:  nil,
		},
		authorization: 2,
	}
	fmt.Println(superadmin, admin)

}

// 结构体类型 A B
type Handler func()

// 此时就没有办法使用指针;除此之外遇事不决用指针 (H Handler) 为结构体接收器。无法对值进行修改。
func (H Handler) Hello() {

}
