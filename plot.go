package fynesimplechart

import "image/color"

type Plot struct {
	Nodes      []Node
	Ticks      int
	XAxisTitle string
	YAxisTitle string
	Title      string

	ShowLine   bool
	LineWidth  float32
	PointSize  float32
	PlotColor  color.Color
	ShowPoints bool
}

func NewPlot(nodes []Node, title string) *Plot {
	plot := &Plot{
		Nodes:      nodes,
		Ticks:      10,
		Title:      title,
		ShowLine:   false,
		LineWidth:  1.5,
		PointSize:  3.0,
		PlotColor:  nil, // Will use auto-generated color if nil
		ShowPoints: true,
	}

	return plot
}
