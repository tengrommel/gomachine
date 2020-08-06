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
	p, err := plot.New()
	if err != nil {
		fmt.Println(err)
	}
	p.Title.Text = "box plots"
	p.Y.Label.Text = "value"
	w := vg.Points(50)
	for idx, colName := range irisDF.Names() {
		if colName != "species" {
			v := make(plotter.Values, irisDF.Nrow())
			for i, floatVal := range irisDF.Col(colName).Float() {
				v[i] = floatVal
			}
			b, err := plotter.NewBoxPlot(w, float64(idx), v)
			if err != nil {
				fmt.Println(err)
			}
			p.Add(b)
		}
	}
	p.NominalX("sepal_length", "sepal_width", "petal_length", "petal_width")
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "nimei.png"); err != nil {
		fmt.Println(err)
	}
}
