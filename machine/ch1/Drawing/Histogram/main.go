package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	p, _ := plot.New()
	p.Title.Text = "Histogram"
	bins := plotter.XYs{
		{10, 10},
		{20, 20},
		{30, 50},
		{40, 20},
		{50, 10},
	}
	h, _ := plotter.NewHistogram(bins, 5)
	p.Add(h)
	p.Save(4*vg.Inch, 4*vg.Inch, "histogram.png")
}
