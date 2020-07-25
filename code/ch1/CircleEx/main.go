package main

import (
	"fmt"
	"gomachine/code/ch1/CircleQueue"
)

func main() {
	myQ := CircleQueue.CircleQueue{}
	CircleQueue.InitQueue(&myQ)
	CircleQueue.EnQueue(&myQ, 1)
	CircleQueue.EnQueue(&myQ, 2)
	CircleQueue.EnQueue(&myQ, 3)
	CircleQueue.EnQueue(&myQ, 4)
	CircleQueue.EnQueue(&myQ, 5)
	fmt.Println(CircleQueue.DeQueue(&myQ))
	fmt.Println(CircleQueue.DeQueue(&myQ))
	fmt.Println(CircleQueue.DeQueue(&myQ))
	fmt.Println(CircleQueue.DeQueue(&myQ))
	fmt.Println(CircleQueue.DeQueue(&myQ))
}
