package main

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"os"
)

// 汇总
func main() {
	path := "/home/teng/Documents/git/gomachine/machine/ch2/LinearRegression/Advertising.csv"
	AdFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer AdFile.Close()
	ADDF := dataframe.ReadCSV(AdFile)
	for _, colName := range ADDF.Names() {
		plotVals := make(plotter.Values, ADDF.Nrow())
		for i, floatVal := range ADDF.Col(colName).Float() {
			plotVals[i] = floatVal
		}
		p, err := plot.New()
		if err != nil {
			fmt.Println(err)
		}
		p.Title.Text = fmt.Sprintf("图的展示字段%s", colName)
		h, err := plotter.NewHist(plotVals, 16)
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
