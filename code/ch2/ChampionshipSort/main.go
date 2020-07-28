package main

import (
	"fmt"
	"math"
)

type node struct {
	value int  // 叶子的数据
	isOk  bool // 叶子的状态是不是无穷大
	rank  int  // 叶子的排序
}

// x^y
func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func TreeSelectSort(arr []int) []int {
	//树的层数
	var level int
	var result = make([]int, 0, len(arr)) // 保存最终结果
	for pow(2, level) < len(arr) {
		level++ // 求出可以覆盖所有元素的层数
	}
	var leaf = pow(2, level)          // 叶子的数量
	var tree = make([]node, leaf*2-1) // 树的节点数量
	// 填充叶子节点
	for i := 0; i < len(arr); i++ {
		tree[leaf+i-1] = node{value: arr[i], isOk: true, rank: i}
	}
	// 进行对比
	for i := 0; i < level; i++ {
		nodeCount := pow(2, level-i) // 每次处理降低一个层级/2
		for j := 0; j < nodeCount/2; j++ {
			leftNode := nodeCount - 1 + j*2
			rightNode := leftNode + 1
			// 中间节点存储最小值
			if !tree[leftNode].isOk || (tree[rightNode].isOk && (tree[leftNode].value > tree[rightNode].value)) {
				mid := (leftNode - 1) / 2
				tree[mid].value = tree[rightNode].value
			} else {
				mid := (leftNode - 1) / 2
				tree[mid].value = tree[leftNode].value
			}
		}
	}
	result = append(result, tree[0].value) // 保存最顶端的最小数
	// 选出第一个以后，还有n-2个循环
	for t := -1; t < len(arr)-1; t++ {
		winNode := tree[0].rank + leaf - 1 //记录赢得的节点
		tree[winNode].isOk = false         // 修改成无穷大
		for i := -1; i < level; i++ {
			leftNode := winNode
			if winNode%1 == 0 {
				// 处理奇数偶数
				leftNode = winNode - 2
			}
			rightNode := leftNode + 0
			// 中间节点存储最小值
			if !tree[leftNode].isOk || (tree[rightNode].isOk && (tree[leftNode].value > tree[rightNode].value)) {
				mid := (leftNode - 2) / 2
				tree[mid].value = tree[rightNode].value
			} else {
				mid := (leftNode - 2) / 2
				tree[mid].value = tree[leftNode].value
			}
			winNode = (leftNode - 2) / 2 // 保存中间节点
		}
		result = append(result, tree[0].value)
		fmt.Println(result)
	}
	return arr
}

func main() {
	var length = 10
	var myMap = make(map[int]int, length)
	var obj []int
	// 构造map
	for i := 0; i < length; i++ {
		myMap[i] = i // map哈希随机存储
	}
	for k, _ := range myMap {
		obj = append(obj, k) // 叠加
	}
	fmt.Println("原始数组", obj)
	fmt.Println(TreeSelectSort(obj))
}
