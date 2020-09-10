package main

import (
	"fmt"
)

// 通过二维数组存储关系
var Nodes = 6
var FindMaps = [9][3]int{
	{1, 2, 1},
	{1, 3, 12},
	{2, 3, 9},

	{2, 4, 3},
	{3, 5, 5},
	{4, 3, 4},

	{4, 5, 13},
	{4, 6, 15},
	{5, 6, 4},
}

var inf = 2<<31 - 1 // 最大代表不可联通
var min int         // 记录最短

func main() {
	var Maps [6][6]int
	for i := 0; i < Nodes; i++ {
		for j := 0; j < Nodes; j++ {
			if i == j {
				Maps[i][j] = 0 // 自己到自己0
			} else {
				Maps[i][j] = -1
			}
		}
	}
	// 读取边长
	for _, v := range FindMaps {
		Maps[v[0]-1][v[1]-1] = v[2]
	}
	for i := 0; i < Nodes; i++ {
		fmt.Printf("%1d\n", Maps[i])
	}
	var dis [6]int // 数组保存距离
	for k, v := range Maps[0] {
		dis[k] = v // 拷贝第一行的数据
	}
	fmt.Println(dis)
	var book [6]int
	book[0] = 1 // 设定从这里开始
	var u int   // 记录节点
	for i := 0; i < Nodes-1; i++ {
		// 寻找距离1号最近的点
		min = inf // 设置当前最大距离
		for j := 0; j < Nodes; j++ {
			if book[j] == 0 && dis[j] < min {
				min = dis[j] // 保存最短
				u = j        // 保存最短的索引
			}
		}
		book[u] = 1
		for v := 0; v < Nodes; v++ {
			if Maps[u][v] != inf && Maps[u][v] != 0 {
				if dis[v] > dis[u]+Maps[u][v] {
					dis[v] = dis[u] + Maps[u][v] // 保存最短距离
				}
			}
		}
	}
}
