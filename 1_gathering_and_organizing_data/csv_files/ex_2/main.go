package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// Open the iris dataset file.
	f, err := os.Open("1_gathering_and_organizing_data/csv_files/data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	// Create a new CSV reader reading from the opened file.
	reader := csv.NewReader(f)
	reader.FieldsPerRecord = 6
	var rawCSVData [][]string
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		rawCSVData = append(rawCSVData, record)
	}
	fmt.Println(rawCSVData)
}
