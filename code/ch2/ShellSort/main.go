package main

import "fmt"

func ShellSortStep(arr []int, start int, gap int) {
	length := len(arr) // 数组的长度
	for i := start + gap; i < length; i += gap {
		// 插入排序的变种
		backup := arr[i] // 备份插入的数据
		j := i - gap     // 上一个位置循环找到位置插入
		for j >= 0 && backup < arr[j] {
			arr[j+gap] = arr[j]
			j -= gap
		}
		arr[j+gap] = backup
	}
}

func ShellSort(arr []int) []int {
	length := len(arr) // 数组长度
	if length <= 1 {
		return arr
	} else {
		gap := len(arr) / 2
		for gap > 0 {
			for i := 0; i < gap; i++ {
				//处理每个元素的步长
				ShellSortStep(arr, i, gap)
			}
			gap /= 2
		}
		return arr
	}
}

func main() {
	arr := []int{1, 9, 29, 8, 3, 7, 6, 5, 0}
	fmt.Println(ShellSort(arr))
}
