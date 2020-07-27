package main

import "fmt"

func merge(leftArr []int, rightArr []int) []int {
	leftIndex := 0  // 左边的索引
	rightIndex := 0 // 右边的索引
	lastArr := make([]int, 0)
	for leftIndex < len(leftArr) && rightIndex < len(rightArr) {
		if leftArr[leftIndex] < rightArr[rightIndex] {
			lastArr = append(lastArr, leftArr[leftIndex])
			leftIndex++
		} else if leftArr[leftIndex] > rightArr[rightIndex] {
			lastArr = append(lastArr, rightArr[rightIndex])
			rightIndex++
		} else {
			lastArr = append(lastArr, leftArr[leftIndex])
			lastArr = append(lastArr, rightArr[rightIndex])
			leftIndex++
			rightIndex++
		}
	}
	for leftIndex < len(leftArr) { // 把没有结束的归并过来
		lastArr = append(lastArr, leftArr[leftIndex])
		leftIndex++
	}
	for rightIndex < len(rightArr) { // 把没有结束的归并过来
		lastArr = append(lastArr, rightArr[rightIndex])
		rightIndex++
	}
	return lastArr
}

func MergeSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		mid := length / 2
		leftArr := MergeSort(arr[:mid])
		rightArr := MergeSort(arr[mid:])
		return merge(leftArr, rightArr)
	}
}

func main() {
	arr := []int{1, 9, 29, 8, 3, 7, 6, 5, 0}
	fmt.Println(MergeSort(arr))
}
