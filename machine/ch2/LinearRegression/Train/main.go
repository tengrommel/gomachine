package main

import (
	"encoding/csv"
	"fmt"
	"github.com/sajari/regression"
	"log"
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
	r.SetVar(0, "TV") // x, y之间的关系
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
		r.Train(regression.DataPoint(yVal, []float64{TVVal}))
	}
	r.Run()
	fmt.Println(r.Formula)
}
