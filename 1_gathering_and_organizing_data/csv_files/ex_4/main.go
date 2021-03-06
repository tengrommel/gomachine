package main

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"log"
	"os"
)

func main() {
	// Open the CSV file.
	irisFile, err := os.Open("1_gathering_and_organizing_data/csv_files/data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	// Create a dataframe from the CSV file.
	// The types of the columns will be inferred.
	irisDF := dataframe.ReadCSV(irisFile)

	// As a sanity check, display the records to stdout.
	// Gota will format the dataframe for pretty printing.
	fmt.Println(irisDF)
}
