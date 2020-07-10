package main

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"github.com/montanaflynn/stats"
	"gonum.org/v1/gonum/stat"
	"log"
	"os"
)

func main() {
	irisFile, err := os.Open("2_Matrices_Probability,_and_Statistics/bivariate_analysis/data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()
	irisDF := dataframe.ReadCSV(irisFile)
	sepalLength := irisDF.Col("sepal_length").Float()
	meanVal := stat.Mean(sepalLength, nil)
	modeVal, modeCount := stat.Mode(sepalLength, nil)
	// Calculate the Median of the variable
	medianVal, err := stats.Median(sepalLength)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nSepal Length Summary Statistics:\n")
	fmt.Printf("Mean value: %0.2f\n", meanVal)
	fmt.Printf("Mode value: %0.2f\n", modeVal)
	fmt.Printf("Mean count: %d\n", int(modeCount))
	fmt.Printf("Median value: %0.2f\n", medianVal)
}
