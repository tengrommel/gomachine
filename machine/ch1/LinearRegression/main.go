package main

import (
	"fmt"
	G "gorgonia.org/gorgonia"
	"gorgonia.org/tensor"
	"math/rand"
	"runtime"
	"time"
)

const vecSize = 1000000

func xy(dt tensor.Dtype) (x tensor.Tensor, y tensor.Tensor) {
	var xBack, yBack interface{}
	switch dt {
	case G.Float32:
		xBack = tensor.Range(tensor.Float32, 1, vecSize+1).([]float32)
		yBackC := tensor.Range(tensor.Float32, 1, vecSize+1).([]float32)
		for i, v := range yBackC {
			yBackC[i] = v*2 + rand.Float32()
		}
		yBack = yBackC
	case G.Float64:
		xBack = tensor.Range(tensor.Float64, 1, vecSize+1).([]float64)
		yBackC := tensor.Range(tensor.Float64, 1, vecSize+1).([]float64)
		for i, v := range yBackC {
			yBackC[i] = v*2 + rand.Float64()
		}
		yBack = yBackC
	}
	x = tensor.New(tensor.WithBacking(xBack), tensor.WithShape(vecSize))
	y = tensor.New(tensor.WithBacking(yBack), tensor.WithShape(vecSize))
	return
}

// 返回随机数
func random(dt tensor.Dtype) interface{} {
	rand.Seed(time.Now().UnixNano())
	switch dt {
	case tensor.Float32:
		return rand.Float32()
	case tensor.Float64:
		return rand.Float64()
	default:
		panic("未知类型")
	}
}

func LinearSetup(Float tensor.Dtype) (m, c *G.Node, machine G.VM) {
	var xT, yT G.Value
	xT, yT = xy(Float)

	g := G.NewGraph()
	x := G.NewVector(g, Float, G.WithShape(vecSize), G.WithName("x"), G.WithValue(xT))
	y := G.NewVector(g, Float, G.WithShape(vecSize), G.WithName("y"), G.WithValue(yT))
	m = G.NewScalar(g, Float, G.WithName("m"), G.WithValue(random(Float)))
	c = G.NewScalar(g, Float, G.WithName("c"), G.WithValue(random(Float)))
	pred := G.Must(G.Add(G.Must(G.Mul(x, m)), c))  // m * x + c
	se := G.Must(G.Square(G.Must(G.Sub(pred, y)))) // y -m*x-c 趋近于零
	cost := G.Must(G.Mean(se))                     // 取得中位数
	if _, err := G.Grad(cost, m, c); err != nil {
		fmt.Println(err)
	}
	machine = G.NewTapeMachine(g, G.BindDualValues(m, c)) // y = m * x + c
	return m, c, machine
}

func LinearRun(m, c *G.Node, machine G.VM, iter int, autoCleanup bool) (retM, retC G.Value) {
	if autoCleanup {
		defer machine.Close()
	}
	model := []G.ValueGrad{m, c}
	solver := G.NewVanillaSolver(G.WithLearnRate(0.001), G.WithClip(5))
	var err error
	for i := 0; i < iter; i++ {
		if err = machine.RunAll(); err != nil {
			fmt.Println(i, err) // 执行第i次发现错误
			break
		}
		if err = solver.Step(model); err != nil {
			fmt.Println(err)
		}
		machine.Reset() // 重置
	}
	return m.Value(), c.Value()
}

func linearRegression(Float tensor.Dtype, iter int) (retM, retC G.Value) {
	m, c, machine := LinearSetup(Float)
	defer runtime.GC()
	return LinearRun(m, c, machine, iter, true)
}

func main() {
	var m, c G.Value
	m, c = linearRegression(G.Float32, 500)
	fmt.Printf("Float32:y=%3.3fx+%3.3f\n", m, c)
	m, c = linearRegression(G.Float64, 500)
	fmt.Printf("Float64:y=%3.3fx+%3.3f", m, c)
}
