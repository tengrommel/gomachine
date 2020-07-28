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
	node5 := SingleLink.NewNode(5)
	list.InsertNodeFront(node1)
	list.InsertNodeFront(node2)
	list.InsertNodeFront(node3)
	list.InsertNodeBack(node4)
	list.InsertNodeValueFront(1, node5)
	fmt.Println("---")
	fmt.Println(list)
}
