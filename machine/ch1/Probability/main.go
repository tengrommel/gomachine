package main

import (
	"fmt"
	"github.com/gonum/stat"
	"github.com/gonum/stat/distuv"
)

func main() {
	//ob := []float64{48, 52}
	//ex := []float64{50, 50}
	//fmt.Println(stat.ChiSquare(ob, ex))
	ob := []float64{300, 200, 100}
	//total := 600
	ex := []float64{0.5, 0.3, 0.2}
	fmt.Println(stat.ChiSquare(ob, ex))
	chiDist := distuv.ChiSquared{
		K:   2.0,
		Src: nil,
	}
	fmt.Println(chiDist.Prob(stat.ChiSquare(ob, ex)))
}
