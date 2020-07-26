package main

import "fmt"

func SelectSortMax(arr []int) int {
	length := len(arr) // 求数组的长度
	if length <= 1 {
		return arr[0]
	} else {
		max := arr[0] // 假定第一个最大
		for i := 1; i < length; i++ {
			if arr[i] > max {
				max = arr[i]
			}
		}
		return max
	}
}

func SelectSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr // 一个元素的数组
	} else {
		for i := 0; i < length-1; i++ { // 只剩一个元素不需要挑选
			min := i                          // 标记索引
			for j := i + 1; j < length; j++ { // 每次选出一个极小值
				if arr[min] < arr[j] {
					min = j // 保存极小值的索引
				}
			}
			if i != min {
				arr[i], arr[min] = arr[min], arr[i] // 数据交换
			}
		}
	}
	return arr
}

// 1 9 2 8 7 6 4 5
func main() {
	arr := []int{1, 9, 2, 8, 3, 7, 6, 5, 0}
	//fmt.Println(SelectSortMax(arr))
	fmt.Println(SelectSort(arr))
}
