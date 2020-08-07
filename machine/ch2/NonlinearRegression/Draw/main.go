package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func Predict(tv, radio, newspaper float64) float64 {
	return 3.013468 + 0.045176*tv + 0.186239*radio + 0.000609*newspaper
}

func main() {
	file, err := os.Open("test.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 4
	testData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	var mAE float64
	for idx, record := range testData {
		if idx == 0 {
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
		NVal, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			fmt.Println(err)
		}
		yPredict := Predict(TVVal, RadioVal, NVal)
		mAE += math.Abs(yVal-yPredict) / float64(len(testData))

	}
	fmt.Println(mAE)
}
