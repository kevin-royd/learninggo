package main

import (
	"fmt"
	"learninggo/homework/student_mgs/stu"
	"os"
)

/*
通过对象简单实现学生管理 添加修改展现
*/

//学生对象为全局对象
var (
	StudentMgr = &stu.StudentMgr{}
)

func main() {
	//	展示菜单相关信息
	for {
		ShowMenu()
		var sel int
		fmt.Scanf("%d\n", &sel)
		switch sel {
		case 1:
			stu := InputStudent()
			StudentMgr.AddStudent(stu)
		case 2:
			stu := InputStudent()
			StudentMgr.AddStudent(stu)
		case 3:
			StudentMgr.ShowAllStudent()
		case 4:
			os.Exit(0)
		}
	}
}

func ShowMenu() {
	fmt.Println("1. 新增学生信息")
	fmt.Println("2. 修改学生信息")
	fmt.Println("3. 展示全部学生信息")
	fmt.Println("4. exit")
}

//用户输入函数
func InputStudent() *stu.Student {
	var (
		username string
		sex      int
		grade    string
		score    float32
	)
	fmt.Println("请输入用户名:")
	fmt.Scanf("%s\n", &username)
	fmt.Println("请输入性别:")
	fmt.Scanf("%s\n", &sex)
	fmt.Println("请输入年级:")
	fmt.Scanf("%s\n", &grade)
	fmt.Println("请输入成绩:")
	fmt.Scanf("%s\n", &score)
	stu := stu.NewStudent(username, sex, grade,score)
	return stu
}
