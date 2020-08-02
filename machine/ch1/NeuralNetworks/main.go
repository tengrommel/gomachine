package main

import (
	"fmt"
	G "gorgonia.org/gorgonia"
	"gorgonia.org/tensor"
	"math/rand"
	"time"
)

var err error

type nn struct {
	g         *G.ExprGraph // 图
	w0, w1    *G.Node      // 节点
	pred      *G.Node      //前置节点
	predValue G.Value
}

// 根据图构造神经网络
func newNN(g *G.ExprGraph) *nn {
	wB := tensor.Random(tensor.Float64, 3)
	wT := tensor.New(tensor.WithBacking(wB), tensor.WithShape(3, 1))
	w0 := G.NewMatrix(g, tensor.Float64, G.WithName("w"), G.WithShape(3, 1), G.WithValue(wT))
	return &nn{g: g, w0: w0}
}

func (m *nn) learnAbles() G.Nodes {
	return G.Nodes{m.w0}
}

func (m *nn) fwd(x *G.Node) error {
	var l0, l1 *G.Node
	l0 = x
	l0Dot := G.Must(G.Mul(l0, m.w0)) // 计算矩阵乘法
	l1 = G.Must(G.Sigmoid(l0Dot))    // 处理隐藏的结果
	m.pred = l1
	G.Read(m.pred, &m.predValue) // 读取
	return nil
}

func main() {
	rand.Seed(time.Now().UnixNano())
	g := G.NewGraph()
	m := newNN(g)
	xB := []float64{0, 0, 1,
		0, 1, 1,
		1, 0, 1,
		1, 1, 1}
	xT := tensor.New(tensor.WithBacking(xB), tensor.WithShape(4, 3))
	x := G.NewMatrix(g, tensor.Float64, G.WithName("x"), G.WithShape(4, 3), G.WithValue(xT))
	yB := []float64{0, 0, 1, 1}
	yT := tensor.New(tensor.WithBacking(yB), tensor.WithShape(4, 1))
	y := G.NewMatrix(g, tensor.Float64, G.WithName("y"), G.WithShape(4, 1), G.WithValue(yT))
	if err := m.fwd(x); err != nil {
		fmt.Printf("%+v", err)
	}
	losses := G.Must(G.Sub(y, m.pred))
	square := G.Must(G.Square(losses)) //损失
	cost := G.Must(G.Mean(square))     // 损失函数
	// 更新验证结果
	if _, err := G.Grad(cost, m.learnAbles()...); err != nil {
		fmt.Println(err)
	}
	vm := G.NewTapeMachine(g, G.BindDualValues(m.learnAbles()...))
	defer vm.Close()
	solver := G.NewVanillaSolver(G.WithLearnRate(1.0))
	// 10000次的训练
	for i := 0; i < 10000; i++ {
		solver = G.NewVanillaSolver(G.WithLearnRate(1.0))
		for i := 0; i < 10000; i++ {
			vm.Reset()
			if err = vm.RunAll(); err != nil { // 训练
				fmt.Println(err)
			}
			solver.Step(G.NodesToValueGrads(m.learnAbles())) // 校验反馈
			vm.Reset()
		}
		fmt.Println("训练之后结果:\n", m.predValue)
	}
}
