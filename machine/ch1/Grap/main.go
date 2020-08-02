package main

// 180 180 180
// 180 20 90
// 178 100 130

import (
	"fmt"
	G "gorgonia.org/gorgonia"
	T "gorgonia.org/tensor"
	"io/ioutil"
	"log"
)

func main() {
	g := G.NewGraph() // 神经网络的图
	matB := []float32{0.9, 0.7, 0.4, 0.2}
	matT := T.New(T.WithBacking(matB), T.WithShape(2, 2)) // 调整为二维数组
	mat := G.NewMatrix(g,
		T.Float32,
		G.WithName("W"),
		G.WithShape(2, 2),
		G.WithValue(matT)) // 新建一个矩阵
	vecB := []float32{5, 7}
	vecT := T.New(T.WithBacking(vecB), T.WithShape(2))
	vec := G.NewVector(g, T.Float32, G.WithName("x"), G.WithShape(2), G.WithValue(vecT))
	z, err := G.Mul(mat, vec)
	if err != nil {
		log.Fatal(err)
	}
	machine := G.NewTapeMachine(g)
	if machine.RunAll() != nil {
		fmt.Println(err)
	}
	fmt.Println(z.Value().Data())
	ioutil.WriteFile("graph.dot", []byte(g.ToDot()), 0644)
}
