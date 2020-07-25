package main

import "fmt"

func main() {
	//node1 := new(Node)
	//node2 := new(Node)
	//node3 := new(Node)
	//node4 := new(Node)
	//node1.data = 1
	//node1.pNext = node2
	//node2.data = 2
	//node2.pNext = node3
	//node3.data = 3
	//node3.pNext = node4
	//node4.data = 4
	//fmt.Println(node1.data)
	//fmt.Println(node2.data)
	//fmt.Println(node3.data)
	//fmt.Println(node4.data)
	//fmt.Println("------------------")
	//fmt.Println(node1.pNext.data, node1.pNext.pNext.data)
	//myStack := NewStack()
	//for i:=0;i<1000;i++ {
	//	myStack.Push(i)
	//}
	//for data := myStack.Pop();data != nil; data=myStack.Pop() {
	//	fmt.Println(data)
	//}
	myQueue := NewLinkQueue()
	for i := 0; i < 100; i++ {
		myQueue.Enqueue(i)
	}
	for data := myQueue.Dequeue(); data != nil; data = myQueue.Dequeue() {
		fmt.Println(data)
	}
}
