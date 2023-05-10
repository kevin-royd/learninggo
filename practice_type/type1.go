package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u User) ChangeName(Name string) User {
	u.Name = Name
	return u
}

func (u *User) ChangeAge(Age int) *User {
	u.Age = Age
	return u
}

func main() {
	u := User{
		Name: "Zhangshan",
		Age:  18,
	}
	//定义了接收器才会有该方法。相当于其他语言的构造器 返回值类型为值类型
	u.ChangeName("lisi")
	u.ChangeAge(20)
	fmt.Printf("u=%v\n", u)

	//测试返回类型为指针类型
	up := &User{
		Name: "wangwu",
		Age:  19,
	}
	up.ChangeName("liuliu")
	up.ChangeAge(21)
	fmt.Printf("up=%v\n", up)

	//输出u={Zhangshan 20}
	//up=&{wangwu 21}
	//结论：只有指针接收器才能对值进行修改

}
