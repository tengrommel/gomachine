package main

import "fmt"

// 堆排序
func HeapSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		lastMessLen := length - i     // 每次截取一段
		HeapSortMax(arr, lastMessLen) // 0-----------0
		fmt.Println(arr)
		if i < length {
			arr[0], arr[lastMessLen-1] = arr[lastMessLen-1], arr[0]
		}
		fmt.Println("ex:", arr)
	}
	return arr
}

func HeapSortMax(arr []int, length int) []int {
	//length := len(arr) // 求数组的长度
	if length <= 1 {
		return arr
	} else {
		depth := length/2 - 1 // 深度 n 2n+1 2n+2
		for i := depth; i >= 0; i-- {
			// 循环所有的三节点
			topMax := i // 假定最大的在i的位置
			leftChild := 2*i + 1
			rightChild := 2*i + 2 // 左右孩子的节点
			if arr[leftChild] > arr[topMax] {
				topMax = leftChild // 如果左边比我大，记录最大
			}
			if rightChild <= length-1 && arr[rightChild] > arr[topMax] {
				// 防止越界
				topMax = rightChild // 如果右边比我大，记录最大
			}
			if topMax != i {
				// 确保i的值就是最大
				arr[i], arr[topMax] = arr[topMax], arr[i]
			}
		}
	}
	return arr
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
	fmt.Println(HeapSort(arr))
}
