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
	// Open the CSV file.
	f, err := os.Open("4_regression/linear_regression/ex5/training.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	// Create a new CSV reader reading from the opened file
	reader := csv.NewReader(f)
	reader.FieldsPerRecord = 4
	trainingData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	var r regression.Regression
	r.SetObserved("Sales")
	r.SetVar(0, "TV")
	for i, record := range trainingData {
		if i == 0 {
			continue
		}
		yVal, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			log.Fatal(err)
		}
		tvVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Fatal(err)
		}
		r.Train(regression.DataPoint(yVal, []float64{tvVal}))
	}
	r.Run()
	fmt.Printf("\nRegression Formula:\n%v\n", r.Formula)
}
