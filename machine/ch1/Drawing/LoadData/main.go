package main

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"os"
)

func main() {
	path := "1_gathering_and_organizing_data/csv_files/data/iris_labeled.csv"
	irisFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer irisFile.Close()
	irisDF := dataframe.ReadCSV(irisFile)
	for _, colName := range irisDF.Names() {
		if colName != "species" {
			v := make(plotter.Values, irisDF.Nrow())
			for i, floatVal := range irisDF.Col(colName).Float() {
				v[i] = floatVal
			}
			p, err := plot.New()
			if err != nil {
				fmt.Println(err)
			}
			p.Title.Text = fmt.Sprintf("histogram of a %s", colName)
			h, err := plotter.NewHist(v, 16)
			if err != nil {
				fmt.Println(err)
			}
			h.Normalize(1)
			p.Add(h)
			if err := p.Save(4*vg.Inch, 4*vg.Inch, colName+"_hist.png"); err != nil {
				fmt.Println(err)
			}
		}
	}
}
