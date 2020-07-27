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

func RadixSort(arr []int) []int {
	max := SelectSortMax(arr) // 寻找数组的极大值
	fmt.Println(max)
	for bit := 1; max/bit > 0; bit *= 10 {
		// 按照数量级分段
		arr = BitSort(arr, bit) // 每次处理一个级别的排序
		fmt.Println(arr)
	}
	return arr
}

func BitSort(arr []int, bit int) []int {
	length := len(arr)           // 数组长度
	bitCounts := make([]int, 10) // 统计长度
	for i := 0; i < length; i++ {
		num := (arr[i] / bit) % 10
		bitCounts[num]++ // 统计余数相等的个数
	}
	fmt.Println(bitCounts)
	for i := 1; i < 10; i++ {
		bitCounts[i] += bitCounts[i-1] // 叠加，计算位置
	}
	fmt.Println(bitCounts)
	tmp := make([]int, 10) // 开辟临时数组
	for i := length - 1; i >= 0; i-- {
		num := (arr[i] / bit) % 10     // 分层处理
		tmp[bitCounts[num]-1] = arr[i] // 计算排序的位置
		bitCounts[num]--
	}
	for i := 0; i < length; i++ {
		arr[i] = tmp[i]
	}
	return arr
}

// 银行，每个人都有存款
// 100 1000 10000
// 银行存款，金额 不固定
func main() {
	arr := []int{11, 91, 222, 878, 348, 7123, 4213, 5123, 1011}
	fmt.Println(RadixSort(arr))
}
