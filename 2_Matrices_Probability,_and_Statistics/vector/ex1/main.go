package main

import "fmt"

func main() {
	// Initialize a "vector" via a slice.
	var myVector []float64

	// Add a couple of components to the vector
	myVector = append(myVector, 1.0)
	myVector = append(myVector, 2.0)
	fmt.Println(myVector)
}
