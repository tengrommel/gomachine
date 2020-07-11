package main

import (
	"bufio"
	"github.com/go-gota/gota/dataframe"
	"log"
	"os"
)

func main() {
	// Open the CSV file.
	advertFile, err := os.Open("4_regression/linear_regression/ex4/Advertising.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer advertFile.Close()
	advertDF := dataframe.ReadCSV(advertFile)
	trainingNum := (4 * advertDF.Nrow()) / 5
	testNum := advertDF.Nrow() / 5
	if trainingNum+testNum < advertDF.Nrow() {
		trainingNum++
	}
	// Create the subset indices
	trainingIdx := make([]int, trainingNum)
	testIdx := make([]int, testNum)
	for i := 0; i < trainingNum; i++ {
		trainingIdx[i] = i
	}
	for i := 0; i < testNum; i++ {
		testIdx[i] = trainingNum + i
	}
	// Create the subset dataframes.
	trainingDF := advertDF.Subset(trainingIdx)
	testDF := advertDF.Subset(testIdx)
	// Create a map that will be used in writing the data
	// to files.
	setMap := map[int]dataframe.DataFrame{
		0: trainingDF,
		1: testDF,
	}
	// Create the respective files.
	for idx, setName := range []string{"training.csv", "test.csv"} {
		// Save the filtered dataset file.
		f, err := os.Create(setName)
		if err != nil {
			log.Fatal(err)
		}
		// Create a buffered writer.
		w := bufio.NewWriter(f)
		// Write the dataframe out as a CSV.
		if err := setMap[idx].WriteCSV(w); err != nil {
			log.Fatal(err)
		}
	}
}
