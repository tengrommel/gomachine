package main

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"log"
	"os"
)

func main() {
	// Open the CSV file.
	advertFile, err := os.Open("4_regression/linear_regression/ex2/Advertising.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer advertFile.Close()
	advertDF := dataframe.ReadCSV(advertFile)
	for _, colName := range advertDF.Names() {
		plotVals := make(plotter.Values, advertDF.Nrow())
		for i, floatVal := range advertDF.Col(colName).Float() {
			plotVals[i] = floatVal
		}
		// Make a plot and set its title
		p, err := plot.New()
		if err != nil {
			log.Fatal(err)
		}
		p.Title.Text = fmt.Sprintf("Histogram of a %s", colName)
		h, err := plotter.NewHist(plotVals, 16)
		if err != nil {
			log.Fatal(err)
		}
		// Normalize the histogram
		h.Normalize(1)
		p.Add(h)
		if err := p.Save(4*vg.Inch, 4*vg.Inch, colName+"hist.png"); err != nil {
			log.Fatal(err)
		}
	}
}
