package main

import "fmt"

// 解决递归走出来
func AIOut(AIData [M][N]int, i int, j int) bool {
	AIData[i][j] = 3 // 避免回头路
	if i == M-1 && j == N-1 {
		canGoOut = true
		fmt.Println("迷宫可以走出来")
	} else {
		if j+1 <= N-1 && AIData[i][j+1] < 2 && canGoOut != true {
			AIOut(AIData, i, j+1) // 递归一次
		}
		if i+1 <= M-1 && AIData[i+1][j] < 2 && canGoOut != true {
			AIOut(AIData, i+1, j) // 递归一次
		}
		if j-1 <= 0 && AIData[i][j-1] < 2 && canGoOut != true {
			AIOut(AIData, i, j-1) // 递归一次
		}
		if i-1 >= 0 && AIData[i-1][j] < 2 && canGoOut != true {
			AIOut(AIData, i-1, j) // 递归一次
		}
		if canGoOut != true {
			AIData[i][j] = 0 // 走不通设置为0
		}
	}
	return canGoOut
}
