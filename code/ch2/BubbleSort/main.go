package main

import "fmt"

// 11, 9, 2, 8, 3, 7, 4, 6, 5, 10

func BubbleFindMax(arr []int) int {
	length := len(arr)
	if length <= 1 {
		return arr[0]
	} else {
		for i := 0; i < length-1; i++ {
			if arr[i] > arr[i+1] { // 两两比较
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
		return arr[length-1]
	}
}

func BubbleSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		for i := 0; i < length-1; i++ { // 只剩一个，不需要冒泡了
			isNeedExchange := false // 如果不需要交换就退出
			for j := 0; j < length-i-1; j++ {
				if arr[j] > arr[j+1] {
					arr[j], arr[j+1] = arr[j+1], arr[j]
					isNeedExchange = true
				}
			}
			if !isNeedExchange {
				break
			}
			fmt.Println(arr)
		}
		return arr
	}
}

func main() {
	arr := []int{1, 9, 29, 8, 3, 7, 6, 5, 0}
	fmt.Println(BubbleSort(arr))
}
