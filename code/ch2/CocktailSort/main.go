package main

import "fmt"

// 鸡尾酒排序
// 正向冒泡一次
// 反向冒泡一次
func cocktail(arr []int) []int {
	for i := 0; i < len(arr)/2; i++ {
		// 每次循环，正向冒泡一次，反向冒泡一次
		left := 0
		right := len(arr) - 1 // 左边， 右边
		for left <= right {
			// 结束的条件
			if arr[left] > arr[left+1] {
				arr[left], arr[left+1] = arr[left+1], arr[left]
			}
			left++
			if arr[right-1] > arr[right] {
				arr[right-1], arr[right] = arr[right], arr[right-1]
			}
			right--
		}
	}
	return arr
}

func main() {
	arr := []int{1, 9, 29, 8, 3, 7, 6, 5, 0}
	fmt.Println(cocktail(arr))
}
