package main

import (
	"fmt"
	G "gorgonia.org/gorgonia"
	T "gorgonia.org/tensor"
	"log"
	"runtime"
)

func main() {
	g := G.NewGraph()
	x := G.NewMatrix(g, G.Float32, G.WithName("x"), G.WithShape(100, 100))
	y := G.NewMatrix(g, G.Float32, G.WithName("y"), G.WithShape(100, 100))
	xpy := G.Must(G.Add(x, y))
	xpy2 := G.Must(G.Tanh(xpy))
	m := G.NewTapeMachine(g, G.UseCudaFor("tanh"))
	G.Let(x, T.New(T.WithShape(100, 100), T.WithBacking(T.Random(T.Float32, 100*100))))
	G.Let(y, T.New(T.WithShape(100, 100), T.WithBacking(T.Random(T.Float32, 100*100))))
	runtime.LockOSThread() // GPU执行
	for i := 0; i < 1000; i++ {
		if err := m.RunAll(); err != nil {
			log.Fatalf("iteration: %d, err: %v", i, err)
		}
	}
	runtime.UnlockOSThread()
	fmt.Printf("%1.1f", xpy2.Value())
}
