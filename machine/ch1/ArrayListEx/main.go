package main

import (
	"fmt"
	"gomachine/machine/ch1/ArrayListEx/ArrayList"
)

func main() {
	list := ArrayList.NewArrayList()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	fmt.Println(list)
}
