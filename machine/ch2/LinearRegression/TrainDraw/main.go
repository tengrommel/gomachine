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

func Predict(Radio float64) float64 {
	return 9.32 + Radio*0.19
}

func main() {
	file, err := os.Open("training.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	ADDF := dataframe.ReadCSV(file)
	yVals := ADDF.Col("Sales").Float()
	pts := make(plotter.XYs, ADDF.Nrow())     // 实际数据
	ptsPred := make(plotter.XYs, ADDF.Nrow()) // 预览数据
	for i, floatVal := range ADDF.Col("Radio").Float() {
		pts[i].X = floatVal
		pts[i].Y = yVals[i]
		ptsPred[i].X = floatVal
		ptsPred[i].Y = Predict(floatVal)
	}
	p, err := plot.New()
	if err != nil {
		fmt.Println(err)
	}
	p.X.Label.Text = "Radio"
	p.Y.Label.Text = "Sales"
	p.Add(plotter.NewGrid())
	s, err := plotter.NewScatter(pts)
	if err != nil {
		fmt.Println(err)
	}
	s.GlyphStyle.Radius = vg.Points(3)
	l, err := plotter.NewLine(ptsPred)
	if err != nil {
		fmt.Println(err)
	}
	l.LineStyle.Width = vg.Points(1)
	l.LineStyle.Dashes = []vg.Length{vg.Points(5), vg.Points(5)}
	p.Add(s, l)
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "last.png"); err != nil {
		fmt.Println(err)
	}
}
