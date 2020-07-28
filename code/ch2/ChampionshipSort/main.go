package main

import "fmt"

type node struct {
	value int  // 叶子的数据
	isOk  bool // 叶子的状态是不是无穷大
	rank  int  // 叶子的排序
}

func main() {
	var length = 10
	var myMap = make(map[int]int, length)
	var obj []int
	// 构造map
	for i := 0; i < length; i++ {
		myMap[i] = i
	}
	for k, _ := range myMap {
		obj = append(obj, k) // 叠加
	}
	fmt.Println(obj)
}
