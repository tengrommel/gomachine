package main

import "fmt"

func InsertTest(arr []int) []int {
	backup := arr[3]
	j := 3 - 1 // 从上一个位置循环找到位置插入
	for j >= 0 && backup < arr[j] {
		arr[j+1] = arr[j] //从前往后移动
		j--
	}
	arr[j+1] = backup
	return arr
}

func InsertSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		for i := 1; i < length; i++ {
			// 跳过第一个
			backup := arr[i] // 备份插入的数据
			j := i - 1       // 上一个位置循环找到位置插入
			for j >= 0 && backup < arr[j] {
				arr[j+1] = arr[j]
				j--
			}
			arr[j+1] = backup
			fmt.Println(arr)
		}
		return arr
	}
}

// 插入排序
func main() {
	arr := []int{1, 9, 29, 8, 3, 7, 6, 5, 0}
	fmt.Println(InsertSort(arr))
}
