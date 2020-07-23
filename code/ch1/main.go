package main

import (
	"fmt"
	"gomachine/code/ch1/ArrayList"
)

func main() {
	//list := ArrayList.NewArrayList()
	//list.Append(1)
	//list.Append(2)
	//list.Append(3)
	//fmt.Println(list)
	//list := ArrayList.NewArrayList()
	//list.Append("a1")
	//list.Append("b2")
	//list.Append("c3")
	//fmt.Println(list.TheSize) // 小写是私有只能内部使用，大写公有

	//list := ArrayList.NewArrayList()
	//list.Append("a1")
	//list.Append("b2")
	//list.Append("c3")
	//for i:=0;i<10;i++ {
	//	list.Insert(1, "x5")
	//	fmt.Println(list) // 小写是私有只能内部使用，大写公有
	//}
	//fmt.Println("delete")
	//list.Delete(5)
	//fmt.Println(list)
	//list := ArrayList.NewArrayList()
	//list.Append("a1")
	//list.Append("b2")
	//list.Append("c3")
	//list.Append("d3")
	//list.Append("f3")
	//for it := list.Iterator(); it.HasNext(); {
	//	item, _ := it.Next("111111")
	//	if item == "d3" {
	//		it.Remove()
	//	}
	//	fmt.Println(item)
	//}
	//fmt.Println(list)
	//myStack := StackArray.NewStack()
	//myStack.Push(1)
	//myStack.Push(2)
	//myStack.Push(3)
	//myStack.Push(4)
	//fmt.Println(myStack.Pop())
	//fmt.Println(myStack.Pop())
	//fmt.Println(myStack.Pop())
	//fmt.Println(myStack.Pop())
	myStack := ArrayList.NewArrayListStack()
	myStack.Push(1)
	myStack.Push(2)
	myStack.Push(3)
	myStack.Push(4)
	fmt.Println(myStack.Pop())
	fmt.Println(myStack.Pop())
	fmt.Println(myStack.Pop())
	fmt.Println(myStack.Pop())
}
