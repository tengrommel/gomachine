package main

import (
	"runtime"
	"sync"
)

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

func ShellSortGoRoutine(arr []int) {
	if len(arr) < 2 || arr == nil {
		return // 数组为空或者数组只有一个元素不需要排序
	}
	wg := sync.WaitGroup{}           // 等待多个线程返回
	GoRoutineNum := runtime.NumCPU() // 抓取系统有几个CPU
	for gap := len(arr) / 2; gap > 0; gap /= 2 {
		wg.Add(GoRoutineNum)
		ch := make(chan int, 10000) // 通道进行线程通信
		go func() {
			// 管道写入任务
			for k := 0; k < gap; k++ {
				ch <- k
			}
			close(ch) // 关闭管道
		}()
		for k := 0; k < GoRoutineNum; k++ {
			go func() {
				for v := range ch {
					ShellSortStep(arr, v, gap) // 完成一个步骤的排序
				}
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func main() {

}
