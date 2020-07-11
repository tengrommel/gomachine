package main

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"log"
	"os"
)

func main() {
	// Open the CSV file.
	advertFile, err := os.Open("4_regression/linear_regression/ex1/Advertising.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer advertFile.Close()
	advertDF := dataframe.ReadCSV(advertFile)
	// Use the Describe method to calculate summary statistics
	// for all of the columns in one shot.
	advertSummary := advertDF.Describe()
	// Output the summary statistics to stdout
	fmt.Println(advertSummary)
}
