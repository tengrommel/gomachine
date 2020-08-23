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
} // 移动的四个方向

var data = [][]int{
	{0, 1, 0, 1, 1},
	{0, 0, 0, 1, 1},
	{0, 1, 0, 1, 1},
	{0, 1, 0, 0, 1},
	{0, 1, 1, 0, 1},
}

// 调整方向
func (p point) add(dir point) point {
	return point{
		p.i + dir.i,
		p.j + dir.j,
	}
}

// 判断的节点是否可以走
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
		steps[i] = make([]int, len(maze[0]))
	}
	Queue := []point{start} // 开始点放入队列
	for len(Queue) > 0 {
		cur := Queue[0]   // 取得第一个
		Queue = Queue[1:] // 删除第一个
		for _, dir := range dirs {
			next := cur.add(dir)
			nextValue, ok := next.at(maze) // 判断是否可
			if !ok || nextValue == 1 {
				continue
			}
			nextValue, ok = next.at(steps) // 避免回头路
			if !ok || nextValue == 1 {
				continue
			}
			if next == start { // 避免环路
				continue
			}
			Queue = append(Queue, next)
			mazeValue, ok := cur.at(steps)
			if ok {
				steps[next.i][next.j] = mazeValue + 1
			}
		}
	}
	return steps
}

func main() {

	fmt.Println(data)
	fmt.Println(dirs)
}
