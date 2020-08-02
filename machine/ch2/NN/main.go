package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"math/rand"
	"time"
)

type NetWork struct {
	w      [][][]float64           // 权重
	b      [][]float64             // 偏差
	d      [][]float64             // delta
	z      [][]float64             // 每层layer的z值
	l      int                     // layer层的数量
	ls     []int                   // 每层神经元的数量
	aFunc  func(z float64) float64 // 激活神经元执行函数
	daFunc func(z float64) float64 // 激活神经元执行函数的导数
}

// 数学中的逻辑回归
func sigmod(z float64) float64 {
	return 1.0 / (1.0 + math.Exp(-z))
}

func sigmodD(z float64) float64 {
	return sigmod(z) * (1 - sigmod(z))
}

// 导出神经网络-写入文件
type exportedNetWork struct {
	W  [][][]float64 // 权重
	B  [][]float64   // 偏差
	D  [][]float64   // delta
	Z  [][]float64   // 每层layer的z值
	L  int           // layer层的数量
	Ls []int         // 每层神经元的数量
}

// 神经网络写入磁盘
func (n NetWork) Export(w io.Writer) error {
	bs, err := json.Marshal(exportedNetWork{
		W:  n.w,
		B:  n.b,
		D:  n.d,
		Z:  n.z,
		L:  n.l,
		Ls: n.ls,
	})
	if err != nil {
		return err
	}
	fmt.Fprint(w, string(bs)) // 写入
	return nil
}

// 二进制载入神经网络
func Load(bs []byte) (*NetWork, error) {
	var en exportedNetWork
	err := json.Unmarshal(bs, &en)
	if err != nil {
		return nil, err
	}
	return &NetWork{
		w:      en.W,
		b:      en.B,
		d:      en.D,
		z:      en.Z,
		l:      en.L,
		ls:     en.Ls,
		aFunc:  sigmod,
		daFunc: sigmodD,
	}, nil
}

// 初始化
func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

}
