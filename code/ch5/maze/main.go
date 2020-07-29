package main

import "fmt"

func main() {
	show(data)
	for {
		var inputString string
		fmt.Scanln(&inputString)
		run(inputString)
	}
}
