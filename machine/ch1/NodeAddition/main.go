package main

import (
	"fmt"
	. "gorgonia.org/gorgonia"
)

func main() {
	g := NewGraph()   // 描述神经网络
	var x, y, z *Node // 三个节点,神经网络
	var err error
	x = NewScalar(g, Float64, WithName("x"))
	y = NewScalar(g, Float64, WithName("y"))
	z, err = Add(x, y)
	if err != nil {
		fmt.Println(err)
	}
	// 构造虚拟机
	machine := NewTapeMachine(g) //神经网络虚拟机
	Let(x, 2.2)
	Let(y, 2.3)
	if machine.RunAll() != nil {
		fmt.Println(err)
	}
	fmt.Println(z.Value())
}
