package main

import "fmt"

func _compare(a, b interface{}) int {
	var newA, newB int
	var ok bool
	if newA, ok = a.(int); !ok {
		return -2
	}
	if newB, ok = b.(int); !ok {
		return -2
	}
	if newA > newB {
		return 1
	} else if newA < newB {
		return -1
	} else {
		return 0
	}
}

func main() {
	myAvl, _ := NewAVLTree(3, _compare)
	myAvl = myAvl.Insert(2)
	myAvl = myAvl.Insert(1)
	myAvl = myAvl.Insert(4)
	myAvl = myAvl.Insert(5)
	myAvl = myAvl.Insert(6)
	myAvl = myAvl.Insert(7)
	myAvl = myAvl.Insert(15)
	myAvl = myAvl.Insert(26)
	myAvl = myAvl.Insert(17)
	//myAvl = myAvl.Delete(7)
	fmt.Println(myAvl.GetAll())
}
