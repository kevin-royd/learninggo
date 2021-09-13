package main

import (
	"log"
	"reflect"
)

func main() {
	// 定义变量
	x := float64(3.14)
	// 设置指时不能对值直接进行操作，所以需要传入指针类型
	reflect_set_value(&x)
	log.Println(x)
}

func reflect_set_value(a interface{}){
	// 直接获取a对象的指
	valueOf := reflect.ValueOf(a)
	// 涉及到值得操作都为指针类型 ptr
	kind := valueOf.Kind()
	// 流程控制判断kind的类型和反射的类型是否相同
	switch kind{
	case reflect.Ptr:
		log.Println("这是指针类型")
		//这里需要和原来值数值类型相同
		valueOf.Elem().SetFloat(6.28)
	}
}