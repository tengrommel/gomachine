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
	path := "/home/teng/Documents/git/gomachine/machine/ch2/LinearRegression/Advertising.csv"
	AdFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer AdFile.Close()
	ADDF := dataframe.ReadCSV(AdFile)
	yVals := ADDF.Col("Sales").Float()
	for _, colName := range ADDF.Names() {
		plotVals := make(plotter.XYs, ADDF.Nrow())
		for i, floatVal := range ADDF.Col(colName).Float() {
			plotVals[i].X = floatVal
			plotVals[i].Y = yVals[i]
		}
		p, err := plot.New()
		if err != nil {
			fmt.Println(err)
		}
		p.X.Label.Text = colName
		p.Y.Label.Text = "y"
		p.Add(plotter.NewGrid())
		s, err := plotter.NewScatter(plotVals)
		if err != nil {
			fmt.Println(err)
		}
		s.GlyphStyle.Radius = vg.Points(3)
		p.Add(s)
		if err := p.Save(4*vg.Inch, 4*vg.Inch, colName+"_scatter.png"); err != nil {
			fmt.Println(err)
		}
	}
}
