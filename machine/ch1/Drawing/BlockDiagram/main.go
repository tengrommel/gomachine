package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

type Squares struct {
	plotter.XYs
}

func (s *Squares) Plot(c draw.Canvas, plt *plot.Plot) {
	trX, trY := plt.Transforms(&c)
	//c.SetColor(color.RGBA{
	//	R: 196,
	//	B: 128,
	//	A: 255,
	//})
	r := vg.Length(10.0)
	for _, p := range s.XYs {
		p1 := vg.Point{trX(p.X) - r, trY(p.Y) - r}
		p2 := vg.Point{trX(p.X) - r, trY(p.Y) + r}
		p3 := vg.Point{trX(p.X) + r, trY(p.Y) + r}
		p4 := vg.Point{trX(p.X) + r, trY(p.Y) - r}

		var p vg.Path
		p.Move(p1)
		p.Line(p2)
		p.Line(p3)
		p.Line(p4)
		p.Line(p1)
		p.Close()
		c.Fill(p)
	}
}

func main() {
	points := plotter.XYs{
		{2, 2},
		{4, 4},
		{6, 6},
		{8, 8},
		{10, 10},
	}

	s := Squares{points}

	p, _ := plot.New()
	p.Title.Text = "Squares"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	p.X.Min = 0
	p.X.Max = 20
	p.Y.Min = 0
	p.Y.Max = 20

	p.Add(&s)
	p.Save(4*vg.Inch, 4*vg.Inch, "squares.png")
}
