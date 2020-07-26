package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(3 > 2)
	fmt.Println("a" > "b") // 地址比较
	// a<b<c 首先比较第一个字母，左边小于右边-1，左边大于右边+1，第一个字母比较不成功比较第二个
	fmt.Println(strings.Compare("a", "c"))
	fmt.Println(strings.Compare("a", "b"))
	fmt.Println(strings.Compare("a2", "a1"))
	fmt.Println(strings.Compare("a2", "a2"))
	fmt.Println("a1" > "a")
	pa := "a"
	pb := "b"
	fmt.Println("pa ", &pa)
	fmt.Println("pb ", &pb)
	fmt.Println(&pa, &pb)
	fmt.Println(pa < pb)
	fmt.Println("a1" < "a2")
	fmt.Println("ab" < "ac")
}
