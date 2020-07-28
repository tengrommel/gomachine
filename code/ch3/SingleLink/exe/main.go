package main

import (
	"fmt"
	"gomachine/code/ch3/SingleLink"
)

func main() {
	list := SingleLink.NewList()
	node1 := SingleLink.NewNode(1)
	node2 := SingleLink.NewNode(2)
	node3 := SingleLink.NewNode(3)
	node4 := SingleLink.NewNode(4)
	list.InsertNodeFront(node1)
	list.InsertNodeFront(node2)
	list.InsertNodeFront(node3)
	list.InsertNodeBack(node4)
	fmt.Println("---")
	fmt.Println(list)
}
