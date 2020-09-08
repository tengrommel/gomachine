package main

import "fmt"

func printResult(A []int, n int) {
	var i int
	for ; i < n; i++ {
		fmt.Println(A[i])
	}
	fmt.Printf("\n")
}

// Function to generate all binary strings
func generateBinaryStrings(n int, A []int, i int) {
	if i == n {
		printResult(A, n)
		return
	}
	// assign "0" at ith position and try for all other permutations for remaining positions
	A[i] = 0
	generateBinaryStrings(n, A, i+1)
	// assign "1" at ith position and try for all other permutations for remaining positions
	A[i] = 1
	generateBinaryStrings(n, A, i+1)
}

func main() {
	var n = 4
	A := make([]int, n)
	generateBinaryStrings(n, A, 0)
	return
}
