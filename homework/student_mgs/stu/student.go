package stu

type Student struct {
	UserName string
	Sex      int
	Grade    string
	Score    float32
}

func NewStudent(username string, sex int, grade string, score float32) (stu *Student) {
	stu = &Student{
		UserName: username,
		Sex:      sex,
		Score:    score,
		Grade:    grade,
	}
	return stu
}
