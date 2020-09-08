package main

import "fmt"

// 坐标
type point struct {
	i int
	j int
}

// 方向
var dirs = [4]point{
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

// 利用深度遍历走出迷宫
var data = [][]int{
	{0, 1, 0, 1, 1},
	{0, 0, 0, 1, 1},
	{0, 1, 0, 1, 1},
	{0, 1, 0, 0, 1},
	{0, 1, 1, 0, 1},
}

// 调节方向
func (p point) add(dir point) point {
	return point{
		p.i + dir.i, p.j + dir.j,
	}
}

// 判断节点是否可以走
func (p point) at(m [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(m) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(m) {
		return 0, false
	}
	return m[p.i][p.j], true
}

func walkMaze(maze [][]int, start, end point) [][]int {
	steps := make([][]int, len(maze)) // 构造矩阵
	for i := range steps {
		steps[i] = make([]int, len(maze[0])) // 二维数组的每个维度展开一下
	}
	Queue := []point{start} // 开始点放入队列
	for len(Queue) > 0 {
		cur := Queue[0]   // 取得第一个
		Queue = Queue[1:] // 删除第一个
		for _, dir := range dirs {
			next := cur.add(dir)           // 循环每一个方向
			nextValue, ok := next.at(maze) // 判断是否可以走
			if !ok || nextValue == 1 {
				continue
			}
			nextValue, ok = next.at(maze) // 避免走回头路
			if !ok || nextValue != 0 {
				continue
			}
			if next == start { // 避免环路
				continue
			}
			Queue = append(Queue, next) // 可以走的路
			mazeValue, ok := cur.at(steps)
			if ok {
				steps[next.i][next.j] = mazeValue + 1 // 记录位置
			}
		}
	}
	return steps
}

func main() {
	for _, line := range data {
		fmt.Printf("%d\n", line)
	}
	//fmt.Println(data)
	//fmt.Println(dirs)
	steps := walkMaze(data, point{0, 0}, point{4, 4})
	for _, line := range steps {
		fmt.Printf("%d\n", line)
	}
}
