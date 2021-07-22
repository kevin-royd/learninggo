package main

import (
	"fmt"
	"os"
)

/*
实现广度迷宫算法
*/
func main() {
	//	首先需要将迷宫导入其中
	filename := "maze/maze.in"
	maze := ReaderMaze(filename)
	//for _, row := range maze {
	//	for _, value := range row {
	//		fmt.Printf("%d ", value)
	//
	//	}
	//	fmt.Println()
	//}
	steps := walk(maze, Point{0, 0},
		Point{len(maze) - 1, len(maze[0]) - 1})

	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}

}

func ReaderMaze(filename string) [][]int {
	//	首先读取文件
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	//	创建声明 行数 列数的变量
	var row, col int
	//	读取文本、将值转换为连续的参数并返回项目数 只有传入内存地址才能发生改变
	fmt.Fscanf(file, "%d %d", &row, &col)
	//	创建一个二维数组 row声明有多少列
	maze := make([][]int, row)
	for i := range maze {
		//	得到单列后声明每行的数量并赋值
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return maze
}

type Point struct {
	//	定义点的坐标位置
	i, j int
}

// 定义4个方向值
var dirs = [4]Point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func walk(maze [][]int, start, end Point) [][]int {
	//	走了多少步才走到这一格
	steps := make([][]int, len(maze)) //得到的行
	for i := range steps {
		steps[i] = make([]int, len(maze[i])) //将maze中点告知steps
	}
	//	创建队列 初始化只有起点 默认值初始为0
	Q := []Point{start}
	for len(Q) > 0 {
		//	发现点
		cur := Q[0]
		Q = Q[1:] //拿掉队头
		//判断是否发现终点、发现终点退出
		if cur == end {
			break
		}
		for _, dir := range dirs {
			//	每次都会发现新的节点 新的节点等于 当前节点加上方向
			next := cur.add(dir) //得到坐标 go中没有方法的重载所以2个坐标无法相加
			//	探索下个节点、下个节点要为探索意思为0，并且下个steps有值表示已经走过了 所有也要为0，并且不能为起点
			value, ok := next.at(maze)
			if !ok || value == 1 { //没有下个节点
				continue
			}
			value, ok = next.at(steps)
			if !ok || value != 0 { //表示已经走过了
				continue
			}
			if next == start {
				continue //起点
			}
			curSteps, _ := cur.at(steps)
			//	将已经走的步数天道steps中
			steps[next.i][next.j] = curSteps + 1
			//	将点添加到队列中
			Q = append(Q, next)
		}
	}
	return steps
}

func (p Point) add(r Point) Point {
	return Point{p.i + r.i, p.j + r.j}
}

// 传入二维数组的范围 返回对应的值和判断是否发生了越界
func (p Point) at(gird [][]int) (int, bool) {
	//对是否越界进行判断
	if p.i < 0 || p.i >= len(gird) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(gird[p.i]) {
		return 0, false
	}
	return gird[p.i][p.j], true
}
