package fynesimplechart

type Plot struct {
	Nodes      []Node
	Ticks      int
	XAxisTitle string
	YAxisTitle string
	Title      string

	ShowLine bool
}

func NewPlot(nodes []Node, title string) *Plot {
	plot := &Plot{
		Nodes: nodes,
		Ticks: 10,
		Title: title,
	}

	return plot
}
