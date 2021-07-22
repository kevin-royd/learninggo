package stu

import (
	"fmt"
)

//学生管理结构体
type StudentMgr struct {
	//	对象为学生
	allStudent []*Student
}

//添加方法
func (s *StudentMgr) AddStudent(stu *Student) (err error) {
	for i, v := range s.allStudent {
		if v.UserName == stu.UserName {
			s.allStudent[i] = stu
			return
		}
	}
	s.allStudent = append(s.allStudent, stu)
	return
}

func (s *StudentMgr) ModifyStudent(stu *Student) (err error) {
//	首先查找学生是否存在
	for i, v := range s.allStudent {
		if v.UserName == stu.UserName{
			s.allStudent[i] = stu
		}
		return
	}
	err = fmt.Errorf("该用户：%s不存在", stu.UserName)
	return err
}

func (s *StudentMgr) ShowAllStudent()  {
	for _, v := range s.allStudent {
		fmt.Printf("user:%s info:%#v\n",v.UserName,v)
	}
}