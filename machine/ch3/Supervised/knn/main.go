package main

import (
	"fmt"
	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	knn2 "github.com/sjwhitworth/golearn/knn"
	"math"
)

func main() {
	path := "/home/teng/Documents/git/gomachine/machine/ch3/iris.csv"
	irisData, err := base.ParseCSVToInstances(path, true)
	if err != nil {
		fmt.Println(err)
	}
	knn := knn2.NewKnnClassifier("euclidean", "linear", 2)
	cv, err := evaluation.GenerateCrossFoldValidationConfusionMatrices(irisData, knn, 5)
	if err != nil {
		fmt.Println(err)
	}
	mean, variance := evaluation.GetCrossValidatedMetric(cv, evaluation.GetAccuracy)
	stdev := math.Sqrt(variance)
	fmt.Println(mean)
	fmt.Println(stdev)
}
