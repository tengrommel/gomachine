package main

import (
	"fmt"
	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/trees"
	"math"
	"math/rand"
	"time"
)

// 贝叶斯

func main() {
	path := "/home/teng/Documents/git/gomachine/machine/ch3/iris.csv"
	irisData, err := base.ParseCSVToInstances(path, true)
	if err != nil {
		fmt.Println(err)
	}
	rand.Seed(time.Now().UnixNano())
	tree := trees.NewID3DecisionTree(0.6)
	cv, err := evaluation.GenerateCrossFoldValidationConfusionMatrices(irisData, tree, 5)
	if err != nil {
		fmt.Println(err)
	}
	mean, variance := evaluation.GetCrossValidatedMetric(cv, evaluation.GetAccuracy)
	stdev := math.Sqrt(variance)
	fmt.Println(mean, stdev)
}
