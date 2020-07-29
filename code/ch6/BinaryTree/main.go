package main

import "fmt"

func main() {
	bst := NewBinaryTree()
	for i := 1; i <= 7; i++ {
		bst.Add(i)
	}
	fmt.Println(" in --------------------------")
	bst.InOrder()
	fmt.Println(" pre --------------------------")
	bst.PreOrder()
	fmt.Println(" post --------------------------")
	bst.PostOrder()

	fmt.Println("******************************************")
	bst = NewBinaryTree()
	bst.Add(4)
	bst.Add(6)
	bst.Add(2)
	bst.Add(1)
	bst.Add(3)
	bst.Add(5)
	bst.Add(7)
	fmt.Println(" in --------------------------")
	bst.InOrder()
	fmt.Println(" pre --------------------------")
	bst.PreOrder()
	fmt.Println(" post --------------------------")
	bst.PostOrder()
	fmt.Println(bst.String())
}
