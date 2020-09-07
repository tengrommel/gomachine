package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

var factVal uint64 = 1

func factorialIterative(n int) uint64 {
	if n < 0 {
		fmt.Print("Factorial of negative number doesn't exist.")
	} else {
		for i := 1; i <= n; i++ {
			factVal *= uint64(i)
		}
	}
	return factVal
}

func main() {
	fmt.Println(factorialIterative(5))
}
