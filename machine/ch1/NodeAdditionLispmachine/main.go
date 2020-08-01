package main

import (
	"fmt"
	. "gorgonia.org/gorgonia"
	"log"
)

func main() {
	g := NewGraph()
	var x, y, z *Node
	var err error
	// define the expression
	x = NewScalar(g, Float64, WithName("x"))
	y = NewScalar(g, Float64, WithName("y"))
	z, err = Add(x, y)
	if err != nil {
		log.Fatal(err)
	}
	Let(x, 2.0)
	Let(y, 2.5)
	m := NewLispMachine(g) // 后台执行
	if m.RunAll() != nil {
		log.Fatal(err)
	}
	fmt.Printf("z: %v\n", z.Value())
	xGrad, err := x.Grad() // 如果有梯度返回梯度
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("dz/dx: %v\n", xGrad)
	yGrad, err := y.Grad()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("dz/dy: %v\n", yGrad)
}
