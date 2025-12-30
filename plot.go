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

	// Area fill properties
	FillArea      bool        // Enable area fill
	FillColor     color.Color // Color for fill (nil uses PlotColor with transparency)
	FillToZero    bool        // Fill from curve to Y=0 axis
	FillToPlotIdx int         // Index of plot to fill to (-1 for none)

	// Bar chart properties
	ShowBars       bool    // Enable bar chart rendering
	BarWidth       float32 // Width of bars (0 = auto, default: 0.8 of available space)
	BarBorderWidth float32 // Border width for bars (0 = no border)
	BarBorderColor color.Color
}

func NewPlot(nodes []Node, title string) *Plot {
	plot := &Plot{
		Nodes:          nodes,
		Ticks:          10,
		Title:          title,
		ShowLine:       false,
		LineWidth:      1.5,
		PointSize:      3.0,
		PlotColor:      nil, // Will use auto-generated color if nil
		ShowPoints:     true,
		FillArea:       false,
		FillColor:      nil,
		FillToZero:     false,
		FillToPlotIdx:  -1,
		ShowBars:       false,
		BarWidth:       0.8, // 80% of available space per bar
		BarBorderWidth: 0,
		BarBorderColor: nil,
	}

	return plot
}
