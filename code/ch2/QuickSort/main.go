package main

import "fmt"

func QuickSort(arr []int) []int {
	length := len(arr) // 数组长度
	if length <= 1 {
		return arr // 一个元素的数组，直接放回
	} else {
		splitData := arr[0]       // 以第一个为基准
		low := make([]int, 0, 0)  // 存储比我小的
		high := make([]int, 0, 0) // 存储比我大的
		mid := make([]int, 0, 0)  // 存储与我相等
		for i := 1; i < length; i++ {
			if arr[i] < splitData {
				low = append(low, arr[i])
			} else if arr[i] > splitData {
				high = append(high, arr[i])
			} else {
				mid = append(mid, arr[i])
			}
		}
		low, high = QuickSort(low), QuickSort(high) // 切割递归处理
		myArr := append(append(low, mid...), high...)
		return myArr
	}
}

func main() {
	arr := []int{1, 9, 2, 8, 3, 7, 6, 5, 0}
	fmt.Println(QuickSort(arr))
}
