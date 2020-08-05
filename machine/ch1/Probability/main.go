package main

import (
	"fmt"
	"github.com/gonum/stat"
	"github.com/gonum/stat/distuv"
)

// 对概率的计算
func main() {
	//ob := []float64{48, 52}
	//ex := []float64{50, 50}
	//fmt.Println(stat.ChiSquare(ob, ex))
	ob := []float64{260, 135, 105}
	total := 500.0
	ex := []float64{total * 0.6, total * 0.25, total * 0.15}
	chi := stat.ChiSquare(ob, ex)
	fmt.Println(total, chi)
	chiDist := distuv.ChiSquared{
		K:   2.0,
		Src: nil,
	}
	fmt.Println(chiDist.Prob(chi))
}
