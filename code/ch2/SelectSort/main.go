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

// 1 9 2 8 7 6 4 5
func main() {
	arr := []int{1, 9, 2, 8, 3, 7, 6, 5, 0}
	fmt.Println(SelectSortMax(arr))
}
