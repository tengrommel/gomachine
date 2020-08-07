package main

import (
	"encoding/csv"
	"fmt"
	"github.com/gonum/matrix/mat64"
	"os"
	"strconv"
)

// 非线性回归
func main() {
	path := "/home/teng/Documents/git/gomachine/machine/ch2/LinearRegression/Advertising.csv"
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 4         // 字段
	trainData, err := reader.ReadAll() // 读取全部
	if err != nil {
		fmt.Println(err)
	}
	featureData := make([]float64, 4*len(trainData)) // 特征数据
	yData := make([]float64, len(trainData))
	var featureIndex int
	var yIndex int
	for idx, record := range trainData {
		if idx == 0 {
			continue
		}
		for i, val := range record {
			valParsed, err := strconv.ParseFloat(val, 64)
			if err != nil {
				fmt.Println(err)
			}
			if i < 3 {
				if i == 0 {
					featureData[featureIndex] = 1
					featureIndex++
				}
				featureData[featureIndex] = valParsed
				featureIndex++
			}
			if i == 3 {
				yData[yIndex] = valParsed
				yIndex++
			}
		}
	}
	features := mat64.NewDense(len(trainData), 4, featureData) // 创建矩阵
	y := mat64.NewVector(len(trainData), yData)                // 创建向量
	fmt.Println(features)
	fmt.Println(y)
}
