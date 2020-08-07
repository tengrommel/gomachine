package main

import (
	"encoding/csv"
	"fmt"
	"github.com/sajari/regression"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	// Open the advertising dataset file.
	file, err := os.Open("training.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 4         // 字段
	trainData, err := reader.ReadAll() // 读取全部
	if err != nil {
		fmt.Println(err)
	}
	var r regression.Regression // 线性回归对象
	r.SetObserved("Sales")
	r.SetVar(0, "TV")
	r.SetVar(1, "Radio")
	for i, record := range trainData {
		if i == 0 {
			continue
		}
		yVal, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			fmt.Println(err)
		}
		TVVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			fmt.Println(err)
		}
		RadioVal, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			fmt.Println(err)
		}
		r.Train(regression.DataPoint(yVal, []float64{TVVal, RadioVal}))
	}
	r.Run()
	fmt.Println(r.Formula)

	testPath := "test.csv"
	fileTest, err := os.Open(testPath)
	if err != nil {
		fmt.Println(err)
	}
	defer fileTest.Close()
	reader = csv.NewReader(fileTest)
	reader.FieldsPerRecord = 4
	testData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	var mAE float64 // 误差
	for i, record := range testData {
		if i == 0 {
			continue
		}
		yVal, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			fmt.Println(err)
		}
		TVVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			fmt.Println(err)
		}
		RadioVal, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			fmt.Println(err)
		}
		// 预测的yPredicted, 正确yVal
		yPredicted, err := r.Predict([]float64{TVVal, RadioVal})
		mAE += math.Abs(yVal-yPredicted) / float64(len(testData))
	}
	fmt.Println(mAE)
}
