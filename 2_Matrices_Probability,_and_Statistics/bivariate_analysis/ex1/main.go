package main

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"github.com/montanaflynn/stats"
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
	sepalWidth := irisDF.Col("sepal_width").Float()
	cor, _ := stats.Correlation(sepalLength, sepalWidth)
	fmt.Println(cor)
}
