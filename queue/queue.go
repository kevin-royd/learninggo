package queue

/*
队列
*/

//定义类型别名
type Queue []int

func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

// 返回值类型int
func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

// 判断是否为空
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
